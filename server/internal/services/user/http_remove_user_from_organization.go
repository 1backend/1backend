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
)

// @ID removeUserFromOrganization
// @Summary Remove a User from an Organization
// @Description Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param organizationId path string true "Organization ID"
// @Param userId path string true "User ID"
// @Param body body user.RemoveUserFromOrganizationRequest false "Add User to Organization Request"
// @Success 200 {object} user.RemoveUserFromOrganizationResponse "User added successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 403 {object} user.ErrorResponse "Forbidden"
// @Failure 404 {object} user.ErrorResponse "Organization/User not found"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/organization/{organizationId}/user/{userId} [delete]
func (s *UserService) RemoveUserFromOrganization(
	w http.ResponseWriter,
	r *http.Request,
) {

	organizationId := mux.Vars(r)["organizationId"]
	userId := mux.Vars(r)["userId"]

	usr, err := s.isAuthorized(
		r,
		user.PermissionOrganizationCreate.Id,
		nil,
		nil,
	)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	req := user.RemoveUserFromOrganizationRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = s.removeUserFromOrganization(usr.Id, userId, organizationId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.RemoveUserFromOrganizationResponse{})
	w.Write(bs)
}

func (s *UserService) removeUserFromOrganization(
	callerId, userId, organizationId string,
) error {
	roleIds, err := s.getRoleIdsByUserId(callerId)
	if err != nil {
		return err
	}

	org, found, err := s.organizationsStore.Query(datastore.Id(organizationId)).
		FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("organization not found")
	}

	if !contains(
		roleIds,
		fmt.Sprintf("user-svc:org:{%v}:admin", org.(*user.Organization).Id),
	) {
		return fmt.Errorf("unauthorized")
	}

	return s.removeRoleFromUser(userId, "user-svc:org:{%v}:user")
}
