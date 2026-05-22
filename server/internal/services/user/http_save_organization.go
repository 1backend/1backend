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
	"log/slog"
	"net/http"
	"time"

	"github.com/pkg/errors"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

var ErrNotAnAdmin = errors.New("user is not an admin of the organization")

// @ID saveOrganization
// @Summary Save an Organization
// @Description Allows a logged-in user to save an organization. The user initiating the request will be assigned the role of admin for that organization.
// @Description The initiating user will receive a dynamic role in the format `user-svc:org:{organizationId}:admin`, where `{organizationId}` is a unique identifier for the saved organization.
// @Description The creator membership also receives the canonical per-user org role `user-svc:org:{organizationId}:{userId}` so org-specific permits can target the creator directly.
// @Description Dynamic roles are generated based on specific user-resource associations (in this case the resource being the organization), offering more flexible permission management compared to static roles.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.SaveOrganizationRequest true "Save User Request"
// @Success 200 {object} user.SaveOrganizationResponse "User saved successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/organization [put]
func (s *UserService) SaveOrganization(
	w http.ResponseWriter,
	r *http.Request) {

	usr, hasPermission, claims, err := s.hasPermission(
		r,
		user.PermissionOrganizationCreate,
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

	req := user.SaveOrganizationRequest{}
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

	org, token, err := s.saveOrganization(
		claims.AppId,
		usr.Id,
		&req,
		claims,
	)
	if err != nil {
		if errors.Is(err, ErrNotAnAdmin) {
			endpoint.WriteErr(w, http.StatusUnauthorized, err)
			return
		}
		if errors.Is(err, ErrOrganizationMembershipNotFound) {
			endpoint.WriteErr(w, http.StatusUnauthorized, err)
			return
		}

		logger.Error(
			"Failed to save organization",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	rsp := user.SaveOrganizationResponse{
		Organization: *org,
		Token:        token,
	}
	endpoint.WriteJSON(w, http.StatusOK, rsp)
}

func (s *UserService) saveOrganization(
	appId string,
	userId string,
	request *user.SaveOrganizationRequest,
	claims *auth.Claims,
) (*user.Organization, *user.Token, error) {
	if request.Slug == "" {
		return nil, nil, errors.New("slug is required")
	}
	if request.Name == "" {
		return nil, nil, errors.New("name is required")
	}

	orgI, exists, err := s.organizationStore.Query(
		datastore.Equals(datastore.Field("appId"), appId),
		datastore.Equals(datastore.Field("slug"), request.Slug),
	).FindOne()
	if err != nil {
		return nil, nil, err
	}

	var final *user.Organization
	now := time.Now()

	shouldActivate := request.Activate

	if exists {
		final = orgI.(*user.Organization)

		roles, err := s.getRolesByUserId(claims.AppId, claims.UserId)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to list effective roles")
		}

		if !hasOrganizationAdminAccess(roles, final.Id) {
			return nil, nil, ErrNotAnAdmin
		}

		if request.Name != "" {
			final.Name = request.Name
		}
		if request.ThumbnailFileId != "" {
			final.ThumbnailFileId = request.ThumbnailFileId
		}
		final.UpdatedAt = now
		err = s.organizationStore.Upsert(final)
		if err != nil {
			return nil, nil, err
		}

		u, err := s.readSelf(userId)
		if err != nil {
			return nil, nil, errors.Wrap(err, "error finding user by id")
		}

		var token *user.Token
		if shouldActivate {
			token, err = s.activateOrganization(appId, u, final.Id, claims.Device)
			if err != nil {
				return nil, nil, errors.Wrap(err, "error activating organization")
			}
		}

		return final, token, nil
	}

	final = &user.Organization{
		AppId:     appId,
		Name:      request.Name,
		Slug:      request.Slug,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if request.Id != "" {
		final.Id = request.Id
	} else {
		final.Id = sdk.Id("org")
	}

	final.InternalId, err = sdk.InternalId(appId, final.Id)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create organization internal id")
	}

	err = s.organizationStore.Upsert(final)
	if err != nil {
		return nil, nil, err
	}

	userI, found, err := s.userStore.Query(
		datastore.Id(userId),
	).FindOne()
	if err != nil {
		return nil, nil, errors.Wrap(err, "error finding user by id")
	}
	if !found {
		return nil, nil, errors.New("user not found by id")
	}
	u := userI.(*user.User)

	roles, err := normalizeMembershipRoles(final.Id, userId, []string{orgAdminRole(final.Id)})
	if err != nil {
		return nil, nil, errors.Wrap(err, "error normalizing creator membership roles")
	}

	acceptedAt := now
	membership, err := s.createMembership(
		appId,
		userId,
		final.Id,
		user.MembershipStatusAccepted,
		roles,
		"",
		&acceptedAt,
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating creator membership")
	}

	if !shouldActivate {
		return final, nil, nil
	}

	token, err := s.issueTokenForActiveOrganization(
		appId,
		u,
		claims.Device,
		final.Id,
		membership.Roles,
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error activating organization")
	}

	return final, token, nil
}

func (s *UserService) deactivateOrganization(
	appId string,
	usr *user.User,
	device string,
) (*user.Token, error) {
	err := s.setActivation(appId, usr.Id, device, "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to clear activation")
	}

	err = s.inactivateTokens(appId, usr.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to inactivate tokens")
	}

	token, err := s.issueToken(appId, usr, device)
	if err != nil {
		return nil, errors.Wrap(err, "failed to issue token")
	}

	return token, nil
}

func (s *UserService) inactivateTokens(appId string, userId string) error {
	filters := []datastore.Filter{
		datastore.Equals(
			datastore.Fields("userId"),
			userId,
		),
	}

	if appId != "*" {
		filters = append(filters, datastore.Equals(
			datastore.Fields("appId"),
			appId,
		))
	}

	return s.tokenStore.Query(
		filters...,
	).UpdateFields(map[string]any{
		"active": false,
	})
}

func (s *UserService) inactivateToken(appId, tokenId string) error {
	return s.tokenStore.Query(
		datastore.Equals(
			datastore.Fields("appId"),
			appId,
		),
		datastore.Equals(
			datastore.Fields("id"),
			tokenId,
		)).UpdateFields(map[string]any{
		"active": false,
	})
}
