/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package firehoseservice

import (
	"encoding/json"
	"net/http"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	firehose "github.com/openorch/openorch/server/internal/services/firehose/types"
)

// @ID publishEvent
// @Summary Publish an Event
// @Description Publishes an event to the firehose service after authorization check
// @Tags Firehose Svc
// @Accept json
// @Produce json
// @Param event body firehose.EventPublishRequest true "Event to publish"
// @Success 200 {object} nil "{}"
// @Failure 400 {object} firehose.ErrorResponse "Invalid JSON"
// @Failure 401 {object} firehose.ErrorResponse "Unauthorized"
// @Security BearerAuth
// @Router /firehose-svc/event [post]
func (p *FirehoseService) Publish(w http.ResponseWriter,
	r *http.Request) {

	isAuthRsp, _, err := p.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *firehose.PermissionEventPublish.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{
			GrantedSlugs: []string{
				"config-svc",
				"file-svc",
				"prompt-svc",
				"chat-svc",
				"model-svc",
			},
		}).
		Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAuthRsp.GetAuthorized() {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	req := firehose.EventPublishRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	p.publish(req.Event)

	w.Write([]byte(`{}`))
}
