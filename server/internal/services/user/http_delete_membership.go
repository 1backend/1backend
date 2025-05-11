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

	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/gorilla/mux"
)

// @ID deleteMembership
// @Summary Delete Membership
// @Description Allows an organization admin to remove a user from an organization.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param organizationId path string true "Organization ID"
// @Param userId path string true "User ID"
// @Param body body user.DeleteMembershipRequest false "Remove User From Organization Request"
// @Success 200 {object} user.DeleteMembershipResponse "User added successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 403 {object} user.ErrorResponse "Forbidden"
// @Failure 404 {object} user.ErrorResponse "Organization/User not found"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/organization/{organizationId}/user/{userId} [delete]
func (s *UserService) DeleteMembership(
	w http.ResponseWriter,
	r *http.Request,
) {

	organizationId := mux.Vars(r)["organizationId"]
	userId := mux.Vars(r)["userId"]

	usr, hasPermission, err := s.hasPermission(
		r,
		user.PermissionOrganizationCreate,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !hasPermission {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	req := user.DeleteMembershipRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = s.deleteMembership(usr.Id, userId, organizationId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.DeleteMembershipResponse{})
	w.Write(bs)
}

func (s *UserService) deleteMembership(
	callerId, userId, organizationId string,
) error {
	roleIds, err := s.getRolesByUserId(callerId)
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
