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
	"net/http"

	"github.com/openorch/openorch/sdk/go/logger"
	chat "github.com/openorch/openorch/server/internal/services/chat/types"
)

// @ID events
// @Summary Events
// @Description Events is a dummy endpoint to display documentation about the events that this service emits.
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Success 200 {object} chat.EventMessageAdded
// @Success 200 {object} chat.EventThreadAdded
// @Success 200 {object} chat.EventThreadUpdate
// @Router /chat-svc/events [get]
func (a *ChatService) Events(
	w http.ResponseWriter,
	r *http.Request,
) {
	logger.Info("Events",
		chat.EventMessageAdded{},
		chat.EventThreadAdded{},
		chat.EventThreadUpdate{},
	)
}
