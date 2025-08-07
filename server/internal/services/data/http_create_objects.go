/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dynamicservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	dynamictypes "github.com/1backend/1backend/server/internal/services/data/types"
)

func (g *DataService) CreateMany(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := g.options.PermissionChecker.HasPermission(
		r,
		dynamictypes.PermissionObjectCreate,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	req := &dynamictypes.CreateManyRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error("Error decoding JSON", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	defer r.Body.Close()

	err = g.createMany(req)
	if err != nil {
		logger.Error("Error creating objects", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	_, err = w.Write([]byte(`{}`))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
