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
	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
	"github.com/pkg/errors"
)

const defaultRedirectStatusCode = http.StatusPermanentRedirect

// @ID saveRedirects
// @Summary Save Redirects
// @Description Save redirects that the edge proxy will apply before routing requests.
// @Tags Proxy Svc
// @Accept json
// @Produce json
// @Param body body proxy.SaveRedirectsRequest true "Save Redirects Request"
// @Success 200 {object} proxy.SaveRedirectsResponse "Redirects saved successfully"
// @Failure 400 {object} proxy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} proxy.ErrorResponse "Unauthorized"
// @Failure 500 {object} proxy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /proxy-svc/redirects [put]
func (cs *ProxyService) SaveRedirects(w http.ResponseWriter, r *http.Request) {
	isAuthRsp, statusCode, err := cs.options.PermissionChecker.HasPermission(
		r,
		proxy.PermissionRouteEdit,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	req := &proxy.SaveRedirectsRequest{}
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

	redirects, err := cs.saveRedirects(req)
	if err != nil {
		logger.Error(
			"Failed to save redirects",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, &proxy.SaveRedirectsResponse{
		Redirects: redirects,
	})
}

func (cs *ProxyService) saveRedirects(req *proxy.SaveRedirectsRequest) ([]proxy.Redirect, error) {
	if len(req.Redirects) == 0 {
		return nil, errors.New("no redirects provided")
	}

	rows := make([]datastore.Row, 0, len(req.Redirects))
	redirects := make([]proxy.Redirect, 0, len(req.Redirects))

	for _, redirect := range req.Redirects {
		if redirect.Id == "" {
			return nil, errors.New("redirect ID is required")
		}
		if redirect.Target == "" {
			return nil, errors.New("redirect target is required")
		}

		statusCode, err := normalizeRedirectStatusCode(redirect.StatusCode)
		if err != nil {
			return nil, err
		}

		r := proxy.Redirect{
			Id:         redirect.Id,
			Target:     redirect.Target,
			StatusCode: statusCode,
		}
		redirects = append(redirects, r)
		rows = append(rows, &r)
	}

	err := cs.redirectStore.UpsertMany(rows)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save redirects")
	}

	cs.invalidateRedirectCache()

	return redirects, nil
}

// @ID listRedirects
// @Summary List Redirects
// @Description List redirects that the edge proxy applies before routing requests.
// @Tags Proxy Svc
// @Accept json
// @Produce json
// @Param body body proxy.ListRedirectsRequest false "List Redirects Request"
// @Success 200 {object} proxy.ListRedirectsResponse "Redirects listed successfully"
// @Failure 400 {object} proxy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} proxy.ErrorResponse "Unauthorized"
// @Failure 500 {object} proxy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /proxy-svc/redirects [post]
func (cs *ProxyService) ListRedirects(w http.ResponseWriter, r *http.Request) {
	isAuthRsp, statusCode, err := cs.options.PermissionChecker.HasPermission(
		r,
		proxy.PermissionRouteView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	req := &proxy.ListRedirectsRequest{}
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

	redirects, err := cs.listRedirects(req)
	if err != nil {
		logger.Error(
			"Failed to list redirects",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, &proxy.ListRedirectsResponse{
		Redirects: redirects,
	})
}

func (cs *ProxyService) listRedirects(req *proxy.ListRedirectsRequest) ([]proxy.Redirect, error) {
	return cs.cachedRedirects(req.Ids)
}

// @ID deleteRedirects
// @Summary Delete Redirects
// @Description Delete specific redirects by their IDs.
// @Tags Proxy Svc
// @Accept json
// @Produce json
// @Param body body proxy.DeleteRedirectsRequest true "Delete Redirects Request"
// @Success 200 {object} proxy.DeleteRedirectsResponse "Redirects deleted successfully"
// @Failure 400 {object} proxy.ErrorResponse "Invalid JSON or missing IDs"
// @Failure 401 {object} proxy.ErrorResponse "Unauthorized"
// @Failure 500 {object} proxy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /proxy-svc/redirects [delete]
func (cs *ProxyService) DeleteRedirects(w http.ResponseWriter, r *http.Request) {
	isAuthRsp, statusCode, err := cs.options.PermissionChecker.HasPermission(
		r,
		proxy.PermissionRouteEdit,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	req := &proxy.DeleteRedirectsRequest{}
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

	err = cs.deleteRedirects(req)
	if err != nil {
		logger.Error(
			"Failed to delete redirects",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, &proxy.DeleteRedirectsResponse{})
}

func (cs *ProxyService) deleteRedirects(req *proxy.DeleteRedirectsRequest) error {
	if len(req.Ids) == 0 {
		return errors.New("no redirect IDs provided")
	}

	ids := []any{}
	for _, v := range req.Ids {
		ids = append(ids, v)
	}

	err := cs.redirectStore.Query(
		datastore.IsInList([]string{"id"}, ids...),
	).Delete()
	if err != nil {
		return errors.Wrap(err, "failed to delete redirects from store")
	}

	cs.invalidateRedirectCache()

	return nil
}

func normalizeRedirectStatusCode(statusCode int) (int, error) {
	if statusCode == 0 {
		return defaultRedirectStatusCode, nil
	}

	switch statusCode {
	case http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusSeeOther,
		http.StatusTemporaryRedirect,
		http.StatusPermanentRedirect:
		return statusCode, nil
	default:
		return 0, errors.Errorf("unsupported redirect status code %d", statusCode)
	}
}
