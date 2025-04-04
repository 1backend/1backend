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

	usr, isAuthorized, err := s.isAuthorized(
		r,
		user.PermissionOrganizationCreate.Id,
		nil,
		nil,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	req := user.SaveOrganizationRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	org, token, err := s.saveOrganization(usr.Id, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.SaveOrganizationResponse{
		Organization: *org,
		Token:        *token,
	})
	w.Write(bs)
}

func (s *UserService) saveOrganization(
	userId string,
	request *user.SaveOrganizationRequest,
) (*user.Organization, *user.AuthToken, error) {
	if request.Slug == "" {
		return nil, nil, errors.New("slug is required")
	}

	orgI, exists, err := s.organizationsStore.Query(
		datastore.Equals(datastore.Field("slug"), request.Slug),
	).FindOne()
	if err != nil {
		return nil, nil, err
	}

	var final *user.Organization

	if exists {
		final = orgI.(*user.Organization)
		if request.Name != "" {
			final.Name = request.Name
		}
		if request.ThumbnailFileId != "" {
			final.ThumbnailFileId = request.ThumbnailFileId
		}
	} else {
		final = &user.Organization{}

		if request.Id != "" {
			final.Id = request.Id
		} else {
			final.Id = sdk.Id("org")
		}

		if request.Name == "" {
			return nil, nil, errors.New("name is required")
		}
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
		Id:             fmt.Sprintf("%v:%v", final.Id, userId),
		UserId:         userId,
		OrganizationId: final.Id,
		Active:         count == 0, // make the first org active
	}

	err = s.organizationUserLinksStore.Upsert(link)
	if err != nil {
		return nil, nil, err
	}

	err = s.organizationsStore.Upsert(final)
	if err != nil {
		return nil, nil, err
	}

	err = s.addDynamicRoleToUser(
		userId,
		fmt.Sprintf("user-svc:org:{%v}:admin", final.Id),
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
