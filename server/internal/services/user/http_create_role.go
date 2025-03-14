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
	"strings"

	"github.com/openorch/openorch/sdk/go/datastore"
	user "github.com/openorch/openorch/server/internal/services/user/types"
)

// @ID createRole
// @Summary Create a New Role
// @Description Create a new role.
// @Description <b>The role ID must be prefixed by the caller's slug.</b>
// @Description Eg. if the caller's slug is `petstore-svc` the role should look like `petstore-svc:admin`.
// @Description The user account who creates the role will become the owner of that role, and only the owner will be able to edit the role.
// @Description
// @Description Requires the `user-svc:role:create` permission.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.CreateRoleRequest true "Create Role Request"
// @Success 200 {object} user.CreateRoleResponse "Role created successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/role [post]
func (s *UserService) CreateRole(w http.ResponseWriter, r *http.Request) {

	rsp, err := s.isAuthorized(r, user.PermissionRoleCreate.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	req := user.CreateRoleRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	role, err := s.createRole(
		rsp.Id,
		rsp.Slug,
		&req,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.CreateRoleResponse{
		Role: role,
	})
	w.Write(bs)
}

func (s *UserService) createRole(
	callerId string,
	callerSlug string,
	req *user.CreateRoleRequest,
) (*user.Role, error) {
	permissions, err := s.permissionsStore.Query(
		datastore.Equals(datastore.Field("id"), req.PermissionIds),
	).Find()
	if err != nil {
		return nil, err
	}
	if len(permissions) < len(req.PermissionIds) {
		return nil, errors.New("nonexistent permissions")
	}

	if req.Id == "" {
		return nil, fmt.Errorf("create role is missing id")
	}
	if !strings.HasPrefix(req.Id, callerSlug) {
		return nil, fmt.Errorf("created role id must be prefixed by caller slug")
	}

	role := &user.Role{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		OwnerId:     callerId,
	}

	err = s.rolesStore.Upsert(role)
	if err != nil {
		return nil, err
	}

	links := []*user.PermissionLink{}
	for _, permissionId := range req.PermissionIds {
		links = append(links, &user.PermissionLink{
			RoleId:       role.Id,
			PermissionId: permissionId,
		})
	}

	err = s.assignPermissions(callerId, links)
	if err != nil {
		return nil, err
	}

	return role, nil
}
