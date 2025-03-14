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
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openorch/openorch/sdk/go/datastore"
	user "github.com/openorch/openorch/server/internal/services/user/types"
	usertypes "github.com/openorch/openorch/server/internal/services/user/types"
)

// @ID setRolePermission
// @Summary Set Role Permissions
// @Description Set permissions for a specified role. The caller can add permissions it owns to any role.
// @Description If the caller tries to add a permission it doesn't own to a role, `StatusBadRequest` will be returned.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param roleId path string true "Role ID"
// @Param body body user.SetRolePermissionsRequest true "Set Role Permissions Request"
// @Success 200 {object} user.SetRolePermissionsResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/role/{roleId}/permissions [put]
func (s *UserService) SetRolePermissions(
	w http.ResponseWriter,
	r *http.Request) {

	usr, err := s.isAuthorized(r, user.PermissionRoleEdit.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	req := user.SetRolePermissionsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	roleId := mux.Vars(r)["roleId"]

	err = s.overwriteRolePermissions(usr.Id, roleId, req.PermissionIds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.SetRolePermissionsResponse{})
	w.Write(bs)
}

func (s *UserService) overwriteRolePermissions(
	userId, roleId string,
	permissionIds []string,
) error {
	q := s.rolesStore.Query(
		datastore.Id(roleId),
	)
	roleI, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("cannot find role %v", roleId)
	}
	role := roleI.(*usertypes.Role)
	if role.OwnerId != userId {
		return fmt.Errorf("cannot add permission to unowned role")
	}

	perms, err := s.permissionsStore.Query(
		datastore.Equals(datastore.Field("id"), permissionIds),
	).Find()
	if err != nil {
		return err
	}
	if len(perms) < len(permissionIds) {
		return fmt.Errorf("cannot find some permissions")
	}

	err = s.userRoleLinksStore.Query(
		datastore.Equals(datastore.Field("roleId"), roleId),
	).Delete()
	if err != nil {
		return err
	}

	for _, permissionId := range permissionIds {
		err = s.assignPermissions(userId, []*user.PermissionLink{
			{
				RoleId:       role.Id,
				PermissionId: permissionId,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
