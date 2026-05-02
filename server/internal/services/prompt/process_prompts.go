/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package promptservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/flusflas/dipper"
	"github.com/pkg/errors"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"

	modelservice "github.com/1backend/1backend/server/internal/services/model"
	modeltypes "github.com/1backend/1backend/server/internal/services/model/types"
	prompttypes "github.com/1backend/1backend/server/internal/services/prompt/types"
)

var TimeNow = time.Now

const (
	maxRetries               = 5
	BaseDelay                = 1 * time.Second
	promptTimeout            = 1 * time.Minute
	promptActivePollInterval = 2 * time.Second
	promptIdlePollInterval   = 1 * time.Minute
)

type promptProcessState int

const (
	promptProcessIdle promptProcessState = iota
	promptProcessActive
)

// a blocking method, call it in a goroutine
func (p *PromptService) processPrompts() {
	timer := time.NewTimer(0)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
		case <-p.trigger:
		}

		stopPromptTimer(timer)

		state, waitFor, err := p.processNextPrompt()
		if err != nil {
			logger.Error("Error processing prompt",
				slog.String("error", err.Error()),
			)
			state = promptProcessActive
		}

		nextPoll := promptIdlePollInterval
		if state == promptProcessActive {
			nextPoll = promptActivePollInterval
		}
		if waitFor > 0 {
			nextPoll = waitFor
			if nextPoll > promptIdlePollInterval {
				nextPoll = promptIdlePollInterval
			}
		}
		timer.Reset(nextPoll)
	}
}

func stopPromptTimer(timer *time.Timer) {
	if !timer.Stop() {
		select {
		case <-timer.C:
		default:
		}
	}
}

func (p *PromptService) processNextPrompt() (promptProcessState, time.Duration, error) {
	p.runMutex.Lock()
	defer p.runMutex.Unlock()

	activePrompts, err := p.promptsStore.Query(
		datastore.IsInList(
			datastore.Field("status"),
			prompttypes.PromptStatusRunning,
			prompttypes.PromptStatusScheduled,
			prompttypes.PromptStatusErrored,
		),
	).
		OrderBy(datastore.OrderByField("createdAt", false)).
		Find()
	if err != nil {
		return promptProcessActive, 0, err
	}

	hasRunning := false
	runningPromptId := ""
	queuedPrompts := make([]datastore.Row, 0, len(activePrompts))

	for _, promptI := range activePrompts {
		prompt := promptI.(*prompttypes.Prompt)

		if prompt.Status == prompttypes.PromptStatusScheduled ||
			prompt.Status == prompttypes.PromptStatusErrored {
			queuedPrompts = append(queuedPrompts, promptI)
			continue
		}

		if prompt.Status != prompttypes.PromptStatusRunning {
			continue
		}

		if prompt.LastRun.Before(TimeNow().Add(-promptTimeout)) {
			logger.Info("Setting prompt as timed out",
				slog.String("promptId", prompt.Id),
			)

			prompt.Status = prompttypes.PromptStatusErrored
			prompt.Error = "timed out"
			err = p.promptsStore.Query(
				datastore.Id(prompt.Id),
			).Update(prompt)
			if err != nil {
				return promptProcessActive, 0, err
			}
			queuedPrompts = append(queuedPrompts, prompt)
			continue
		}
		hasRunning = true
		runningPromptId = prompt.Id
	}

	if hasRunning {
		logger.Debug("Prompt is already running",
			slog.String("promptId", runningPromptId),
		)
		return promptProcessActive, 0, nil
	}

	currentPrompt, hasQueuedPrompt, nextDue := selectPromptFromRows(queuedPrompts)
	if currentPrompt == nil {
		if hasQueuedPrompt {
			return promptProcessActive, nextDue, nil
		}
		return promptProcessIdle, 0, nil
	}

	return promptProcessActive, 0, p.runPrompt(currentPrompt)
}

func (p *PromptService) runPrompt(currentPrompt *prompttypes.Prompt) error {
	if p.processPromptFunc != nil {
		return p.processPromptFunc(currentPrompt)
	}
	return p.processPrompt(currentPrompt)
}

func (p *PromptService) processPrompt(
	currentPrompt *prompttypes.Prompt,
) (err error) {
	updateCurr := func() {
		logger.Info("Prompt finished",
			slog.String("promptId", currentPrompt.Id),
			slog.String("status", string(currentPrompt.Status)),
			slog.Any("error", err),
		)

		err = p.promptsStore.Query(
			datastore.Id(currentPrompt.Id),
		).Update(currentPrompt)
		if err != nil {
			logger.Error("Error updating prompt",
				slog.String("promptId", currentPrompt.Id),
				slog.String("error", err.Error()),
			)
		}
	}

	token, err := p.getToken()
	if err != nil {
		return errors.Wrap(err, "failed to get token")
	}

	defer func() {
		if err != nil {
			currentPrompt.Error = err.Error()
			currentPrompt.Status = prompttypes.PromptStatusErrored
			updateCurr()
		} else {
			currentPrompt.Status = prompttypes.PromptStatusCompleted
			updateCurr()
		}

		ev := prompttypes.EventPromptProcessingFinished{
			PromptId: currentPrompt.Id,
			Error:    errToString(err),
		}

		var m map[string]interface{}
		js, _ := json.Marshal(ev)
		err = json.Unmarshal(js, &m)
		if err != nil {
			logger.Error("Failed to unmarshal event",
				slog.Any("error", err),
			)
		}

		_, err = p.options.ClientFactory.Client(client.WithToken(token)).
			FirehoseSvcAPI.PublishEvent(context.Background()).
			Event(openapi.FirehoseSvcEventPublishRequest{
				Event: &openapi.FirehoseSvcEvent{
					Name: openapi.PtrString(ev.Name()),
					Data: m,
				},
			}).
			Execute()
		if err != nil {
			logger.Error(
				"Failed to publish firehose event",
				slog.Any("error", err),
			)
		}
	}()

	logger.Info("Picking up prompt from queue",
		slog.String("promptId", currentPrompt.Id),
	)

	currentPrompt.LastRun = time.Now()
	currentPrompt.Error = ""
	currentPrompt.Status = prompttypes.PromptStatusRunning
	currentPrompt.RunCount++

	err = p.promptsStore.Upsert(currentPrompt)
	if err != nil {
		return errors.Wrap(err, "error updating currently running prompt")
	}

	ev := prompttypes.EventPromptProcessingStarted{
		PromptId: currentPrompt.Id,
	}

	var m map[string]interface{}
	js, _ := json.Marshal(ev)
	json.Unmarshal(js, &m)

	_, err = p.options.ClientFactory.Client(client.WithToken(token)).
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

	modelId := currentPrompt.ModelId
	if modelId == "" {
		readConfigRsp, _, err := p.options.ClientFactory.Client(client.WithToken(token)).
			ConfigSvcAPI.ListConfigs(context.Background()).
			Body(openapi.ConfigSvcListConfigsRequest{
				AppHost: sdk.DefaultAppHost,
				Ids:     []string{"modelSvc"},
			}).
			Execute()
		if err != nil {
			return err
		}

		modelIdI := dipper.Get(readConfigRsp.Configs["modelSvc"].Data, "currentModelId")
		var ok bool
		modelId, ok = modelIdI.(string)
		if !ok {
			modelId = modelservice.DefaultModelId
		}
	}
	if currentPrompt.ModelId == "" {
		currentPrompt.ModelId = modelId
	}

	getModelRsp, _, err := p.options.ClientFactory.Client(client.WithToken(token)).
		ModelSvcAPI.GetModel(context.Background(), modelId).
		Execute()
	if err != nil {
		return err
	}
	_, _, err = p.options.ClientFactory.Client(client.WithToken(token)).
		ChatSvcAPI.SaveMessage(context.Background(), currentPrompt.ThreadId).
		Body(openapi.ChatSvcSaveMessageRequest{
			// not a fan of taking the prompt id but at least it makes this idempotent
			// in case prompts get retried over and over again
			Id:       &currentPrompt.Id,
			ThreadId: &currentPrompt.ThreadId,
			UserId:   openapi.PtrString(currentPrompt.UserId),
			Text:     openapi.PtrString(currentPrompt.Prompt),
			Meta: map[string]interface{}{
				"modelId":    getModelRsp.Model.Id,
				"platformId": getModelRsp.Platform.Id,
			},
		}).
		Execute()
	if err != nil {
		return err
	}

	statusRsp, _, err := p.options.ClientFactory.Client(client.WithToken(token)).
		ModelSvcAPI.GetModelStatus(context.Background(), modelId).
		Execute()
	if err != nil {
		return err
	}

	stat := statusRsp.Status
	if !stat.Running {
		return fmt.Errorf("model '%v' is not running", modelId)
	}
	if stat.Address == "" {
		return errors.Wrap(err, "missing model address")
	}
	if !strings.HasPrefix(stat.Address, "http") {
		stat.Address = "http://" + stat.Address
	}

	err = p.processPlatform(stat.Address, currentPrompt, getModelRsp)

	logger.Debug("Finished streaming prompt",
		slog.String("error", fmt.Sprintf("%v", err)),
	)
	if err != nil {
		return errors.Wrap(err, "error streaming llm")
	}

	return nil
}

func (p *PromptService) processPlatform(
	address string,
	currentPrompt *prompttypes.Prompt,
	model *openapi.ModelSvcGetModelResponse,
) error {

	switch *model.Platform.Id {
	case modeltypes.PlatformLlamaCpp.Id:
		return p.processLlamaCpp(address, currentPrompt, model)
	case modeltypes.PlatformStableDiffusion.Id:
		return p.processStableDiffusion(address, currentPrompt, model)
	}

	return fmt.Errorf("cannot find platform %v", model.Platform.Id)
}
