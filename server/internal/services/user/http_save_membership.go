/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/gorilla/mux"
)

// @ID saveMembership
// @Summary Save Membership
// @Description Allows an organization admin to add a user to the organization.
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

	usr, hasPermission, claims, err := s.hasPermission(
		r,
		user.PermissionOrganizationAddUser,
	)
	if err != nil {
		logger.Error(
			"Failed to check permission",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}
	if !hasPermission {
		endpoint.Unauthorized(w)
		return
	}

	req := user.SaveMembershipRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	err = s.saveMembership(claims.App, usr.Id, userId, organizationId)
	if err != nil {
		logger.Error(
			"Failed to save membership",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.SaveMembershipResponse{})
}

func (s *UserService) saveMembership(
	app string,
	callerId,
	userId,
	organizationId string,
) error {
	roles, err := s.getRolesByUserId(app, callerId)
	if err != nil {
		return err
	}

	orgI, found, err := s.organizationsStore.Query(
		datastore.Equals(datastore.Field("app"), app),
		datastore.Id(organizationId),
	).
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
		app,
		userId,
		newRole,
	)
	if err != nil {
		return err
	}

	// When creating a new org, the user switches to that org as the active one
	link := &user.Membership{
		Id:             sdk.Id("memb"),
		App:            app,
		UserId:         userId,
		OrganizationId: org.Id,
		// @todo null out the other active orgs for correctness
		Active: true,
	}

	err = s.membershipsStore.Upsert(link)
	if err != nil {
		return err
	}

	return s.inactivateTokens(app, userId)
}

func contains(ss []string, s string) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}

	return false
}
