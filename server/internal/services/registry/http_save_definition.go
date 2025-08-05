/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package registryservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
)

// Register a new definition
// @ID saveDefinition
// @Summary Register a Definition
// @Description Registers a new definition, associating an definition address with a slug acquired from the bearer token.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param body body registry.SaveDefinitionRequest true "Register Service Definition Request"
// @Success 201 {object} registry.SaveDefinitionResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid JSON"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/definition [put]
func (rs *RegistryService) SaveDefinition(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := rs.options.PermissionChecker.HasPermission(
		r,
		registry.PermissionNodeView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	req := &registry.SaveDefinitionRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logger.Error(
			"Error decoding request body",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	err = rs.saveServiceDefinition(req)
	if err != nil {
		logger.Error(
			"Error saving service definition",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(`{}`))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (rs *RegistryService) saveServiceDefinition(
	req *registry.SaveDefinitionRequest,
) error {
	return rs.definitionStore.Upsert(req.Definition)
}
