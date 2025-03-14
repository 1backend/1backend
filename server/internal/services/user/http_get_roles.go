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

	"github.com/openorch/openorch/sdk/go/datastore"
	user "github.com/openorch/openorch/server/internal/services/user/types"
	usertypes "github.com/openorch/openorch/server/internal/services/user/types"
)

// @ID getRoles
// @Summary Get all Roles
// @Description Retrieve all roles from the user service.
// @Tags User Svc
// @Accept json
// @Produce json
// @Success 200 {object} user.GetRolesResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/roles [get]
func (s *UserService) GetRoles(
	w http.ResponseWriter,
	r *http.Request) {

	_, err := s.isAuthorized(r, user.PermissionRoleView.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	roles, err := s.getRoles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.GetRolesResponse{
		Roles: roles,
	})
	w.Write(bs)
}

func (s *UserService) getRoles() ([]*usertypes.Role, error) {
	rolesI, err := s.rolesStore.Query().
		OrderBy(datastore.OrderByField("name", false)).
		Find()

	if err != nil {
		return nil, err
	}

	roles := []*usertypes.Role{}
	for _, roleI := range rolesI {
		roles = append(roles, roleI.(*usertypes.Role))
	}

	return roles, err
}
