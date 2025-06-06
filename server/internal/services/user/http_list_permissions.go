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

// @ID listPermissions
// @Summary List Permissions
// @Description List permissions by roles.
// @Description Caller can only list permissions for roles they have.
// @Tags User Svc
// @Accept  json
// @Produce  json
// @Param   roleId     path    string     true        "Role ID"
// @Success 200 {object} user.ListPermissionsResponse
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/permissions [post]
func (s *UserService) ListPermissions(
	w http.ResponseWriter,
	r *http.Request) {

	req := user.ListPermissionsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	claim, err := s.options.Authorizer.ParseJWTFromRequest(s.publicKeyPem, r)
	if err != nil || claim == nil {
		endpoint.Unauthorized(w)
		return
	}
	rolesIndex := map[string]bool{}
	for _, role := range claim.Roles {
		rolesIndex[role] = true
	}
	for _, role := range req.Roles {
		if !rolesIndex[role] {
			endpoint.WriteString(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
	}

	permissions, err := s.listPermissions(req)
	if err != nil {
		logger.Error(
			"Failed to list permission",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	rsp := user.ListPermissionsResponse{
		Permissions: permissions,
	}
	endpoint.WriteJSON(w, http.StatusOK, rsp)
}

func (s *UserService) listPermissions(
	request user.ListPermissionsRequest,
) ([]string, error) {
	roles := []any{}
	for _, role := range request.Roles {
		roles = append(roles, role)
	}

	permissionsI, err := s.permitsStore.Query(
		datastore.Intersects([]string{"roles"}, roles),
	).
		OrderBy(
			datastore.OrderByField("createdAt", false),
		).
		Find()

	if err != nil {
		return nil, err
	}

	permissions := []string{}
	for _, permissionI := range permissionsI {
		permissions = append(permissions, permissionI.(*user.Permit).Permission)
	}

	return permissions, nil
}
