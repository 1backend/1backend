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
	"strings"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"

	"github.com/pkg/errors"
)

// @ID assignPermissions
// @Summary Assign Permissions
// @Description Assign permissions to roles.
// @Description
// @Description Requires the `user-svc:permission:assign` permission.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.AssignPermissionsRequest true "Assign Permissions Request"
// @Success 200 {object} user.AssignPermissionsResponse "Assign Permissions successfully"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/roles/permissions [put]
func (s *UserService) AssignPermissions(
	w http.ResponseWriter,
	r *http.Request,
) {

	usr, err := s.getUserFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	req := user.AssignPermissionsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	for _, link := range req.PermissionLinks {
		if !strings.HasPrefix(link.PermissionId, usr.Slug) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
	}

	err = s.assignPermissions(req.PermissionLinks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.CreateUserResponse{})
	w.Write(bs)
}

func (s *UserService) assignPermissions(
	permissionLinks []*user.PermissionLink,
) error {
	for _, permissionLink := range permissionLinks {
		roleQ := s.rolesStore.Query(
			datastore.Id(permissionLink.RoleId),
		)
		roleI, found, err := roleQ.FindOne()
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("cannot find role %v", permissionLink.RoleId)
		}
		role := roleI.(*usertypes.Role)

		permQ := s.permissionsStore.Query(
			datastore.Id(permissionLink.PermissionId),
		)
		permissionI, found, err := permQ.FindOne()
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("cannot find permission %v", permissionLink.PermissionId)
		}
		permission := permissionI.(*usertypes.Permission)

		err = s.permissionRoleLinksStore.Upsert(&usertypes.PermissionRoleLink{
			Id:           fmt.Sprintf("%v:%v", permission.Id, role.Id),
			CreatedAt:    time.Now(),
			RoleId:       permissionLink.RoleId,
			PermissionId: permissionLink.PermissionId,
		})
		if err != nil {
			return errors.Wrap(err, "error saving permission role link")
		}
	}

	return nil
}
