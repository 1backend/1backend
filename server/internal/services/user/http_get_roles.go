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

// @ID listRoles
// @Summary List Roles
// @Description Retrieve all roles from the user service.
// @Tags User Svc
// @Accept json
// @Produce json
// @Success 200 {object} user.ListRolesResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/roles [get]
func (s *UserService) ListRoles(
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

	roles, err := s.listRoles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.ListRolesResponse{
		Roles: roles,
	})
	w.Write(bs)
}

func (s *UserService) listRoles() ([]string, error) {
	linkIs, err := s.permissionRoleLinksStore.Query().
		OrderBy(datastore.OrderByField("name", false)).
		Find()

	if err != nil {
		return nil, err
	}

	roles := []string{}
	for _, linkI := range linkIs {
		roles = append(roles, linkI.(*user.PermissionRoleLink).Role)
	}

	return roles, err
}
