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

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/gorilla/mux"
)

// @ID saveMembership
// @Summary Save Membership
// @Description Allows and organization admint to add a user to an organization.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param organizationId path string true "Organization ID"
// @Param userId path string true "User ID"
// @Param body body user.SaveMembershipRequest false "Add User to Organization Request"
// @Success 200 {object} user.SaveMembershipResponse "User added successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 403 {object} user.ErrorResponse "Forbidden"
// @Failure 404 {object} user.ErrorResponse "Organization/User not found"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/organization/{organizationId}/user/{userId} [put]
func (s *UserService) SaveMembership(
	w http.ResponseWriter,
	r *http.Request,
) {

	organizationId := mux.Vars(r)["organizationId"]
	userId := mux.Vars(r)["userId"]

	usr, hasPermission, err := s.hasPermission(
		r,
		user.PermissionOrganizationAddUser,
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

	req := user.SaveMembershipRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = s.saveMembership(usr.Id, userId, organizationId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.SaveMembershipResponse{})
	w.Write(bs)
}

func (s *UserService) saveMembership(
	callerId, userId, organizationId string,
) error {
	roles, err := s.getRolesByUserId(callerId)
	if err != nil {
		return err
	}

	orgI, found, err := s.organizationsStore.Query(datastore.Id(organizationId)).
		FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("organization not found")
	}

	org := orgI.(*user.Organization)

	if !contains(
		roles,
		fmt.Sprintf("user-svc:org:{%v}:admin", org.Id),
	) {
		return fmt.Errorf("not an admin of the organization")
	}

	newRole := fmt.Sprintf("user-svc:org:{%v}:user", org.Id)

	for _, role := range roles {
		if newRole == role {
			return nil
		}
	}

	err = s.assignRole(
		userId,
		newRole,
	)
	if err != nil {
		return err
	}

	// When creating a new org, the user switches to that org as the active one
	link := &user.Membership{
		Id:             sdk.Id("memb"),
		UserId:         userId,
		OrganizationId: org.Id,
		// @todo null out the other active orgs for correctness
		Active: true,
	}

	err = s.membershipsStore.Upsert(link)
	if err != nil {
		return err
	}

	return s.inactivateTokens(userId)
}

func contains(ss []string, s string) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}

	return false
}
