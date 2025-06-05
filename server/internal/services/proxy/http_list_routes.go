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

// @ID listRoutes
// @Summary List Routes
// @Description List routes that the edge proxy will use to route requests.
// @Tags Proxy Svc
// @Accept json
// @Produce json
// @Param body body proxy.ListRoutesRequest false "List Routes Request"
// @Success 200 {object} proxy.ListRoutesResponse "Routes listd successfully"
// @Failure 400 {object} proxy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} proxy.ErrorResponse "Unauthorized"
// @Failure 500 {object} proxy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /proxy-svc/routes [post]
func (cs *ProxyService) ListRoutes(w http.ResponseWriter, r *http.Request) {
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

	req := &proxy.ListRoutesRequest{}

	if r.ContentLength > 0 {
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
	}

	routes, err := cs.listRoutes(req)
	if err != nil {
		logger.Error(
			"Failed to list routes",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
	}

	endpoint.WriteJSON(w, http.StatusOK, &proxy.ListRoutesResponse{
		Routes: routes,
	})
}

func (cs *ProxyService) listRoutes(req *proxy.ListRoutesRequest) ([]proxy.Route, error) {
	filters := []datastore.Filter{}

	if len(req.Ids) > 0 {
		ids := []any{}
		for _, id := range req.Ids {
			ids = append(ids, id)
		}
		filters = append(filters, datastore.IsInList([]string{"id"}, ids...))
	}

	routeIs, err := cs.routeStore.Query(filters...).
		Find()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query routes")
	}

	routes := make([]proxy.Route, 0, len(routeIs))
	for _, routeI := range routeIs {
		route, ok := routeI.(*proxy.Route)
		if !ok {
			return nil, errors.Errorf("expected route type, got %T", routeI)
		}
		routes = append(routes, *route)
	}

	return routes, nil
}
