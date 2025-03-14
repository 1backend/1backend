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
	"log/slog"
	"strings"
	"sync"
	"time"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/clients/llamacpp"
	"github.com/openorch/openorch/sdk/go/logger"

	streammanager "github.com/openorch/openorch/server/internal/services/prompt/stream"
	prompttypes "github.com/openorch/openorch/server/internal/services/prompt/types"
)

func (p *PromptService) processLlamaCpp(
	address string,
	currentPrompt *prompttypes.Prompt,
	model *openapi.ModelSvcGetModelResponse,
) error {
	logger.Debug("Starting LLamaCPP stream")

	fullPrompt := currentPrompt.Prompt

	template := ""
	switch {
	case currentPrompt.Parameters != nil && currentPrompt.Parameters.TextToText != nil:
		template = currentPrompt.Parameters.TextToText.Template
	case currentPrompt.EngineParameters != nil && currentPrompt.EngineParameters.LlamaCpp != nil:
		template = currentPrompt.EngineParameters.LlamaCpp.Template
	}

	if template == "" {
		modelRsp, _, err := p.clientFactory.Client(sdk.WithToken(p.token)).
			ModelSvcAPI.GetModel(context.Background(), currentPrompt.ModelId).
			Execute()
		if err != nil {
			return err
		}

		if modelRsp.Model.PromptTemplate != nil {
			template = *modelRsp.Model.PromptTemplate
		}
	}

	if template != "" {
		fullPrompt = strings.Replace(
			template,
			"{prompt}",
			currentPrompt.Prompt,
			-1,
		)
	}

	var llamaCppClient llamacpp.ClientI
	if p.llamaCppCLient != nil {
		llamaCppClient = p.llamaCppCLient
	} else {
		llamaCppClient = &llamacpp.Client{
			LLamaCppAddress: address,
		}
	}

	start := time.Now()
	var responseCount int
	var mu sync.Mutex

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				mu.Lock()
				logger.Debug(
					"LLamaCPP is streaming",
					slog.String("promptId", currentPrompt.Id),
					slog.Float64(
						"responsesPerSecond",
						float64(responseCount/int(time.Since(start).Seconds())),
					),
					slog.Int("totalResponses", responseCount),
				)
				mu.Unlock()
			case <-done:
				return
			}
		}
	}()

	err := llamaCppClient.PostCompletionsStreamed(llamacpp.PostCompletionsRequest{
		Prompt:    fullPrompt,
		Stream:    true,
		MaxTokens: 1000000,
	}, func(resp *llamacpp.CompletionResponse) {
		mu.Lock()
		responseCount++
		mu.Unlock()

		if len(resp.Choices) > 0 && resp.Choices[0].FinishReason == "stop" {
			defer func() {
				done <- true
			}()

			messageId := sdk.Id("msg")

			_, _, err := p.clientFactory.Client(sdk.WithToken(p.token)).
				ChatSvcAPI.AddMessage(context.Background(), currentPrompt.ThreadId).
				Body(
					openapi.ChatSvcAddMessageRequest{
						Message: &openapi.ChatSvcMessage{
							Id:       messageId,
							ThreadId: currentPrompt.ThreadId,
							Text: openapi.PtrString(
								p.streamManager.ConcatHistoryText(currentPrompt.ThreadId),
							),
							Meta: map[string]interface{}{
								"modelId":    model.Model.Id,
								"platformId": model.Platform.Id,
							},
						},
					},
				).
				Execute()
			if err != nil {
				logger.Error("Error when saving chat message after broadcast",
					slog.String("error", err.Error()))
				return
			}

			p.streamManager.DeleteHistory(currentPrompt.ThreadId)

			p.streamManager.Broadcast(currentPrompt.ThreadId, &streammanager.Chunk{
				Text:      resp.Choices[0].Text,
				Type:      streammanager.ChunkTypeDone,
				MessageId: messageId,
			})
		} else {
			p.streamManager.Broadcast(currentPrompt.ThreadId, &streammanager.Chunk{
				Text: resp.Choices[0].Text,
				Type: streammanager.ChunkTypeProgress,
			})
		}
	})

	return err
}

func errToString(err error) string {
	if err != nil {
		return err.Error()
	}

	return ""
}
