/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxyservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/pkg/errors"

	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
)

// @ID deleteRoutes
// @Summary Delete Routes
// @Description Delete specific routes by their IDs.
// @Tags Proxy Svc
// @Accept json
// @Produce json
// @Param body body proxy.DeleteRoutesRequest true "Delete Routes Request"
// @Success 200 {object} proxy.DeleteRoutesResponse "Routes deleted successfully"
// @Failure 400 {object} proxy.ErrorResponse "Invalid JSON or missing IDs"
// @Failure 401 {object} proxy.ErrorResponse "Unauthorized"
// @Failure 500 {object} proxy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /proxy-svc/routes [delete]
func (cs *ProxyService) DeleteRoutes(w http.ResponseWriter, r *http.Request) {
	isAuthRsp, statusCode, err := cs.options.PermissionChecker.HasPermission(
		r,
		proxy.PermissionRouteEdit, // Reusing edit permission
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	req := &proxy.DeleteRoutesRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error(
			"Failed to decode delete request body",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	err = cs.deleteRoutes(req)
	if err != nil {
		logger.Error(
			"Failed to delete routes",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, &proxy.DeleteRoutesResponse{})
}

func (cs *ProxyService) deleteRoutes(req *proxy.DeleteRoutesRequest) error {
	if len(req.Ids) == 0 {
		return errors.New("no route IDs provided")
	}

	ids := []any{}
	for _, v := range req.Ids {
		ids = append(ids, v)
	}

	err := cs.routeStore.Query(
		datastore.IsInList([]string{"id"}, ids...),
	).Delete()
	if err != nil {
		return errors.Wrap(err, "failed to delete routes from store")
	}

	cs.invalidateRouteCache()

	return nil
}
