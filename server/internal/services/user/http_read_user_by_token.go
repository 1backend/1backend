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
	"net/http"

	"github.com/openorch/openorch/sdk/go/datastore"
	user "github.com/openorch/openorch/server/internal/services/user/types"
)

// @ID readUserByToken
// @Summary Read User by Token
// @Description Retrieve user information based on an authentication token.
// @Tags User Svc
// @Accept json
// @Produce json
// @Success 200 {object} user.ReadUserByTokenResponse
// @Failure 400 {object} user.ErrorResponse "Token Missing"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/user/by-token [post]
func (s *UserService) ReadUserByToken(w http.ResponseWriter, r *http.Request) {

	token, exists := s.authorizer.TokenFromRequest(r)
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Token Missing`))
		return
	}

	usr, err := s.readUserByToken(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	usr.PasswordHash = ""

	orgs, activeOrgId, err := s.getUserOrganizations(usr.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.ReadUserByTokenResponse{
		User:                 usr,
		Organizations:        orgs,
		ActiveOrganizationId: activeOrgId,
	})
	w.Write(bs)
}

func (s *UserService) readUserByToken(token string) (*user.User, error) {
	authTokenI, found, err := s.authTokensStore.Query(
		datastore.Equals(datastore.Field("token"), token),
	).FindOne()
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("token not found")
	}
	authToken := authTokenI.(*user.AuthToken)

	userI, found, err := s.usersStore.Query(
		datastore.Equals(datastore.Field("id"), authToken.UserId),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("user not found")
	}

	u := userI.(*user.User)

	ret := &user.User{
		Id:        u.Id,
		Name:      u.Name,
		Slug:      u.Slug,
		Contacts:  u.Contacts,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return ret, nil
}

func (s *UserService) getUserOrganizations(
	userId string,
) ([]*user.Organization, string, error) {
	links, err := s.organizationUserLinksStore.Query(
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
		link := linkI.(*user.OrganizationUserLink)
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
