package chatservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/server/internal/di"
	chattypes "github.com/1backend/1backend/server/internal/services/chat/types"
)

func TestMessageCreatesThread(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe.Router)

	err = universe.StarterFunc()
	require.NoError(t, err)

	token, err := boot.RegisterUserAccount(
		options.ClientFactory.Client().UserSvcAPI,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := options.ClientFactory.Client(client.WithToken(token.Token))

	t.Run("no thread id", func(t *testing.T) {
		req := &chattypes.SaveMessageRequest{
			Id:   sdk.Id("msg"),
			Text: "hi there",
		}

		_, _, err := userClient.ChatSvcAPI.SaveMessage(context.Background(), "-").
			Body(
				openapi.ChatSvcSaveMessageRequest{
					Id:   &req.Id,
					Text: openapi.PtrString(req.Text),
				},
			).
			Execute()
		require.Error(t, err)
	})

	t.Run("thread does not exist", func(t *testing.T) {
		req := &chattypes.SaveMessageRequest{
			Id:       sdk.Id("msg"),
			ThreadId: "1",
			Text:     "hi there",
		}

		require.NotEmpty(t, req.ThreadId)

		_, _, err := userClient.ChatSvcAPI.SaveMessage(context.Background(), req.ThreadId).
			Body(
				openapi.ChatSvcSaveMessageRequest{
					Id:   &req.Id,
					Text: openapi.PtrString(req.Text),
				},
			).
			Execute()
		require.Error(t, err)
	})

	t.Run("no user id should not fail", func(t *testing.T) {
		tid := sdk.Id("thr")
		title := "Test Thread Title"

		req := &chattypes.SaveThreadRequest{
			Id:    tid,
			Title: title,
		}

		_, rsp, err := userClient.ChatSvcAPI.SaveThread(context.Background()).
			Body(
				openapi.ChatSvcSaveThreadRequest{
					Id:    &req.Id,
					Title: openapi.PtrString(req.Title),
				},
			).
			Execute()
		require.NoError(t, err, rsp)
	})

	userId := "usr-1"
	var threadId string

	t.Run("create thread", func(t *testing.T) {
		tid := sdk.Id("thr")
		title := "Test Thread Title"

		req := &chattypes.SaveThreadRequest{
			Id:      tid,
			Title:   title,
			UserIds: []string{userId},
		}

		rsp, _, err := userClient.ChatSvcAPI.SaveThread(context.Background()).
			Body(
				openapi.ChatSvcSaveThreadRequest{
					Id:      &req.Id,
					Title:   &req.Title,
					UserIds: []string{userId},
				},
			).
			Execute()
		require.NoError(t, err)

		thread := rsp.Thread

		require.Equal(t, tid, thread.Id)
		require.Equal(t, title, *thread.Title)
		threadId = req.Id
	})

	t.Run("no user id", func(t *testing.T) {
		req := chattypes.SaveMessageRequest{
			Id:       sdk.Id("msg"),
			ThreadId: threadId,
			Text:     "hi there",
		}

		require.NotEmpty(t, req.ThreadId)
		_, _, err := userClient.ChatSvcAPI.SaveMessage(context.Background(), req.ThreadId).
			Body(
				openapi.ChatSvcSaveMessageRequest{
					Id:   &req.Id,
					Text: &req.Text,
				},
			).
			Execute()
		require.NoError(t, err)
	})
}
