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

func newRefreshRaceTestService(t *testing.T) *UserService {
	t.Helper()

	dataStoreFactory, err := infra.NewDataStoreFactory(infra.DataStoreConfig{
		Test:        true,
		HomeDir:     t.TempDir(),
		TablePrefix: "t_" + sdk.Id(""),
	})
	require.NoError(t, err)

	s, err := NewUserService(&universe.Options{
		Test:             true,
		TokenExpiration:  20 * time.Millisecond,
		DataStoreFactory: dataStoreFactory,
		Lock:             distlock.NewLocalDistributedLock(),
		Authorizer:       auth.AuthorizerImpl{},
	})
	require.NoError(t, err)
	require.NoError(t, s.bootstrap())

	return s
}

func waitForReplacementCacheExpiry() {
	time.Sleep(30 * time.Millisecond)
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
