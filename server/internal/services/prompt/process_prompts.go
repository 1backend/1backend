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
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/flusflas/dipper"
	"github.com/pkg/errors"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/logger"

	modelservice "github.com/openorch/openorch/server/internal/services/model"
	modeltypes "github.com/openorch/openorch/server/internal/services/model/types"
	prompttypes "github.com/openorch/openorch/server/internal/services/prompt/types"
)

var TimeNow = time.Now

const (
	maxRetries    = 5
	BaseDelay     = 1 * time.Second
	promptTimeout = 1 * time.Minute
)

// a blocking method, call it in a goroutine
func (p *PromptService) processPrompts() {
	ticker := time.NewTicker(2000 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		case <-p.trigger:
		}

		err := p.processNextPrompt()
		if err != nil {
			logger.Error("Error processing prompt",
				slog.String("error", err.Error()),
			)
		}
	}
}

func (p *PromptService) processNextPrompt() error {
	p.runMutex.Lock()
	defer p.runMutex.Unlock()

	runningPrompts, err := p.promptsStore.Query(
		datastore.Equals(
			datastore.Field("status"),
			prompttypes.PromptStatusRunning,
		),
	).Find()
	if err != nil {
		return err
	}

	hasRunning := false
	runningPromptId := ""
	for _, runningPromptI := range runningPrompts {
		runningPrompt := runningPromptI.(*prompttypes.Prompt)

		if runningPrompt.LastRun.Before(time.Now().Add(-promptTimeout)) {
			logger.Info("Setting prompt as timed out",
				slog.String("promptId", runningPrompt.Id),
			)

			runningPrompt.Status = prompttypes.PromptStatusErrored
			runningPrompt.Error = "timed out"
			err = p.promptsStore.Query(
				datastore.Id(runningPrompt.Id),
			).Update(runningPrompt)
			if err != nil {
				return err
			}
			continue
		}
		hasRunning = true
		runningPromptId = runningPrompt.Id
	}

	if hasRunning {
		logger.Debug("Prompt is already running",
			slog.String("promptId", runningPromptId),
		)
		return nil
	}

	currentPrompt, err := SelectPrompt(p.promptsStore)
	if err != nil {
		return err
	}
	if currentPrompt == nil {
		return nil
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

	modelId := currentPrompt.ModelId
	if modelId == "" {
		getConfigRsp, _, err := p.clientFactory.Client(sdk.WithToken(p.token)).
			ConfigSvcAPI.GetConfig(context.Background()).
			Execute()
		if err != nil {
			return err
		}

		modelIdI := dipper.Get(getConfigRsp.Config.Data, "model-svc.currentModelId")
		var ok bool
		modelId, ok = modelIdI.(string)
		if !ok {
			modelId = modelservice.DefaultModelId
		}
	}
	if currentPrompt.ModelId == "" {
		currentPrompt.ModelId = modelId
	}

	getModelRsp, _, err := p.clientFactory.Client(sdk.WithToken(p.token)).
		ModelSvcAPI.GetModel(context.Background(), modelId).
		Execute()
	if err != nil {
		return err
	}
	_, _, err = p.clientFactory.Client(sdk.WithToken(p.token)).
		ChatSvcAPI.AddMessage(context.Background(), currentPrompt.ThreadId).
		Body(openapi.ChatSvcAddMessageRequest{
			Message: &openapi.ChatSvcMessage{
				// not a fan of taking the prompt id but at least it makes this idempotent
				// in case prompts get retried over and over again
				Id:       currentPrompt.Id,
				ThreadId: currentPrompt.ThreadId,
				UserId:   openapi.PtrString(currentPrompt.UserId),
				Text:     openapi.PtrString(currentPrompt.Prompt),
				CreatedAt: openapi.PtrString(
					time.Now().Format(time.RFC3339Nano),
				),
				Meta: map[string]interface{}{
					"modelId":    getModelRsp.Model.Id,
					"platformId": getModelRsp.Platform.Id,
				},
			},
		}).
		Execute()
	if err != nil {
		return err
	}

	statusRsp, _, err := p.clientFactory.Client(sdk.WithToken(p.token)).
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
