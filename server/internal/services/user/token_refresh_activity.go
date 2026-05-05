/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	stdErrors "errors"
	"log/slog"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

const (
	tokenRefreshActivityTopic         = "user.token.refresh"
	tokenRefreshActivityFlushInterval = time.Minute
	tokenRefreshActivityFlushLimit    = int64(100)
)

type tokenRefreshActivityEvent struct {
	Id string `json:"id"`

	AppId   string `json:"appId"`
	AppHost string `json:"appHost"`
	UserId  string `json:"userId"`
	Device  string `json:"device"`

	BucketStart time.Time `json:"bucketStart"`
	BucketEnd   time.Time `json:"bucketEnd"`

	FirstRefreshedAt time.Time `json:"firstRefreshedAt"`
	LastRefreshedAt  time.Time `json:"lastRefreshedAt"`
	RefreshCount     int64     `json:"refreshCount"`
}

func (s *UserService) recordTokenRefreshActivity(
	token *user.Token,
	refreshedAt time.Time,
) error {
	if s.tokenActivityStore == nil {
		return errors.New("token refresh activity store is not configured")
	}
	if token == nil {
		return errors.New("token is nil")
	}

	refreshedAt = refreshedAt.UTC()
	device := token.Device
	if device == "" {
		device = unknownDevice
	}

	bucketStart := refreshedAt.Truncate(time.Hour)
	bucketEnd := bucketStart.Add(time.Hour)
	id := tokenRefreshActivityId(token.AppId, token.UserId, device, bucketStart)
	appHost, err := s.tokenRefreshActivityAppHost(token)
	if err != nil {
		return err
	}

	activityI, found, err := s.tokenActivityStore.Query(datastore.Id(id)).FindOne()
	if err != nil {
		return errors.Wrap(err, "failed to query token refresh activity")
	}

	if found {
		activity := activityI.(*user.TokenRefreshActivity)

		fields := map[string]interface{}{
			"refreshCount":    activity.RefreshCount + 1,
			"lastRefreshedAt": refreshedAt,
			"updatedAt":       refreshedAt,
			"appHost":         appHost,
		}
		if activity.FirstRefreshedAt.IsZero() ||
			refreshedAt.Before(activity.FirstRefreshedAt) {
			fields["firstRefreshedAt"] = refreshedAt
		}

		return s.tokenActivityStore.Query(datastore.Id(id)).UpdateFields(fields)
	}

	activity := &user.TokenRefreshActivity{
		Id:               id,
		CreatedAt:        refreshedAt,
		UpdatedAt:        refreshedAt,
		AppId:            token.AppId,
		AppHost:          appHost,
		UserId:           token.UserId,
		Device:           device,
		BucketStart:      bucketStart,
		BucketEnd:        bucketEnd,
		FirstRefreshedAt: refreshedAt,
		LastRefreshedAt:  refreshedAt,
		RefreshCount:     1,
	}

	err = s.tokenActivityStore.Create(activity)
	if err == nil {
		return nil
	}
	if !stdErrors.Is(err, datastore.ErrEntryAlreadyExists) {
		return errors.Wrap(err, "failed to create token refresh activity")
	}

	return s.recordTokenRefreshActivity(token, refreshedAt)
}

func (s *UserService) tokenRefreshActivityAppHost(token *user.Token) (string, error) {
	if token.App != nil && token.App.Host != "" {
		return token.App.Host, nil
	}

	app, found, err := s.appByID(token.AppId)
	if err != nil {
		return "", errors.Wrap(err, "failed to query token refresh activity app")
	}
	if !found {
		return "", errors.Errorf("app not found for token refresh activity: %s", token.AppId)
	}

	return app.Host, nil
}

func (s *UserService) publishTokenRefreshActivitiesLoop() {
	_, err := s.publishDueTokenRefreshActivities(
		context.Background(),
		time.Now(),
		tokenRefreshActivityFlushLimit,
	)
	if err != nil {
		logger.Error("Failed to publish token refresh activities", slog.Any("error", err))
	}

	ticker := time.NewTicker(tokenRefreshActivityFlushInterval)
	defer ticker.Stop()

	for range ticker.C {
		_, err := s.publishDueTokenRefreshActivities(
			context.Background(),
			time.Now(),
			tokenRefreshActivityFlushLimit,
		)
		if err != nil {
			logger.Error("Failed to publish token refresh activities", slog.Any("error", err))
		}
	}
}

func (s *UserService) publishDueTokenRefreshActivities(
	ctx context.Context,
	now time.Time,
	limit int64,
) (int, error) {
	if s.tokenActivityStore == nil || s.options.PubSub == nil {
		return 0, nil
	}
	if limit <= 0 {
		limit = tokenRefreshActivityFlushLimit
	}

	releaseLock, err := s.acquireTokenRefreshActivityPublishLock(ctx)
	if err != nil {
		return 0, err
	}
	defer releaseLock()

	return s.publishDueTokenRefreshActivitiesLocked(ctx, now.UTC(), limit)
}

func (s *UserService) publishDueTokenRefreshActivitiesLocked(
	ctx context.Context,
	now time.Time,
	limit int64,
) (int, error) {
	activities, err := s.tokenActivityStore.Query(
		datastore.Equals(datastore.Field("published"), false),
		datastore.LessThanOrEqual(datastore.Field("bucketEnd"), now),
	).OrderBy(datastore.OrderByField("bucketEnd", false)).
		Limit(limit).
		Find()
	if err != nil {
		return 0, errors.Wrap(err, "failed to query due token refresh activities")
	}

	published := 0
	for _, activityI := range activities {
		activity := activityI.(*user.TokenRefreshActivity)
		payload, err := json.Marshal(tokenRefreshActivityEvent{
			Id:               activity.Id,
			AppId:            activity.AppId,
			AppHost:          activity.AppHost,
			UserId:           activity.UserId,
			Device:           activity.Device,
			BucketStart:      activity.BucketStart,
			BucketEnd:        activity.BucketEnd,
			FirstRefreshedAt: activity.FirstRefreshedAt,
			LastRefreshedAt:  activity.LastRefreshedAt,
			RefreshCount:     activity.RefreshCount,
		})
		if err != nil {
			return published, errors.Wrap(err, "failed to marshal token refresh activity event")
		}

		_, err = s.options.PubSub.Publish(ctx, tokenRefreshActivityTopic, payload)
		if err != nil {
			return published, errors.Wrap(err, "failed to publish token refresh activity event")
		}

		publishedAt := now.UTC()
		err = s.tokenActivityStore.Query(
			datastore.Id(activity.Id),
			datastore.Equals(datastore.Field("published"), false),
		).UpdateFields(map[string]interface{}{
			"published":   true,
			"publishedAt": publishedAt,
			"updatedAt":   publishedAt,
		})
		if err != nil {
			return published, errors.Wrap(err, "failed to mark token refresh activity published")
		}
		published++
	}

	return published, nil
}

func (s *UserService) acquireTokenRefreshActivityPublishLock(
	ctx context.Context,
) (func(), error) {
	if s.options.Lock == nil {
		return func() {}, nil
	}

	key := "user-svc:token-refresh-activity:publish"
	err := s.options.Lock.Acquire(ctx, key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to acquire token refresh activity publish lock")
	}

	return func() {
		if err := s.options.Lock.Release(context.Background(), key); err != nil {
			logger.Error(
				"Failed to release token refresh activity publish lock",
				slog.String("lockKey", key),
				slog.Any("error", err),
			)
		}
	}, nil
}

func tokenRefreshActivityId(
	appId string,
	userId string,
	device string,
	bucketStart time.Time,
) string {
	raw := appId + "\x00" +
		userId + "\x00" +
		device + "\x00" +
		bucketStart.UTC().Format(time.RFC3339)
	sum := sha256.Sum256([]byte(raw))

	return "trfa_" + hex.EncodeToString(sum[:16])
}
