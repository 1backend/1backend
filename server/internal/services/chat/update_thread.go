/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package chatservice

import (
	"context"
	"encoding/json"
	"log/slog"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/logger"
	chattypes "github.com/openorch/openorch/server/internal/services/chat/types"
)

func (a *ChatService) updateThread(
	chatThread *chattypes.Thread,
) (*chattypes.Thread, error) {
	err := a.threadsStore.Query(
		datastore.Equals(datastore.Field("id"), chatThread.Id),
	).Update(chatThread)

	if err != nil {
		return nil, err
	}

	ev := chattypes.EventThreadUpdate{
		ThreadId: chatThread.Id,
	}

	var m map[string]interface{}
	js, _ := json.Marshal(ev)
	json.Unmarshal(js, &m)

	_, err = a.clientFactory.Client(sdk.WithToken(a.token)).
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

	return chatThread, nil
}
