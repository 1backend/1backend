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
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/gorilla/mux"
)

// @ID saveMembership
// @Summary Save Membership
// @Description Creates or updates an organization membership invite.
// @Description Memberships automatically include a canonical per-user org role in the form `user-svc:org:{organizationId}:{userId}`.
// @Description Callers may add additional roles under the same organization prefix to model org-local groups or custom permission bundles.
// @Description Additional roles are still subject to role ownership validation via `OwnsRole`.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param organizationId path string true "Organization ID"
// @Param userId path string true "User ID"
// @Param body body user.SaveMembershipRequest false "Add User to Organization Request"
// @Success 200 {object} user.SaveMembershipResponse "Membership saved successfully"
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
	if err != nil && err != io.EOF {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	membership, err := s.saveMembership(
		claims.AppId,
		usr.Id,
		usr.Slug,
		userId,
		organizationId,
		req.Roles,
	)
	if err != nil {
		if errors.Is(err, ErrOrganizationAdminRequired) || errors.Is(err, ErrMembershipRoleNotOwned) {
			endpoint.WriteString(w, http.StatusForbidden, "Forbidden")
			return
		}
		if errors.Is(err, ErrMembershipRoleOutsideOrganizationScope) || errors.Is(err, ErrMembershipRolesEmpty) {
			endpoint.WriteString(w, http.StatusBadRequest, err.Error())
			return
		}

		logger.Error(
			"Failed to save membership",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.SaveMembershipResponse{
		Membership: *membership,
	})
}

func (s *UserService) saveMembership(
	appId string,
	callerId,
	callerSlug,
	userId,
	organizationId string,
	Roles []string,
) (*user.Membership, error) {
	roles, err := s.getRolesByUserId(appId, callerId)
	if err != nil {
		return nil, err
	}

	orgI, found, err := s.organizationStore.Query(
		datastore.Equals(datastore.Field("appId"), appId),
		datastore.Id(organizationId),
	).
		FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("organization not found")
	}

	org := orgI.(*user.Organization)

	if !hasOrganizationAdminAccess(roles, org.Id) {
		return nil, ErrOrganizationAdminRequired
	}

	_, targetUserFound, err := s.userStore.Query(datastore.Id(userId)).FindOne()
	if err != nil {
		return nil, err
	}
	if !targetUserFound {
		return nil, fmt.Errorf("user not found")
	}

	normalizedRoles, err := normalizeMembershipRoles(org.Id, userId, Roles)
	if err != nil {
		return nil, err
	}
	if err := validateMembershipRoleOwnership(callerSlug, roles, normalizedRoles); err != nil {
		return nil, err
	}

	existingMembership, found, err := s.findMembership(appId, userId, org.Id)
	if err != nil {
		return nil, err
	}

	if found {
		if existingMembership.Status == user.MembershipStatusAccepted {
			updatedMembership, changed, err := s.updateMembership(
				existingMembership,
				user.MembershipStatusAccepted,
				normalizedRoles,
				existingMembership.InvitedBy,
				existingMembership.AcceptedAt,
			)
			if err != nil {
				return nil, err
			}
			if changed {
				if err := s.inactivateTokens(appId, userId); err != nil {
					return nil, err
				}
			}
			return updatedMembership, nil
		}

		updatedMembership, _, err := s.updateMembership(
			existingMembership,
			user.MembershipStatusPending,
			normalizedRoles,
			callerId,
			nil,
		)
		if err != nil {
			return nil, err
		}
		return updatedMembership, nil
	}

	return s.createMembership(
		appId,
		userId,
		org.Id,
		user.MembershipStatusPending,
		normalizedRoles,
		callerId,
		nil,
	)
}

func contains(ss []string, s string) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}

	return false
}
