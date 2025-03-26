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
	"time"

	"github.com/pkg/errors"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID createOrganization
// @Summary Create an Organization
// @Description Allows a logged-in user to create a new organization. The user initiating the request will be assigned the role of admin for that organization.
// @Description The initiating user will receive a dynamic role in the format `user-svc:org:{organizationId}:admin`, where `{organizationId}` is a unique identifier for the created organization.
// @Description Dynamic roles are generated based on specific user-resource associations (in this case the resource being the organization), offering more flexible permission management compared to static roles.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.CreateOrganizationRequest true "Create User Request"
// @Success 200 {object} user.CreateOrganizationResponse "User created successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/organization [post]
func (s *UserService) CreateOrganization(
	w http.ResponseWriter,
	r *http.Request) {

	usr, _, err := s.isAuthorized(
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

	req := user.CreateOrganizationRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	org, token, err := s.createOrganization(usr.Id, req.Id, req.Name, req.Slug)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.CreateOrganizationResponse{
		Organization: *org,
		Token:        *token,
	})
	w.Write(bs)
}

func (s *UserService) createOrganization(
	userId, orgId, name, slug string,
) (*user.Organization, *user.AuthToken, error) {
	_, exists, err := s.organizationsStore.Query(
		datastore.Equals(datastore.Field("slug"), slug),
	).FindOne()
	if err != nil {
		return nil, nil, err
	}

	if exists {
		return nil, nil, errors.New("organization already exists")
	}

	if orgId == "" {
		orgId = sdk.Id("org")
	}

	org := &user.Organization{
		Id:   orgId,
		Name: name,
		Slug: slug,
	}

	count, err := s.organizationUserLinksStore.Query(
		datastore.Equals(
			datastore.Field("userId"),
			userId,
		),
	).Count()
	if err != nil {
		return nil, nil, err
	}

	link := &user.OrganizationUserLink{
		Id:             fmt.Sprintf("%v:%v", org.Id, userId),
		UserId:         userId,
		OrganizationId: org.Id,
		Active:         count == 0, // make the first org active
	}

	err = s.organizationUserLinksStore.Create(link)
	if err != nil {
		return nil, nil, err
	}

	err = s.organizationsStore.Create(org)
	if err != nil {
		return nil, nil, err
	}

	err = s.addDynamicRoleToUser(
		userId,
		fmt.Sprintf("user-svc:org:{%v}:admin", org.Id),
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

	err = s.authTokensStore.Create(token)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating token")
	}

	return org, token, nil
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

func (s *UserService) addStaticRoleToUser(userId, roleId string) error {
	roleQ := s.rolesStore.Query(
		datastore.Id(roleId),
	)
	roleI, found, err := roleQ.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("cannot find role %v", roleId)
	}
	role := roleI.(*user.Role)

	userQ := s.usersStore.Query(
		datastore.Id(userId),
	)
	userI, found, err := userQ.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("cannot find user %v", userId)
	}
	u := userI.(*user.User)

	return s.userRoleLinksStore.Upsert(&user.UserRoleLink{
		Id:        fmt.Sprintf("%v:%v", u.Id, role.Id),
		CreatedAt: time.Now(),

		RoleId: roleId,
		UserId: userId,
	})
}

func (s *UserService) addDynamicRoleToUser(userId, roleId string) error {
	userQ := s.usersStore.Query(
		datastore.Id(userId),
	)
	userI, found, err := userQ.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("cannot find user %v", userId)
	}
	u := userI.(*user.User)

	return s.userRoleLinksStore.Upsert(&user.UserRoleLink{
		Id:        fmt.Sprintf("%v:%v", u.Id, roleId),
		CreatedAt: time.Now(),

		RoleId: roleId,
		UserId: userId,
	})
}
