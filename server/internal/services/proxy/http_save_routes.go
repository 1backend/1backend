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

// @ID saveRoutes
// @Summary Save Routes
// @Description Save routes that the edge proxy will use to route requests.
// @Tags Proxy Svc
// @Accept json
// @Produce json
// @Param body body proxy.SaveRoutesRequest true "Save Routes Request"
// @Success 200 {object} proxy.SaveRoutesResponse "Routes saved successfully"
// @Failure 400 {object} proxy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} proxy.ErrorResponse "Unauthorized"
// @Failure 500 {object} proxy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /proxy-svc/routes [put]
func (cs *ProxyService) SaveRoutes(w http.ResponseWriter, r *http.Request) {
	isAuthRsp, statusCode, err := cs.options.PermissionChecker.HasPermission(
		r,
		proxy.PermissionRouteEdit,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &proxy.SaveRoutesRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error(
			"Failed to decode request body",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	routes, err := cs.saveRoutes(req)
	if err != nil {
		logger.Error(
			"Failed to save routes",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, &proxy.SaveRoutesResponse{
		Routes: routes,
	})
}

func (cs *ProxyService) saveRoutes(req *proxy.SaveRoutesRequest) ([]proxy.Route, error) {
	if len(req.Routes) == 0 {
		return nil, errors.New("no routes provided")
	}

	rows := make([]datastore.Row, 0, len(req.Routes))
	routes := make([]proxy.Route, 0, len(req.Routes))

	for _, route := range req.Routes {
		if route.Id == "" {
			return nil, errors.New("route ID is required")
		}
		if route.Target == "" {
			return nil, errors.New("route target is required")
		}

		r := proxy.Route{
			Id:     route.Id,
			Target: route.Target,
		}
		routes = append(routes, r)
		rows = append(rows, &r)
	}

	err := cs.routeStore.UpsertMany(rows)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save routes")
	}

	return routes, nil
}
