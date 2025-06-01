/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	"context"
	"encoding/json"
	"log/slog"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"
	chat "github.com/1backend/1backend/server/internal/services/chat/types"
)

func (a *ChatService) updateThread(
	callerUserId string,
	thread *chat.Thread,
	req *chat.SaveThreadRequest,
) (*chat.Thread, error) {
	if req.Title != "" {
		thread.Title = req.Title
	}
	if req.TopicIds != nil {
		thread.TopicIds = req.TopicIds
	}
	if req.UserIds != nil {
		thread.UserIds = req.UserIds
	}

	err := a.threadsStore.Query(
		datastore.Equals(datastore.Field("id"), req.Id),
	).Update(thread)

	if err != nil {
		return nil, err
	}

	ev := chat.EventThreadUpdate{
		ThreadId: req.Id,
	}

	var m map[string]interface{}
	js, _ := json.Marshal(ev)
	json.Unmarshal(js, &m)

	_, err = a.options.ClientFactory.Client(client.WithToken(a.token)).
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

	return thread, nil
}
