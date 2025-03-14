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
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openorch/openorch/sdk/go/datastore"
	user "github.com/openorch/openorch/server/internal/services/user/types"
	usertypes "github.com/openorch/openorch/server/internal/services/user/types"
)

// @ID deleteRole
// @Summary Delete a Role
// @Description Delete a role based on the role ID.
// @Tags User Svc
// @Accept  json
// @Produce  json
// @Param   roleId     path    string  true  "Role ID"
// @Success 200 {object} user.DeleteRoleResponse
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/role/{roleId} [delete]
func (s *UserService) DeleteRole(w http.ResponseWriter, r *http.Request) {

	_, err := s.isAuthorized(r, user.PermissionRoleDelete.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	roleId := mux.Vars(r)["roleId"]

	err = s.deleteRole(roleId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.DeleteRoleResponse{})
	w.Write(bs)
}

func (s *UserService) deleteRole(roleId string) error {
	q := s.rolesStore.Query(
		datastore.Id(roleId),
	)
	roleI, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}
	role := roleI.(*usertypes.Role)

	if role.Id == usertypes.RoleAdmin.Id {
		return errors.New("cannot delete default role")
	}

	return q.Delete()
}

func (s *UserService) removeRole(userId string, roleId string) error {
	query := s.usersStore.Query(
		datastore.Equals(datastore.Field("id"), userId),
	)
	userI, found, err := query.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	return s.userRoleLinksStore.Query(
		datastore.Id(fmt.Sprintf("%v:%v", user.Id, roleId)),
	).Delete()

}
