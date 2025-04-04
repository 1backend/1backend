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

// @ID getPermissionsByRole
// @Summary Get Permissions by Role
// @Description Retrieve permissions associated with a specific role ID.
// @Tags User Svc
// @Accept  json
// @Produce  json
// @Param   roleId     path    string     true        "Role ID"
// @Success 200 {object} user.GetPermissionsResponse
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/role/{roleId}/permissions [get]
func (s *UserService) GetPermissions(
	w http.ResponseWriter,
	r *http.Request) {

	_, isAuthorized, err := s.isAuthorized(r, user.PermissionRoleView, nil, nil)
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

	permissions, err := s.getPermissions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.GetPermissionsResponse{
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
