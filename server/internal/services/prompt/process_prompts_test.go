package promptservice

import (
	"context"
	"testing"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore/localstore"
	"github.com/1backend/1backend/sdk/go/pubsub/localpubsub"
	prompttypes "github.com/1backend/1backend/server/internal/services/prompt/types"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/stretchr/testify/require"
)

func TestPromptQueueNotificationTriggersProcessing(t *testing.T) {
	ps, err := localpubsub.NewLocalPubSub("")
	require.NoError(t, err)
	defer ps.Close()

	service := &PromptService{
		options: &universe.Options{PubSub: ps},
		trigger: make(chan bool, 1),
	}
	service.subscribePromptQueue()

	_, err = ps.Publish(context.Background(), promptQueueTopic, []byte(`{"promptId":"prom-test"}`))
	require.NoError(t, err)

	select {
	case <-service.trigger:
	case <-time.After(time.Second):
		t.Fatal("prompt queue notification did not trigger prompt processing")
	}
}

func TestProcessNextPromptDispatchesReadyPrompt(t *testing.T) {
	fixedTime := time.Date(2023, 6, 1, 12, 0, 0, 0, time.UTC)
	originalTimeNow := TimeNow
	TimeNow = func() time.Time {
		return fixedTime
	}
	defer func() {
		TimeNow = originalTimeNow
	}()

	store, err := localstore.NewLocalStore(&prompttypes.Prompt{}, "")
	require.NoError(t, err)

	prompt := &prompttypes.Prompt{
		Id:        "prom-ready",
		Status:    prompttypes.PromptStatusScheduled,
		CreatedAt: fixedTime,
	}
	require.NoError(t, store.Create(prompt))

	var processedPrompt *prompttypes.Prompt
	service := &PromptService{
		promptsStore: store,
		processPromptFunc: func(p *prompttypes.Prompt) error {
			processedPrompt = p
			return nil
		},
	}

	state, waitFor, err := service.processNextPrompt()
	require.NoError(t, err)
	require.Equal(t, promptProcessActive, state)
	require.Zero(t, waitFor)
	require.NotNil(t, processedPrompt)
	require.Equal(t, prompt.Id, processedPrompt.Id)
}

func TestProcessNextPromptBacksOffQueuedPrompt(t *testing.T) {
	fixedTime := time.Date(2023, 6, 1, 12, 0, 0, 0, time.UTC)
	originalTimeNow := TimeNow
	TimeNow = func() time.Time {
		return fixedTime
	}
	defer func() {
		TimeNow = originalTimeNow
	}()

	store, err := localstore.NewLocalStore(&prompttypes.Prompt{}, "")
	require.NoError(t, err)

	require.NoError(t, store.Create(&prompttypes.Prompt{
		Id:        "prom-waiting",
		Status:    prompttypes.PromptStatusScheduled,
		RunCount:  1,
		LastRun:   fixedTime.Add(-BaseDelay / 2),
		CreatedAt: fixedTime,
	}))

	service := &PromptService{promptsStore: store}

	state, waitFor, err := service.processNextPrompt()
	require.NoError(t, err)
	require.Equal(t, promptProcessActive, state)
	require.Equal(t, BaseDelay/2, waitFor)
}

func TestProcessNextPromptIdleWhenNoQueuedPrompt(t *testing.T) {
	store, err := localstore.NewLocalStore(&prompttypes.Prompt{}, "")
	require.NoError(t, err)

	service := &PromptService{promptsStore: store}

	state, waitFor, err := service.processNextPrompt()
	require.NoError(t, err)
	require.Equal(t, promptProcessIdle, state)
	require.Zero(t, waitFor)
}
