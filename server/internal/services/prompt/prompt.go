/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package promptservice

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/logger"
	"github.com/pkg/errors"

	chat "github.com/openorch/openorch/server/internal/services/chat/types"

	apptypes "github.com/openorch/openorch/server/internal/services/chat/types"
	streammanager "github.com/openorch/openorch/server/internal/services/prompt/stream"
	prompttypes "github.com/openorch/openorch/server/internal/services/prompt/types"
)

const maxThreadTitle = 100

func (p *PromptService) prompt(
	ctx context.Context,
	promptReq *prompttypes.PromptRequest,
	userId string,
) (*prompttypes.PromptResponse, error) {
	prompt := &prompttypes.Prompt{
		Id:               promptReq.Id,
		Prompt:           promptReq.Prompt,
		Sync:             promptReq.Sync,
		ThreadId:         promptReq.ThreadId,
		ModelId:          promptReq.ModelId,
		MaxRetries:       promptReq.MaxRetries,
		Parameters:       promptReq.Parameters,
		EngineParameters: promptReq.EngineParameters,
	}

	prompt.Status = prompttypes.PromptStatusScheduled
	now := TimeNow()
	prompt.CreatedAt = now
	prompt.UpdatedAt = now
	prompt.UserId = userId

	if prompt.Id == "" {
		prompt.Id = sdk.Id("prom")
	}

	if prompt.ThreadId == "" {
		prompt.ThreadId = prompt.Id
	}

	err := p.promptsStore.Create(prompt)
	if err != nil {
		return nil, err
	}

	logger.Info("Created prompt",
		slog.String("promptId", prompt.Id),
	)

	threadId := prompt.ThreadId

	getThreadRsp, _, err := p.clientFactory.Client(sdk.WithToken(p.token)).
		ChatSvcAPI.GetThread(ctx, threadId).
		Execute()
	if err != nil {
		return nil, err
	}

	if !*getThreadRsp.Exists {
		logger.Info("Creating thread", slog.String("threadId", threadId))

		// threads can be created when a message is sent
		now := time.Now()

		thread := &apptypes.Thread{
			Id:        threadId,
			UserIds:   []string{userId},
			CreatedAt: now,
			UpdatedAt: now,
		}

		if thread.Title == "" {
			if len(prompt.Prompt) > maxThreadTitle {
				thread.Title = prompt.Prompt[:maxThreadTitle]
			} else {
				thread.Title = prompt.Prompt
			}
		}

		_, _, err := p.clientFactory.Client(sdk.WithToken(p.token)).
			ChatSvcAPI.AddThread(ctx).
			Body(openapi.ChatSvcAddThreadRequest{
				Thread: &openapi.ChatSvcThread{
					Id:      thread.Id,
					Title:   openapi.PtrString(thread.Title),
					UserIds: thread.UserIds,
					CreatedAt: openapi.PtrString(
						thread.CreatedAt.Format(time.RFC3339Nano),
					),
					UpdatedAt: openapi.PtrString(
						thread.UpdatedAt.Format(time.RFC3339Nano),
					),
				},
			}).
			Execute()
		if err != nil {
			return nil, err
		}
	}

	ev := prompttypes.EventPromptAdded{
		PromptId: prompt.Id,
	}

	var m map[string]interface{}
	js, _ := json.Marshal(ev)
	json.Unmarshal(js, &m)

	_, err = p.clientFactory.Client(sdk.WithToken(p.token)).
		FirehoseSvcAPI.PublishEvent(context.Background()).
		Event(openapi.FirehoseSvcEventPublishRequest{
			Event: &openapi.FirehoseSvcEvent{
				Name: openapi.PtrString(ev.Name()),
				Data: m,
			},
		}).
		Execute()
	if err != nil {
		logger.Error("Failed to publish firehose event", slog.Any("error", err))
	}

	go p.triggerPromptProcessing()

	rsp := &prompttypes.PromptResponse{
		Prompt: prompt,
	}

	if prompt.Sync {
		subscriber := make(chan *streammanager.Chunk)
		p.streamManager.Subscribe(threadId, subscriber)

		go func() {
			<-ctx.Done()
			p.streamManager.Unsubscribe(threadId, subscriber)
		}()

		for resp := range subscriber {
			if resp.Type == streammanager.ChunkTypeDone {
				r, _, err := p.clientFactory.Client(sdk.WithToken(p.token)).
					ChatSvcAPI.
					GetMessage(ctx, resp.MessageId).
					Execute()
				if err != nil {
					return nil, errors.Wrap(err, "error reading message")
				}

				m := r.Message

				// @todo should use openapi type here but there are issues
				// with the generation
				rsp.ResponseMessage = &chat.Message{
					Id:       m.Id,
					ThreadId: m.ThreadId,
					FileIds:  m.FileIds,
				}
				if m.Text != nil {
					rsp.ResponseMessage.Text = *m.Text
				}
				if m.UserId != nil {
					rsp.ResponseMessage.UserId = *m.UserId
				}

				return rsp, nil
			}
		}
	}

	return rsp, nil
}

func (p *PromptService) triggerPromptProcessing() {
	select {
	case p.trigger <- true:
		logger.Debug("Prompt trigger signal sent")
	default:
		logger.Debug("Prompt trigger signal skipped, already pending")
	}
}
