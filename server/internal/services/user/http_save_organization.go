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
	"log/slog"
	"net/http"
	"time"

	"github.com/pkg/errors"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID saveOrganization
// @Summary Save an Organization
// @Description Allows a logged-in user to save an organization. The user initiating the request will be assigned the role of admin for that organization.
// @Description The initiating user will receive a dynamic role in the format `user-svc:org:{organizationId}:admin`, where `{organizationId}` is a unique identifier for the saved organization.
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

	usr, hasPermission, err := s.hasPermission(
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

	org, token, err := s.saveOrganization(usr.Id, &req)
	if err != nil {
		logger.Error(
			"Failed to save organization",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(user.SaveOrganizationResponse{
		Organization: *org,
		Token:        *token,
	})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (s *UserService) saveOrganization(
	userId string,
	request *user.SaveOrganizationRequest,
) (*user.Organization, *user.AuthToken, error) {
	if request.Slug == "" {
		return nil, nil, errors.New("slug is required")
	}
	if request.Name == "" {
		return nil, nil, errors.New("name is required")
	}

	orgI, exists, err := s.organizationsStore.Query(
		datastore.Equals(datastore.Field("slug"), request.Slug),
	).FindOne()
	if err != nil {
		return nil, nil, err
	}

	var final *user.Organization
	now := time.Now()

	if exists {
		final = orgI.(*user.Organization)
		if request.Name != "" {
			final.Name = request.Name
		}
		if request.ThumbnailFileId != "" {
			final.ThumbnailFileId = request.ThumbnailFileId
		}
		final.UpdatedAt = now
	} else {
		final = &user.Organization{
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

		// When creating a new org, the user switches to that org as the active one
		link := &user.Membership{
			Id:             sdk.Id("oul"),
			UserId:         userId,
			OrganizationId: final.Id,
			// @todo null out the other active orgs for correctness
			Active: true,
		}

		err = s.membershipsStore.Upsert(link)
		if err != nil {
			return nil, nil, err
		}

	}

	err = s.organizationsStore.Upsert(final)
	if err != nil {
		return nil, nil, err
	}

	_, err = s.saveEnrolls(
		userId,
		&user.SaveEnrollsRequest{
			Enrolls: []user.EnrollInput{
				{
					UserId: userId,
					Role:   fmt.Sprintf("user-svc:org:{%v}:admin", final.Id),
				},
			},
		},
	)
	if err != nil {
		return nil, nil, err
	}

	err = s.inactivateTokens(userId)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error inactivating tokens")
	}

	userI, found, err := s.usersStore.Query(
		datastore.Id(userId),
	).FindOne()
	if err != nil {
		return nil, nil, errors.Wrap(err, "error finding user by id")
	}
	if !found {
		return nil, nil, errors.New("user not found by id")
	}
	u := userI.(*user.User)

	token, err := s.generateAuthToken(u)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error generating token")
	}

	err = s.authTokensStore.Upsert(token)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating token")
	}

	return final, token, nil
}

func (s *UserService) inactivateTokens(userId string) error {
	return s.authTokensStore.Query(
		datastore.Equals(
			datastore.Fields("userId"),
			userId,
		)).UpdateFields(map[string]any{
		"active": false,
	})
}

func (s *UserService) inactivateToken(tokenId string) error {
	return s.authTokensStore.Query(
		datastore.Equals(
			datastore.Fields("id"),
			tokenId,
		)).UpdateFields(map[string]any{
		"active": false,
	})
}
