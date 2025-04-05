/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package userservice

import (
	"encoding/json"
	"net/http"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	authr := auth.AuthorizerImpl{}
	claim, err := authr.ParseJWTFromRequest(s.publicKeyPem, r)
	if err != nil || claim == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}
	rolesIndex := map[string]bool{}
	for _, role := range claim.RoleIds {
		rolesIndex[role] = true
	}
	for _, role := range req.Roles {
		if !rolesIndex[role] {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
	}

	permissions, err := s.listPermissions(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.ListPermissionsResponse{
		Permissions: permissions,
	})
	w.Write(bs)
}

func (s *UserService) listPermissions(
	request user.ListPermissionsRequest,
) ([]string, error) {
	roles := []any{}
	for _, role := range request.Roles {
		roles = append(roles, role)
	}

	permissionsI, err := s.permissionRoleLinksStore.Query(
		datastore.IsInList([]string{"role"}, roles...),
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
		permissions = append(permissions, permissionI.(*user.PermissionRoleLink).Permission)
	}

	return permissions, nil
}
