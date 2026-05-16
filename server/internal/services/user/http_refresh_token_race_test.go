package userservice

import (
	"context"
	"encoding/json"
	"sync"
	"testing"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/infra"
	distlock "github.com/1backend/1backend/sdk/go/lock/local"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/stretchr/testify/require"
)

func TestRefreshTokenConcurrentDifferentTokensKeepsSingleActiveToken(t *testing.T) {
	t.Parallel()

	s := newRefreshRaceTestService(t)

	tok0, err := s.register(
		user.DefaultAppHost,
		"refresh-race-"+sdk.Id(""),
		"password",
		"Refresh Race",
		[]string{user.RoleUser},
	)
	require.NoError(t, err)

	waitForReplacementCacheExpiry()
	tok1, err := s.refreshToken(context.Background(), tok0.Token)
	require.NoError(t, err)

	waitForReplacementCacheExpiry()
	tok2, err := s.refreshToken(context.Background(), tok1.Token)
	require.NoError(t, err)

	waitForReplacementCacheExpiry()

	realTokenStore := s.tokenStore
	barrier := newActiveTokenLookupBarrier()
	s.tokenStore = &activeTokenLookupBarrierStore{
		DataStore: realTokenStore,
		barrier:   barrier,
	}

	start := make(chan struct{})
	inputs := []*user.Token{tok0, tok1}
	outputs := make([]*user.Token, len(inputs))
	errs := make([]error, len(inputs))

	var wg sync.WaitGroup
	for i, input := range inputs {
		wg.Add(1)
		go func(i int, input *user.Token) {
			defer wg.Done()
			<-start
			outputs[i], errs[i] = s.refreshToken(context.Background(), input.Token)
		}(i, input)
	}

	close(start)
	wg.Wait()

	for _, err := range errs {
		require.NoError(t, err)
	}

	for _, output := range outputs {
		_, found, err := s.tokenStore.Query(
			datastore.Equals(datastore.Field("token"), output.Token),
		).FindOne()
		require.NoError(t, err)
		require.True(t, found, "refresh returned a token that pruning deleted")
	}

	activeTokens, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("appId"), tok2.AppId),
		datastore.Equals(datastore.Field("userId"), tok2.UserId),
		datastore.Equals(datastore.Field("device"), tok2.Device),
		datastore.Equals(datastore.Field("active"), true),
	).Find()
	require.NoError(t, err)
	require.Len(t, activeTokens, 1)

	deviceTokens, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("appId"), tok2.AppId),
		datastore.Equals(datastore.Field("userId"), tok2.UserId),
		datastore.Equals(datastore.Field("device"), tok2.Device),
	).Find()
	require.NoError(t, err)
	require.LessOrEqual(t, len(deviceTokens), 3)
}

func TestRefreshTokenMarksInFlightTokenBeforeLockKeyParsing(t *testing.T) {
	s := newRefreshRaceTestService(t)

	tok0 := registerRefreshRaceUser(t, s, "refresh-prune-window")

	waitForReplacementCacheExpiry()
	tok1, err := s.refreshToken(context.Background(), tok0.Token)
	require.NoError(t, err)

	waitForReplacementCacheExpiry()
	_, err = s.refreshToken(context.Background(), tok1.Token)
	require.NoError(t, err)

	waitForReplacementCacheExpiry()

	blocker := &blockingParseUnverifiedAuthorizer{
		AuthorizerImpl: auth.AuthorizerImpl{},
		token:          tok0.Token,
		entered:        make(chan struct{}),
		release:        make(chan struct{}),
	}
	s.options.Authorizer = blocker

	var blockedErr error
	blockedDone := make(chan struct{})
	go func() {
		_, blockedErr = s.refreshToken(context.Background(), tok0.Token)
		close(blockedDone)
	}()

	require.Eventually(t, func() bool {
		select {
		case <-blocker.entered:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)

	var competingErr error
	competingDone := make(chan struct{})
	go func() {
		_, competingErr = s.refreshToken(context.Background(), tok1.Token)
		close(competingDone)
	}()

	close(blocker.release)

	require.Eventually(t, func() bool {
		select {
		case <-blockedDone:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)
	require.NoError(t, blockedErr)

	require.Eventually(t, func() bool {
		select {
		case <-competingDone:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)
	require.NoError(t, competingErr)
}

func TestRefreshTokenPruneDoesNotDeleteTokenMarkedAfterPruneRead(t *testing.T) {
	s := newRefreshRaceTestService(t)

	tok0 := registerRefreshRaceUser(t, s, "refresh-prune-stale-read")

	waitForReplacementCacheExpiry()
	tok1, err := s.refreshToken(context.Background(), tok0.Token)
	require.NoError(t, err)

	waitForReplacementCacheExpiry()
	tok2, err := s.refreshToken(context.Background(), tok1.Token)
	require.NoError(t, err)

	waitForReplacementCacheExpiry()

	realTokenStore := s.tokenStore
	pruner := &stalePruneWindowStore{
		DataStore:    realTokenStore,
		token:        tok0.Token,
		pruneRead:    make(chan struct{}),
		releasePrune: make(chan struct{}),
		markerDone:   make(chan struct{}),
	}
	s.tokenStore = pruner

	var pruneErr error
	pruneDone := make(chan struct{})
	go func() {
		_, pruneErr = s.refreshToken(context.Background(), tok2.Token)
		close(pruneDone)
	}()

	require.Eventually(t, func() bool {
		select {
		case <-pruner.pruneRead:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)

	var originalErr error
	originalDone := make(chan struct{})
	go func() {
		_, originalErr = s.refreshToken(context.Background(), tok0.Token)
		close(originalDone)
	}()

	require.Eventually(t, func() bool {
		select {
		case <-pruner.markerDone:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)

	close(pruner.releasePrune)

	require.Eventually(t, func() bool {
		select {
		case <-pruneDone:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)
	require.NoError(t, pruneErr)

	require.Eventually(t, func() bool {
		select {
		case <-originalDone:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)
	require.NoError(t, originalErr)

	_, found, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("token"), tok0.Token),
	).FindOne()
	require.NoError(t, err)
	require.True(t, found, "token marked as recently refreshed was pruned from a stale read")
}

func TestRefreshTokenPruneDeleteIsConditionalOnCurrentMarker(t *testing.T) {
	s := newRefreshRaceTestService(t)

	tok0 := registerRefreshRaceUser(t, s, "refresh-prune-delete-window")

	waitForReplacementCacheExpiry()
	tok1, err := s.refreshToken(context.Background(), tok0.Token)
	require.NoError(t, err)

	waitForReplacementCacheExpiry()
	tok2, err := s.refreshToken(context.Background(), tok1.Token)
	require.NoError(t, err)

	waitForReplacementCacheExpiry()

	realTokenStore := s.tokenStore
	pruner := &deleteWindowStore{
		DataStore:     realTokenStore,
		token:         tok0.Token,
		targetId:      tok0.Id,
		deleteStarted: make(chan struct{}),
		releaseDelete: make(chan struct{}),
		markerDone:    make(chan struct{}),
	}
	s.tokenStore = pruner

	var pruneErr error
	pruneDone := make(chan struct{})
	go func() {
		_, pruneErr = s.refreshToken(context.Background(), tok2.Token)
		close(pruneDone)
	}()

	require.Eventually(t, func() bool {
		select {
		case <-pruner.deleteStarted:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)

	var originalErr error
	originalDone := make(chan struct{})
	go func() {
		_, originalErr = s.refreshToken(context.Background(), tok0.Token)
		close(originalDone)
	}()

	require.Eventually(t, func() bool {
		select {
		case <-pruner.markerDone:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)

	close(pruner.releaseDelete)

	require.Eventually(t, func() bool {
		select {
		case <-pruneDone:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)
	require.NoError(t, pruneErr)

	require.Eventually(t, func() bool {
		select {
		case <-originalDone:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)
	require.NoError(t, originalErr)

	_, found, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("token"), tok0.Token),
	).FindOne()
	require.NoError(t, err)
	require.True(t, found, "token marked before prune delete completed was deleted")
}

func TestRefreshTokenRepeatedOriginalTokenUseStaysBounded(t *testing.T) {
	s := newRefreshRaceTestServiceWithExpiration(t, 2*time.Millisecond)

	original := registerRefreshRaceUser(t, s, "refresh-repeated-original")
	startedAt := time.Now()

	const refreshCount = 200
	for i := 0; i < refreshCount; i++ {
		waitForReplacementCacheKeyExpiry(t, s, original.Token)

		refreshed, err := s.refreshToken(context.Background(), original.Token)
		require.NoError(t, err)
		require.NotEmpty(t, refreshed.Token)

		if i%10 == 0 {
			assertRefreshRaceTokenState(t, s, original)
		}
	}

	assertRefreshRaceTokenState(t, s, original)

	tokenI, found, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("token"), original.Token),
	).FindOne()
	require.NoError(t, err)
	require.True(t, found, "repeatedly used original token was pruned")

	storedOriginal := tokenI.(*user.Token)
	require.NotNil(t, storedOriginal.LastRefreshedAt)
	require.True(t, storedOriginal.LastRefreshedAt.After(startedAt))
}

func TestRefreshTokenConcurrentIssueTokenKeepsSingleActiveToken(t *testing.T) {
	s := newRefreshRaceTestServiceWithExpiration(t, time.Minute)

	tok := registerRefreshRaceUser(t, s, "refresh-issue-race")
	usr, err := s.readSelf(tok.UserId)
	require.NoError(t, err)

	realTokenStore := s.tokenStore
	blocker := &blockActiveFalseStore{
		DataStore: realTokenStore,
		blocked:   make(chan struct{}),
		release:   make(chan struct{}),
	}
	s.tokenStore = blocker

	var refreshErr error
	refreshDone := make(chan struct{})
	go func() {
		_, refreshErr = s.refreshToken(context.Background(), tok.Token)
		close(refreshDone)
	}()

	require.Eventually(t, func() bool {
		select {
		case <-blocker.blocked:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)

	var issued *user.Token
	var issueErr error
	issueDone := make(chan struct{})
	go func() {
		issued, issueErr = s.issueToken(tok.AppId, usr, tok.Device)
		close(issueDone)
	}()

	close(blocker.release)

	require.Eventually(t, func() bool {
		select {
		case <-refreshDone:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)
	require.NoError(t, refreshErr)

	require.Eventually(t, func() bool {
		select {
		case <-issueDone:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)
	require.NoError(t, issueErr)
	require.NotNil(t, issued)
	require.NotEmpty(t, issued.Token)

	_, found, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("token"), issued.Token),
	).FindOne()
	require.NoError(t, err)
	require.True(t, found, "issueToken returned a token that concurrent refresh pruning deleted")

	activeTokens, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("appId"), tok.AppId),
		datastore.Equals(datastore.Field("userId"), tok.UserId),
		datastore.Equals(datastore.Field("device"), tok.Device),
		datastore.Equals(datastore.Field("active"), true),
	).Find()
	require.NoError(t, err)
	require.Len(t, activeTokens, 1)
}

func TestRefreshTokenReplacementCacheDoesNotReturnStaleTokenAfterInactivation(t *testing.T) {
	s := newRefreshRaceTestServiceWithExpiration(t, time.Minute)

	tok0 := registerRefreshRaceUser(t, s, "refresh-cache-stale")

	tok1, err := s.refreshToken(context.Background(), tok0.Token)
	require.NoError(t, err)
	require.NotEmpty(t, tok1.Token)

	role := "user-svc:cache-stale:" + sdk.Id("")
	err = s.assignRole(tok0.AppId, tok0.UserId, role)
	require.NoError(t, err)

	refreshed, err := s.refreshToken(context.Background(), tok0.Token)
	require.NoError(t, err)

	claims, err := s.options.Authorizer.ParseJWT(s.publicKeyPem, refreshed.Token)
	require.NoError(t, err)
	require.Contains(t, claims.Roles, role)
}

func TestRefreshTokenDifferentDevicesDoNotShareEntryLock(t *testing.T) {
	s := newRefreshRaceTestServiceWithExpiration(t, time.Minute)

	tokA := registerRefreshRaceUser(t, s, "refresh-entry-device-a")
	usr, err := s.readSelf(tokA.UserId)
	require.NoError(t, err)

	tokB, err := s.issueToken(
		tokA.AppId,
		usr,
		"refresh-entry-device-b-"+sdk.Id(""),
	)
	require.NoError(t, err)

	realTokenStore := s.tokenStore
	blocker := &blockTokenLastRefreshedStore{
		DataStore: realTokenStore,
		token:     tokA.Token,
		blocked:   make(chan struct{}),
		release:   make(chan struct{}),
	}
	s.tokenStore = blocker

	var refreshAErr error
	refreshADone := make(chan struct{})
	go func() {
		_, refreshAErr = s.refreshToken(context.Background(), tokA.Token)
		close(refreshADone)
	}()

	require.Eventually(t, func() bool {
		select {
		case <-blocker.blocked:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)

	refreshBDone := make(chan error, 1)
	go func() {
		_, err := s.refreshToken(context.Background(), tokB.Token)
		refreshBDone <- err
	}()

	select {
	case err := <-refreshBDone:
		require.NoError(t, err)
	case <-time.After(200 * time.Millisecond):
		t.Fatal("refresh on another device was blocked by token A marker")
	}

	close(blocker.release)

	require.Eventually(t, func() bool {
		select {
		case <-refreshADone:
			return true
		default:
			return false
		}
	}, time.Second, time.Millisecond)
	require.NoError(t, refreshAErr)
}

func BenchmarkRefreshTokenPruneWithConditionalDelete(b *testing.B) {
	s := newRefreshRaceTestServiceWithExpiration(b, time.Minute)
	current := registerRefreshRaceUser(b, s, "refresh-prune-bench")

	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		next, err := s.refreshToken(context.Background(), current.Token)
		if err != nil {
			b.Fatal(err)
		}
		current = next
	}
}

func newRefreshRaceTestService(t testing.TB) *UserService {
	t.Helper()

	return newRefreshRaceTestServiceWithExpiration(t, 20*time.Millisecond)
}

func newRefreshRaceTestServiceWithExpiration(
	t testing.TB,
	tokenExpiration time.Duration,
) *UserService {
	t.Helper()

	dataStoreFactory, err := infra.NewDataStoreFactory(infra.DataStoreConfig{
		Test:        true,
		HomeDir:     t.TempDir(),
		TablePrefix: "t_" + sdk.Id(""),
	})
	require.NoError(t, err)

	s, err := NewUserService(&universe.Options{
		Test:             true,
		TokenExpiration:  tokenExpiration,
		DataStoreFactory: dataStoreFactory,
		Lock:             distlock.NewLocalDistributedLock(),
		Authorizer:       auth.AuthorizerImpl{},
	})
	require.NoError(t, err)
	require.NoError(t, s.bootstrap())

	return s
}

func registerRefreshRaceUser(
	t testing.TB,
	s *UserService,
	slugPrefix string,
) *user.Token {
	t.Helper()

	tok, err := s.register(
		user.DefaultAppHost,
		slugPrefix+"-"+sdk.Id(""),
		"password",
		"Refresh Race",
		[]string{user.RoleUser},
	)
	require.NoError(t, err)

	return tok
}

func waitForReplacementCacheExpiry() {
	time.Sleep(30 * time.Millisecond)
}

func waitForReplacementCacheKeyExpiry(
	t testing.TB,
	s *UserService,
	token string,
) {
	t.Helper()

	require.Eventually(t, func() bool {
		_, found := s.cachedReplacementToken(generateCacheKey(token))
		return !found
	}, time.Second, time.Millisecond)
}

func assertRefreshRaceTokenState(
	t testing.TB,
	s *UserService,
	original *user.Token,
) {
	t.Helper()

	_, found, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("token"), original.Token),
	).FindOne()
	require.NoError(t, err)
	require.True(t, found, "original token was pruned")

	activeTokens, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("appId"), original.AppId),
		datastore.Equals(datastore.Field("userId"), original.UserId),
		datastore.Equals(datastore.Field("device"), original.Device),
		datastore.Equals(datastore.Field("active"), true),
	).Find()
	require.NoError(t, err)
	require.Len(t, activeTokens, 1)

	deviceTokens, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("appId"), original.AppId),
		datastore.Equals(datastore.Field("userId"), original.UserId),
		datastore.Equals(datastore.Field("device"), original.Device),
	).Find()
	require.NoError(t, err)
	require.LessOrEqual(t, len(deviceTokens), 3)
}

type blockingParseUnverifiedAuthorizer struct {
	auth.AuthorizerImpl
	token   string
	entered chan struct{}
	release chan struct{}
	once    sync.Once
}

func (a *blockingParseUnverifiedAuthorizer) ParseJWTUnverified(
	token string,
) (*auth.Claims, error) {
	claims, err := a.AuthorizerImpl.ParseJWTUnverified(token)
	if err != nil {
		return nil, err
	}

	if token == a.token {
		a.once.Do(func() {
			close(a.entered)
			<-a.release
		})
	}

	return claims, nil
}

type blockActiveFalseStore struct {
	datastore.DataStore
	blocked chan struct{}
	release chan struct{}
	once    sync.Once
}

func (s *blockActiveFalseStore) Query(filters ...datastore.Filter) datastore.QueryBuilder {
	return &blockActiveFalseQuery{
		QueryBuilder: s.DataStore.Query(filters...),
		store:        s,
	}
}

type blockActiveFalseQuery struct {
	datastore.QueryBuilder
	store *blockActiveFalseStore
}

func (q *blockActiveFalseQuery) UpdateFields(fields map[string]interface{}) error {
	err := q.QueryBuilder.UpdateFields(fields)
	if err != nil {
		return err
	}

	active, ok := fields["active"].(bool)
	if ok && !active {
		q.store.once.Do(func() {
			close(q.store.blocked)
			<-q.store.release
		})
	}

	return nil
}

type blockTokenLastRefreshedStore struct {
	datastore.DataStore
	token   string
	blocked chan struct{}
	release chan struct{}
	once    sync.Once
}

func (s *blockTokenLastRefreshedStore) Query(filters ...datastore.Filter) datastore.QueryBuilder {
	return &blockTokenLastRefreshedQuery{
		QueryBuilder: s.DataStore.Query(filters...),
		store:        s,
		filters:      filters,
	}
}

type blockTokenLastRefreshedQuery struct {
	datastore.QueryBuilder
	store   *blockTokenLastRefreshedStore
	filters []datastore.Filter
}

func (q *blockTokenLastRefreshedQuery) UpdateFields(fields map[string]interface{}) error {
	_, updatesLastRefreshedAt := fields["lastRefreshedAt"]
	if updatesLastRefreshedAt && hasEqualsFilterValue(q.filters, "token", q.store.token) {
		q.store.once.Do(func() {
			close(q.store.blocked)
			<-q.store.release
		})
	}

	return q.QueryBuilder.UpdateFields(fields)
}

type stalePruneWindowStore struct {
	datastore.DataStore
	token        string
	pruneRead    chan struct{}
	releasePrune chan struct{}
	markerDone   chan struct{}
	pruneOnce    sync.Once
	markerOnce   sync.Once
}

func (s *stalePruneWindowStore) Query(filters ...datastore.Filter) datastore.QueryBuilder {
	return &stalePruneWindowQuery{
		QueryBuilder: s.DataStore.Query(filters...),
		store:        s,
		filters:      filters,
	}
}

type stalePruneWindowQuery struct {
	datastore.QueryBuilder
	store   *stalePruneWindowStore
	filters []datastore.Filter
}

func (q *stalePruneWindowQuery) Find() ([]datastore.Row, error) {
	rows, err := q.QueryBuilder.Find()
	if err != nil {
		return nil, err
	}

	if isRefreshPruneLookup(q.filters) {
		q.store.pruneOnce.Do(func() {
			close(q.store.pruneRead)
			<-q.store.releasePrune
		})
	}

	return rows, nil
}

func (q *stalePruneWindowQuery) UpdateFields(fields map[string]interface{}) error {
	err := q.QueryBuilder.UpdateFields(fields)
	if err != nil {
		return err
	}

	_, updatesLastRefreshedAt := fields["lastRefreshedAt"]
	if updatesLastRefreshedAt && hasEqualsFilterValue(q.filters, "token", q.store.token) {
		q.store.markerOnce.Do(func() {
			close(q.store.markerDone)
		})
	}

	return nil
}

type deleteWindowStore struct {
	datastore.DataStore
	token         string
	targetId      string
	deleteStarted chan struct{}
	releaseDelete chan struct{}
	markerDone    chan struct{}
	deleteOnce    sync.Once
	markerOnce    sync.Once
}

func (s *deleteWindowStore) Query(filters ...datastore.Filter) datastore.QueryBuilder {
	return &deleteWindowQuery{
		QueryBuilder: s.DataStore.Query(filters...),
		store:        s,
		filters:      filters,
	}
}

type deleteWindowQuery struct {
	datastore.QueryBuilder
	store   *deleteWindowStore
	filters []datastore.Filter
}

func (q *deleteWindowQuery) Delete() error {
	if hasEqualsFilterValue(q.filters, "id", q.store.targetId) {
		q.store.deleteOnce.Do(func() {
			close(q.store.deleteStarted)
			<-q.store.releaseDelete
		})
	}

	return q.QueryBuilder.Delete()
}

func (q *deleteWindowQuery) UpdateFields(fields map[string]interface{}) error {
	err := q.QueryBuilder.UpdateFields(fields)
	if err != nil {
		return err
	}

	_, updatesLastRefreshedAt := fields["lastRefreshedAt"]
	if updatesLastRefreshedAt && hasEqualsFilterValue(q.filters, "token", q.store.token) {
		q.store.markerOnce.Do(func() {
			close(q.store.markerDone)
		})
	}

	return nil
}

type activeTokenLookupBarrierStore struct {
	datastore.DataStore
	barrier *activeTokenLookupBarrier
}

func (s *activeTokenLookupBarrierStore) Query(filters ...datastore.Filter) datastore.QueryBuilder {
	query := s.DataStore.Query(filters...)
	if isActiveTokenLookup(filters) {
		return &activeTokenLookupBarrierQuery{
			QueryBuilder: query,
			barrier:      s.barrier,
		}
	}
	return query
}

type activeTokenLookupBarrierQuery struct {
	datastore.QueryBuilder
	barrier *activeTokenLookupBarrier
}

func (q *activeTokenLookupBarrierQuery) FindOne() (datastore.Row, bool, error) {
	row, found, err := q.QueryBuilder.FindOne()
	q.barrier.wait()
	return row, found, err
}

type activeTokenLookupBarrier struct {
	mu    sync.Mutex
	count int
	done  chan struct{}
}

func newActiveTokenLookupBarrier() *activeTokenLookupBarrier {
	return &activeTokenLookupBarrier{
		done: make(chan struct{}),
	}
}

func (b *activeTokenLookupBarrier) wait() {
	b.mu.Lock()
	b.count++
	if b.count == 2 {
		close(b.done)
	}
	done := b.done
	b.mu.Unlock()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
	}
}

func isActiveTokenLookup(filters []datastore.Filter) bool {
	return hasEqualsFilter(filters, "appId") &&
		hasEqualsFilter(filters, "userId") &&
		hasEqualsFilter(filters, "device") &&
		hasEqualsFilterValue(filters, "active", true)
}

func isRefreshPruneLookup(filters []datastore.Filter) bool {
	return hasEqualsFilter(filters, "appId") &&
		hasEqualsFilter(filters, "userId") &&
		hasEqualsFilter(filters, "device") &&
		!hasEqualsFilter(filters, "active") &&
		!hasEqualsFilter(filters, "token")
}

func hasEqualsFilter(filters []datastore.Filter, field string) bool {
	for _, filter := range filters {
		if filter.Op == datastore.OpEquals && filter.FieldIs(field) {
			return true
		}
	}
	return false
}

func hasEqualsFilterValue(filters []datastore.Filter, field string, value any) bool {
	for _, filter := range filters {
		if filter.Op != datastore.OpEquals || !filter.FieldIs(field) {
			continue
		}

		var values []any
		if err := json.Unmarshal([]byte(filter.ValuesJson), &values); err != nil {
			return false
		}
		if len(values) == 1 && values[0] == value {
			return true
		}
	}
	return false
}
