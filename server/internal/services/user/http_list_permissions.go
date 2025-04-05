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

	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID listPermissions
// @Summary List Permissions
// @Description Retrieve permissions by roles.
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

	_, isAuthorized, err := s.hasPermission(r, user.PermissionRoleView, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	req := user.ListPermissionsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	permissions, err := s.getPermissions()
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

func (s *UserService) getPermissions() ([]string, error) {
	permissionsI, err := s.permissionRoleLinksStore.Query().
		OrderBy(datastore.OrderByField("createdAt", false)).
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
