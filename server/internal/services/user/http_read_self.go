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
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID readSelf
// @Summary Read Self
// @Description Retrieves user information based on the authentication token in the request header.
// @Description Typically called by single-page applications during the initial page load.
// @Description While some details (such as roles, slug, user ID, and active organization ID) can be extracted from the JWT,
// @Description this endpoint returns additional data, including the full user object and associated organizations.
// @Description
// @Description ReadSelf intentionally still works after token revocation until the token expires.
// @Description This is to ensure that the user is not notified of token revocation (though some information is
// @Description leaked by the count token functionality @todo).
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ReadSelfRequest false "Read Self Request"
// @Success 200 {object} user.ReadSelfResponse
// @Failure 400 {object} user.ErrorResponse "Token Missing"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/self [post]
func (s *UserService) ReadSelf(w http.ResponseWriter, r *http.Request) {
	request := user.ReadSelfRequest{}
	if r.ContentLength != 0 {
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			logger.Error(
				"Failed to decode request",
				slog.Any("error", err),
			)
			endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
			return
		}
		defer r.Body.Close()
	}

	claim, err := s.parseJWTFromRequest(r)
	if err != nil {
		logger.Error(
			"Failed to parse JWT",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	usr, err := s.readSelf(claim.UserId)
	if err != nil {
		logger.Error(
			"Failed to read self",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	orgs, activeOrgId, err := s.getUserOrganizations(usr.Id)
	if err != nil {
		logger.Error(
			"Failed to get organizations",
			slog.String("userId", usr.Id),
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	contacts, err := s.getContactsByUserId(usr.Id)
	if err != nil {
		logger.Error(
			"Failed to get contacts",
			slog.String("userId", usr.Id),
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	rsp := user.ReadSelfResponse{
		User:                 usr,
		Roles:                claim.Roles,
		Organizations:        orgs,
		ActiveOrganizationId: activeOrgId,
		Contacts:             contacts,
	}

	if request.CountTokens {
		tokenCount, err := s.countTokens(usr.Id)
		if err != nil {
			logger.Error(
				"Failed to count tokens",
				slog.String("userId", usr.Id),
				slog.Any("error", err),
			)
			endpoint.InternalServerError(w)
			return
		}
		rsp.TokenCount = tokenCount
	}

	bs, _ := json.Marshal(rsp)
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (s *UserService) readSelf(userId string) (*user.User, error) {
	userI, found, err := s.usersStore.Query(
		datastore.Equals(datastore.Field("id"), userId),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("user not found")
	}

	u := userI.(*user.User)

	return u, nil
}

func (s *UserService) getUserOrganizations(
	userId string,
) ([]*user.Organization, string, error) {
	links, err := s.membershipsStore.Query(
		datastore.Equals(
			datastore.Field("userId"),
			userId,
		),
	).Find()
	if err != nil {
		return nil, "", err
	}

	organizationIds := []any{}
	activeOrganizationId := ""
	for _, linkI := range links {
		link := linkI.(*user.Membership)
		if link.Active {
			activeOrganizationId = link.OrganizationId
		}
		organizationIds = append(organizationIds, link.OrganizationId)
	}

	orgIs, err := s.organizationsStore.Query(
		datastore.IsInList(
			datastore.Field("id"),
			organizationIds...,
		),
	).Find()
	if err != nil {
		return nil, "", err
	}

	orgs := []*user.Organization{}
	for _, orgI := range orgIs {
		org := orgI.(*user.Organization)
		orgs = append(orgs, org)
	}

	return orgs, activeOrganizationId, nil
}

func (s *UserService) countTokens(userId string) (int64, error) {
	tokenCount, err := s.authTokensStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
	).Count()
	if err != nil {
		return 0, err
	}

	return tokenCount, nil
}
