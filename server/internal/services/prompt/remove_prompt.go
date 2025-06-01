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

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/pkg/errors"

	prompttypes "github.com/1backend/1backend/server/internal/services/prompt/types"
)

func (p *PromptService) removePrompt(promptId string) error {
	logger.Info("Removing prompt",
		slog.String("promptId", promptId),
	)

	err := p.promptsStore.Query(
		datastore.Id(promptId),
	).Delete()

	if err != nil {
		return err
	}

	ev := prompttypes.EventPromptRemoved{
		PromptId: promptId,
	}

	var m map[string]interface{}
	js, _ := json.Marshal(ev)
	json.Unmarshal(js, &m)

	token, err := p.getToken()
	if err != nil {
		return errors.Wrap(err, "failed to get token")
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
		logger.Error("Failed to publish firehose event", slog.Any("error", err))
	}

	return nil
}
