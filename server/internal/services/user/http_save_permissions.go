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

	"github.com/openorch/openorch/sdk/go/datastore"
	user "github.com/openorch/openorch/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID savePermissions
// @Summary Save Permissions
// @Description Creates or updates a list of permissions.
// @Description <b>The permission ID must be prefixed by the callers slug.</b>
// @Description Eg. if the owner's slug is `petstore-svc` the permission should look like `petstore-svc:pet:edit`.
// @Descripion The user account who creates the permission will become the owner of that permission, and only the owner will be able to edit the permission.
// @Description
// @Description Requires the `user-svc:permission:create` permission.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.SavePermissionsRequest true "Permission Details"
// @Success 200 {object} user.SavePermissionsResponse
// @Failure 400 {string} string "Bad Request: Invalid JSON or Bad Namespace"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/permissions [put]
func (s *UserService) SavePermissions(
	w http.ResponseWriter,
	r *http.Request,
) {

	usr, err := s.isAuthorized(r, user.PermissionPermissionCreate.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	req := &user.SavePermissionsRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	permissions, err := s.savePermissions(
		usr.Id,
		req,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.SavePermissionsResponse{
		Permissions: permissions,
	})
	w.Write(bs)
}

func (s *UserService) savePermissions(
	userId string,
	req *user.SavePermissionsRequest,
) ([]*user.Permission, error) {
	ret := []*user.Permission{}

	for _, permission := range req.Permissions {
		query := s.permissionsStore.Query(
			datastore.Equals(datastore.Field("id"), permission.Id),
		)

		permI, found, err := query.FindOne()
		if err != nil {
			return nil, err
		}

		if found {
			perm := permI.(*user.Permission)
			if perm.OwnerId != userId {
				return nil, fmt.Errorf("cannot update unowned permission")
			}

			perm.Name = permission.Name
			perm.Description = permission.Description
			err = query.Update(perm)
			if err != nil {
				return nil, errors.Wrap(err, "error updating permission")
			}

			ret = append(ret, perm)
			continue
		}

		perm := &user.Permission{
			Id:          permission.Id,
			Name:        permission.Name,
			Description: permission.Description,
			OwnerId:     userId,
		}

		err = s.permissionsStore.Create(perm)
		if err != nil {
			return nil, errors.Wrap(err, "error creating permission")
		}

		ret = append(ret, perm)
	}

	return ret, nil
}
