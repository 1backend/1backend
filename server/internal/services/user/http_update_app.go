/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID updateApp
// @Summary Update App Host
// @Description Change the hostname of an existing app.
// @Description Requires the `user-svc:app:edit` permission.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.UpdateAppRequest true "Update App Request"
// @Success 200 {object} user.UpdateAppResponse "App updated successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid Request"
// @Failure 404 {object} user.ErrorResponse "App Not Found"
// @Security BearerAuth
// @Router /user-svc/app [put]
func (s *UserService) UpdateApp(w http.ResponseWriter, r *http.Request) {
	req := &user.UpdateAppRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	_, hasPermission, _, err := s.hasPermission(r, user.PermissionAppEdit)
	if err != nil || !hasPermission {
		endpoint.Unauthorized(w)
		return
	}

	err = s.updateAppHost(req.Id, req.NewHost)
	if err != nil {
		logger.Error("Failed to update app host", slog.String("id", req.Id), slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, &user.UpdateAppResponse{})
}

func (s *UserService) updateAppHost(id string, newHost string) error {
	return s.appStore.Query(
		datastore.Id(id),
	).UpdateFields(map[string]interface{}{
		"host": newHost,
	})
}
