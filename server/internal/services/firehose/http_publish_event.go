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
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	firehose "github.com/1backend/1backend/server/internal/services/firehose/types"
)

// @ID publishEvent
// @Summary Publish an Event
// @Description Publishes an event to the firehose service after authorization check
// @Tags Firehose Svc
// @Accept json
// @Produce json
// @Param body body firehose.EventPublishRequest true "Event to publish"
// @Success 200 {object} nil "{}"
// @Failure 400 {object} firehose.ErrorResponse "Invalid JSON"
// @Failure 401 {object} firehose.ErrorResponse "Unauthorized"
// @Security BearerAuth
// @Router /firehose-svc/event [post]
func (p *FirehoseService) Publish(w http.ResponseWriter,
	r *http.Request) {

	isAuthRsp, statusCode, err := p.options.PermissionChecker.HasPermission(
		r,
		firehose.PermissionEventPublish,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	req := firehose.EventPublishRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error(
			"Failed to decode request body",
			"error", err,
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	p.publish(req.Event)

	_, err = w.Write([]byte(`{}`))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
