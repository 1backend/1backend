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
		req := &chattypes.AddMessageRequest{
			Message: &chattypes.Message{
				Id:   sdk.Id("msg"),
				Text: "hi there",
			},
		}

		_, _, err := userClient.ChatSvcAPI.AddMessage(context.Background(), "-").
			Body(
				openapi.ChatSvcAddMessageRequest{
					Message: &openapi.ChatSvcMessage{
						Id:   req.Message.Id,
						Text: openapi.PtrString(req.Message.Text),
					},
				},
			).
			Execute()
		require.Error(t, err)
	})

	t.Run("thread does not exist", func(t *testing.T) {
		req := &chattypes.AddMessageRequest{
			Message: &chattypes.Message{
				Id:       sdk.Id("msg"),
				ThreadId: "1",
				Text:     "hi there",
			},
		}

		_, _, err := userClient.ChatSvcAPI.AddMessage(context.Background(), req.Message.ThreadId).
			Body(
				openapi.ChatSvcAddMessageRequest{
					Message: &openapi.ChatSvcMessage{
						Id:   req.Message.Id,
						Text: openapi.PtrString(req.Message.Text),
					},
				},
			).
			Execute()
		require.Error(t, err)
	})

	t.Run("no user id should not fail", func(t *testing.T) {
		tid := sdk.Id("thr")
		title := "Test Thread Title"

		req := &chattypes.AddThreadRequest{
			Thread: &chattypes.Thread{
				Id:    tid,
				Title: title,
			},
		}

		_, _, err = userClient.ChatSvcAPI.AddThread(context.Background()).
			Body(
				openapi.ChatSvcAddThreadRequest{
					Thread: &openapi.ChatSvcThread{
						Id:    req.Thread.Id,
						Title: openapi.PtrString(req.Thread.Title),
					},
				},
			).
			Execute()
		require.NoError(t, err)
	})

	userId := "usr-1"
	var threadId string

	t.Run("create thread", func(t *testing.T) {
		tid := sdk.Id("thr")
		title := "Test Thread Title"

		req := &chattypes.AddThreadRequest{
			Thread: &chattypes.Thread{
				Id:      tid,
				Title:   title,
				UserIds: []string{userId},
			},
		}

		rsp, _, err := userClient.ChatSvcAPI.AddThread(context.Background()).
			Body(
				openapi.ChatSvcAddThreadRequest{
					Thread: &openapi.ChatSvcThread{
						Id:      req.Thread.Id,
						Title:   openapi.PtrString(req.Thread.Title),
						UserIds: []string{userId},
					},
				},
			).
			Execute()
		require.NoError(t, err)

		thread := rsp.Thread

		require.Equal(t, tid, thread.Id)
		require.Equal(t, title, *thread.Title)
		threadId = req.Thread.Id
	})

	t.Run("no user id", func(t *testing.T) {
		req := chattypes.AddMessageRequest{
			Message: &chattypes.Message{
				Id:       sdk.Id("msg"),
				ThreadId: threadId,
				Text:     "hi there",
			}}

		_, _, err := userClient.ChatSvcAPI.AddMessage(context.Background(), req.Message.ThreadId).
			Body(
				openapi.ChatSvcAddMessageRequest{
					Message: &openapi.ChatSvcMessage{
						Id:   req.Message.Id,
						Text: openapi.PtrString(req.Message.Text),
					},
				},
			).
			Execute()
		require.NoError(t, err)
	})
}
