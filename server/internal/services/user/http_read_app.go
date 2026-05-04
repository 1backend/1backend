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

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID readApp
// @Summary Read or Create App
// @Description Get an app by host, or create it if it does not exist.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ReadAppRequest true "Read App Request"
// @Success 200 {object} user.ReadAppResponse "App retrieved or created successfully"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/app [post]
func (s *UserService) ReadApp(w http.ResponseWriter, r *http.Request) {
	req := &user.ReadAppRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logger.Error("Failed to decode request", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	if req.Host == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "Host is required")
		return
	}

	// lookup
	app, err := s.findApp(req)
	if err != nil {
		logger.Error("Failed to lookup app", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	if app == nil {
		app, err = s.createApp(req)
		if err != nil {
			logger.Error("Failed to create app", slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}
	}

	endpoint.WriteJSON(w, http.StatusOK, &user.ReadAppResponse{
		App: *app,
	})
}

func (s *UserService) findApp(req *user.ReadAppRequest) (*user.App, error) {
	if req.Host == "" {
		return nil, nil
	}

	app, found, err := s.appByHost(req.Host)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return app, nil
}

func (s *UserService) createApp(req *user.ReadAppRequest) (*user.App, error) {
	app := &user.App{
		Id:   sdk.Id("app"),
		Host: req.Host,
	}
	if err := s.appStore.Upsert(app); err != nil {
		return nil, err
	}
	s.cacheApp(app)
	return app, nil
}
