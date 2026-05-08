package userservice

import (
	"testing"
	"time"

	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/stretchr/testify/require"
)

func TestUserLifecycleEventPayloadIncludesAppHost(t *testing.T) {
	eventTime := time.Date(2026, 5, 7, 12, 30, 0, 0, time.UTC)

	payload := userLifecycleEventPayload(
		&user.App{Id: "app-example", Host: "example.com"},
		&user.Token{
			AppId:  "app-example",
			UserId: "usr-1",
			Device: "desktop",
		},
		"user-slug",
		eventTime,
	)

	require.Equal(t, "app-example", payload["appId"])
	require.Equal(t, "example.com", payload["appHost"])
	require.Equal(t, "usr-1", payload["userId"])
	require.Equal(t, "desktop", payload["device"])
	require.Equal(t, "user-slug", payload["slug"])
	require.Equal(t, eventTime, payload["time"])
}

func TestUserLifecycleEventPayloadFallsBackToTokenAppHost(t *testing.T) {
	payload := userLifecycleEventPayload(
		nil,
		&user.Token{
			AppId:  "app-example",
			App:    &user.App{Id: "app-example", Host: "example.com"},
			UserId: "usr-1",
			Device: "desktop",
		},
		"",
		time.Date(2026, 5, 7, 12, 30, 0, 0, time.UTC),
	)

	require.Equal(t, "example.com", payload["appHost"])
	require.NotContains(t, payload, "slug")
}
