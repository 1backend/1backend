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

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

var ErrOrganizationMembershipNotFound = errors.New("organization membership not found")

// @ID activateOrganization
// @Summary Activate Organization
// @Description Sets the caller user's active organization and returns a fresh token reflecting the new active organization.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ActivateOrganizationRequest true "Activate Organization Request"
// @Success 200 {object} user.ActivateOrganizationResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/organization/activate [post]
func (s *UserService) ActivateOrganization(
	w http.ResponseWriter,
	r *http.Request,
) {
	claims, err := s.options.Authorizer.ParseJWTFromRequest(s.publicKeyPem, r)
	if err != nil {
		logger.Error(
			"Failed to parse JWT",
			slog.Any("error", err),
		)
		endpoint.Unauthorized(w)
		return
	}

	usr, err := s.readSelf(claims.UserId)
	if err != nil {
		logger.Error(
			"Failed to read self",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	req := user.ActivateOrganizationRequest{}
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

	if req.OrganizationId == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "organizationId is required")
		return
	}

	token, err := s.activateOrganization(
		claims.AppId,
		usr,
		req.OrganizationId,
		claims.Device,
	)
	if err != nil {
		if errors.Is(err, ErrOrganizationMembershipNotFound) {
			endpoint.Unauthorized(w)
			return
		}

		logger.Error(
			"Failed to activate organization",
			slog.Any("error", err),
			slog.String("organizationId", req.OrganizationId),
			slog.String("userId", usr.Id),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.ActivateOrganizationResponse{
		Token: *token,
	})
}

func (s *UserService) activateOrganization(
	appId string,
	usr *user.User,
	organizationId string,
	device string,
) (*user.Token, error) {
	links, err := s.membershipStore.Query(
		datastore.Equals(datastore.Field("appId"), appId),
		datastore.Equals(datastore.Field("userId"), usr.Id),
	).Find()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query memberships")
	}

	found := false
	for _, linkI := range links {
		link := linkI.(*user.Membership)
		if link.OrganizationId == organizationId {
			found = true
			break
		}
	}

	if !found {
		return nil, ErrOrganizationMembershipNotFound
	}

	now := time.Now()
	deviceLinks := []*user.Membership{}
	deviceOrganizationFound := false
	for _, linkI := range links {
		link := linkI.(*user.Membership)
		if link.Device != device {
			continue
		}
		deviceLinks = append(deviceLinks, link)
		if link.OrganizationId == organizationId {
			deviceOrganizationFound = true
		}
	}

	for _, link := range deviceLinks {
		isActive := link.OrganizationId == organizationId
		if link.Active == isActive {
			continue
		}

		err = s.membershipStore.Query(
			datastore.Id(link.Id),
		).UpdateFields(map[string]any{
			"active":    isActive,
			"updatedAt": now,
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to update membership")
		}
	}

	if !deviceOrganizationFound {
		id := sdk.Id("memb")
		internalId, err := sdk.InternalId(appId, id)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create membership internal id")
		}

		err = s.membershipStore.Upsert(&user.Membership{
			InternalId:     internalId,
			Id:             id,
			AppId:          appId,
			UserId:         usr.Id,
			OrganizationId: organizationId,
			Device:         device,
			Active:         true,
			CreatedAt:      now,
			UpdatedAt:      now,
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to create membership for device")
		}
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
