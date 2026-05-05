package userservice

import (
	"context"
	"encoding/json"
	stdErrors "errors"
	"fmt"
	"sync"
	"testing"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/infra"
	distlock "github.com/1backend/1backend/sdk/go/lock/local"
	"github.com/1backend/1backend/sdk/go/pubsub"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/stretchr/testify/require"
)

func TestTokenRefreshActivityAggregatesManyRefreshesIntoOneHourlyEvent(t *testing.T) {
	t.Parallel()

	s, ps := newRefreshActivityTestService(t)
	ctx := context.Background()

	tok := registerRefreshRaceUser(t, s, "refresh-activity-aggregate")
	for i := 0; i < 12; i++ {
		refreshed, err := s.refreshToken(ctx, tok.Token)
		require.NoError(t, err)
		tok = refreshed
	}

	activities := tokenRefreshActivities(t, s)
	require.Len(t, activities, 1)
	require.Equal(t, int64(12), activities[0].RefreshCount)
	require.False(t, activities[0].Published)
	require.Equal(t, 0, ps.publishCallCount())

	published, err := s.publishDueTokenRefreshActivities(ctx, time.Now().Add(2*time.Hour), 100)
	require.NoError(t, err)
	require.Equal(t, 1, published)

	messages := ps.messages()
	require.Len(t, messages, 1)
	require.Equal(t, tokenRefreshActivityTopic, messages[0].Topic)

	event := tokenRefreshActivityEvent{}
	require.NoError(t, json.Unmarshal(messages[0].Payload, &event))
	require.Equal(t, activities[0].Id, event.Id)
	require.Equal(t, tok.AppId, event.AppId)
	require.Equal(t, user.DefaultAppHost, event.AppHost)
	require.Equal(t, int64(12), event.RefreshCount)

	published, err = s.publishDueTokenRefreshActivities(ctx, time.Now().Add(2*time.Hour), 100)
	require.NoError(t, err)
	require.Equal(t, 0, published)
	require.Equal(t, 1, ps.publishCallCount())
}

func TestTokenRefreshActivityDoesNotCountCachedDuplicateRefreshes(t *testing.T) {
	t.Parallel()

	s, _ := newRefreshActivityTestService(t)
	ctx := context.Background()

	tok := registerRefreshRaceUser(t, s, "refresh-activity-cache")

	refreshed, err := s.refreshToken(ctx, tok.Token)
	require.NoError(t, err)

	cached, err := s.refreshToken(ctx, tok.Token)
	require.NoError(t, err)
	require.Equal(t, refreshed.Id, cached.Id)

	activities := tokenRefreshActivities(t, s)
	require.Len(t, activities, 1)
	require.Equal(t, int64(1), activities[0].RefreshCount)
}

func TestTokenRefreshActivityPublishRetryDoesNotMarkFailedPublish(t *testing.T) {
	t.Parallel()

	s, ps := newRefreshActivityTestService(t)
	ctx := context.Background()

	refreshedAt := time.Now().UTC().Add(-2 * time.Hour).Truncate(time.Hour).Add(5 * time.Minute)
	token := &user.Token{
		AppId:  "app-retry",
		App:    &user.App{Id: "app-retry", Host: "retry.example.com"},
		UserId: "usr-retry",
		Device: "device-retry",
	}
	require.NoError(t, s.recordTokenRefreshActivity(token, refreshedAt))

	ps.setFail(true)
	published, err := s.publishDueTokenRefreshActivities(ctx, time.Now(), 100)
	require.Error(t, err)
	require.Equal(t, 0, published)
	require.Equal(t, 1, ps.publishCallCount())

	activities := tokenRefreshActivities(t, s)
	require.Len(t, activities, 1)
	require.False(t, activities[0].Published)
	require.Nil(t, activities[0].PublishedAt)

	ps.setFail(false)
	published, err = s.publishDueTokenRefreshActivities(ctx, time.Now(), 100)
	require.NoError(t, err)
	require.Equal(t, 1, published)
	require.Equal(t, 2, ps.publishCallCount())

	activities = tokenRefreshActivities(t, s)
	require.Len(t, activities, 1)
	require.True(t, activities[0].Published)
	require.NotNil(t, activities[0].PublishedAt)
}

func TestTokenRefreshActivityFlushLimitBoundsPublishWork(t *testing.T) {
	t.Parallel()

	s, ps := newRefreshActivityTestService(t)
	ctx := context.Background()

	base := time.Now().UTC().Add(-3 * time.Hour).Truncate(time.Hour)
	for i := 0; i < 25; i++ {
		token := &user.Token{
			AppId:  "app-limit",
			App:    &user.App{Id: "app-limit", Host: "limit.example.com"},
			UserId: "usr-limit",
			Device: fmt.Sprintf("device-%02d", i),
		}
		require.NoError(t, s.recordTokenRefreshActivity(token, base.Add(time.Duration(i)*time.Second)))
	}

	published, err := s.publishDueTokenRefreshActivities(ctx, time.Now(), 10)
	require.NoError(t, err)
	require.Equal(t, 10, published)
	require.Equal(t, 10, ps.publishCallCount())
	require.Equal(t, int64(15), unpublishedTokenRefreshActivityCount(t, s))

	published, err = s.publishDueTokenRefreshActivities(ctx, time.Now(), 10)
	require.NoError(t, err)
	require.Equal(t, 10, published)
	require.Equal(t, 20, ps.publishCallCount())
	require.Equal(t, int64(5), unpublishedTokenRefreshActivityCount(t, s))
}

func newRefreshActivityTestService(t *testing.T) (*UserService, *recordingPubSub) {
	t.Helper()

	dataStoreFactory, err := infra.NewDataStoreFactory(infra.DataStoreConfig{
		Test:        true,
		HomeDir:     t.TempDir(),
		TablePrefix: "t_" + sdk.Id(""),
	})
	require.NoError(t, err)

	ps := &recordingPubSub{}
	s, err := NewUserService(&universe.Options{
		Test:             true,
		TokenExpiration:  time.Minute,
		DataStoreFactory: dataStoreFactory,
		Lock:             distlock.NewLocalDistributedLock(),
		Authorizer:       auth.AuthorizerImpl{},
		PubSub:           ps,
	})
	require.NoError(t, err)
	require.NoError(t, s.bootstrap())

	return s, ps
}

func tokenRefreshActivities(t *testing.T, s *UserService) []*user.TokenRefreshActivity {
	t.Helper()

	rows, err := s.tokenActivityStore.Query().
		OrderBy(datastore.OrderByField("bucketStart", false)).
		Find()
	require.NoError(t, err)

	activities := make([]*user.TokenRefreshActivity, 0, len(rows))
	for _, row := range rows {
		activities = append(activities, row.(*user.TokenRefreshActivity))
	}

	return activities
}

func unpublishedTokenRefreshActivityCount(t *testing.T, s *UserService) int64 {
	t.Helper()

	count, err := s.tokenActivityStore.Query(
		datastore.Equals(datastore.Field("published"), false),
	).Count()
	require.NoError(t, err)

	return count
}

type recordingPubSub struct {
	mu      sync.Mutex
	fail    bool
	calls   int
	records []pubsub.Message
}

func (p *recordingPubSub) Publish(
	ctx context.Context,
	topic string,
	payload []byte,
) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.calls++
	if p.fail {
		return "", stdErrors.New("publish failed")
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	id := fmt.Sprintf("msg-%d", p.calls)
	p.records = append(p.records, pubsub.Message{
		ID:      id,
		Topic:   topic,
		Payload: append([]byte(nil), payload...),
	})

	return id, nil
}

func (p *recordingPubSub) Subscribe(
	ctx context.Context,
	subscriberId string,
	topic string,
	options ...pubsub.SubscribeOption,
) (pubsub.Subscription, error) {
	_ = ctx
	_ = subscriberId
	_ = topic
	_ = options

	return nil, stdErrors.New("subscribe not implemented")
}

func (p *recordingPubSub) Close() error {
	return nil
}

func (p *recordingPubSub) setFail(fail bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.fail = fail
}

func (p *recordingPubSub) publishCallCount() int {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.calls
}

func (p *recordingPubSub) messages() []pubsub.Message {
	p.mu.Lock()
	defer p.mu.Unlock()

	messages := make([]pubsub.Message, len(p.records))
	copy(messages, p.records)

	return messages
}
