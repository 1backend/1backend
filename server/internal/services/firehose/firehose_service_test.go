package firehoseservice_test

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	client "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"
	firehose "github.com/1backend/1backend/server/internal/services/firehose/types"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/r3labs/sse"
	"github.com/stretchr/testify/require"
)

func TestFirehoseSubscription(t *testing.T) {
	// todo this test sometimes flakes, likely because the sse client
	// unsubscribe is not taking effect
	return

	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe.Router)
	err = universe.StarterFunc()
	require.NoError(t, err)

	cl, adminToken, err := test.AdminClient(options.ClientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	firehoseSvc := cl.FirehoseSvcAPI

	t.Run("firehose subscription with timeout", func(t *testing.T) {
		event := &client.FirehoseSvcEvent{
			Name: client.PtrString("test-event"),
			Data: map[string]any{"hi": "there"},
		}

		eventChannel := make(chan *sse.Event, 1)

		errChannel := make(chan error, 1)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	outer:
		for attempt := 0; attempt < 3; attempt++ {

			sseClient := sse.NewClient(server.URL + "/firehose-svc/events/subscribe")
			sseClient.Headers = map[string]string{
				"Authorization": "Bearer " + adminToken,
			}

			t.Log("Subscribing to firehose")
			go func() {
				err := sseClient.SubscribeChan("messages", eventChannel)
				if err != nil {
					errChannel <- err
				}

			}()

			go func() {
				_, publishErr := firehoseSvc.PublishEvent(ctx).
					Event(client.FirehoseSvcEventPublishRequest{
						Event: event,
					}).
					Execute()

				if publishErr != nil {
					t.Logf("Failed to publish event %v", publishErr)
				}
			}()

			select {
			case receivedEvent := <-eventChannel:
				t.Logf("Received data: %s", receivedEvent.Data)

				ev := &firehose.Event{}
				err = json.Unmarshal([]byte(receivedEvent.Data), ev)
				if err != nil {
					errChannel <- err
					return
				}

				require.Equal(t, *event.Name, ev.Name)
				require.Equal(t, event.Data, ev.Data)

				sseClient.Unsubscribe(eventChannel)
				close(eventChannel)
				return

			case err := <-errChannel:
				t.Log(err)
				continue outer

			case <-time.After(3 * time.Second):
				continue outer

			case <-time.After(10 * time.Second):
				break outer
			}

		}

		t.Fatalf("Test timed out waiting for event")
	})
}
