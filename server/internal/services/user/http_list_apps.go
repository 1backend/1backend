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

// @ID listApps
// @Summary List Apps
// @Description List apps. Role, user ID or contact ID must be specified.
// @Description
// @Description Requires the `user-svc:app:view` permission, which by default all users have.
// @Description Caller can only list apps of roles they own (unless they are an admin).
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ListAppsRequest true "List Apps Request"
// @Success 200 {object} user.ListAppsResponse "Apps listed successfully"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/apps [post]
func (s *UserService) ListApps(w http.ResponseWriter, r *http.Request) {
	req := &user.ListAppsRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	if len(req.Ids) == 0 && len(req.Hosts) == 0 {
		// Only admins can list all apps
		_, hasPermission, _, err := s.hasPermission(r, user.PermissionAppMultiView)
		if err != nil {
			logger.Error(
				"Failed to check permission",
				slog.Any("error", err),
			)
			endpoint.InternalServerError(w)
			return
		}
		if !hasPermission {
			endpoint.Unauthorized(w)
			return
		}
	}

	apps, err := s.listApps(req)
	if err != nil {
		logger.Error(
			"Failed to list apps",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	rsp := &user.ListAppsResponse{
		Apps: apps,
	}
	endpoint.WriteJSON(w, http.StatusOK, rsp)
}

func (s *UserService) listApps(
	req *user.ListAppsRequest,
) ([]user.App, error) {
	filters := []datastore.Filter{}

	if len(req.Ids) > 0 {
		ids := []any{}
		for _, id := range req.Ids {
			ids = append(ids, id)
		}
		filters = append(filters, datastore.IsInList(
			[]string{"id"},
			ids...,
		))
	}
	if len(req.Hosts) > 0 {
		hosts := []any{}
		for _, host := range req.Hosts {
			hosts = append(hosts, host)
		}
		filters = append(filters, datastore.IsInList(
			[]string{"host"},
			hosts...,
		))
	}

	appIs, err := s.appStore.Query(filters...).Find()
	if err != nil {
		return nil, err
	}

	apps := make([]user.App, 0, len(appIs))
	for _, ai := range appIs {
		app, ok := ai.(*user.App)
		if !ok {
			continue
		}
		apps = append(apps, *app)
	}

	return apps, nil
}
