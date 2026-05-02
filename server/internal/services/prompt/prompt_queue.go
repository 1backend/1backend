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
	"log/slog"

	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/1backend/1backend/sdk/go/pubsub"
	prompttypes "github.com/1backend/1backend/server/internal/services/prompt/types"
)

const (
	promptQueueTopic        = "prompt-svc.prompts"
	promptQueueSubscriberID = "prompt-svc-worker"
)

func (p *PromptService) subscribePromptQueue() {
	if p.options.PubSub == nil {
		return
	}

	sub, err := p.options.PubSub.Subscribe(
		context.Background(),
		promptQueueSubscriberID,
		promptQueueTopic,
		pubsub.WithBackfillSince(TimeNow().UTC()),
	)
	if err != nil {
		logger.Warn("Failed to subscribe to prompt queue notifications",
			slog.Any("error", err),
		)
		return
	}

	go p.forwardPromptQueueNotifications(sub)
}

func (p *PromptService) forwardPromptQueueNotifications(sub pubsub.Subscription) {
	for msg := range sub.Chan() {
		if err := sub.Ack(context.Background(), msg.ID); err != nil {
			logger.Warn("Failed to ack prompt queue notification",
				slog.String("messageId", msg.ID),
				slog.Any("error", err),
			)
		}

		p.triggerPromptProcessing()
	}
}

func (p *PromptService) notifyPromptQueued(ctx context.Context, promptId string) {
	if p.options.PubSub != nil {
		payload, _ := json.Marshal(prompttypes.EventPromptAdded{
			PromptId: promptId,
		})
		if _, err := p.options.PubSub.Publish(ctx, promptQueueTopic, payload); err != nil {
			logger.Error("Failed to publish prompt queue notification",
				slog.String("promptId", promptId),
				slog.Any("error", err),
			)
		}
	}

	p.triggerPromptProcessing()
}
