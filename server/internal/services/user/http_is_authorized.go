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
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/logger"

	user "github.com/openorch/openorch/server/internal/services/user/types"
)

// @ID isAuthorized
// @Summary Is Authorized
// @Description Check if a user is authorized for a specific permission.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param permissionId path string true "Permission ID"
// @Param body body user.IsAuthorizedRequest false "Is Authorized Request"
// @Success 200 {object} user.IsAuthorizedResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON or missing permission id"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Security BearerAuth
// @Router /user-svc/permission/{permissionId}/is-authorized [post]
func (s *UserService) IsAuthorized(
	w http.ResponseWriter,
	r *http.Request,
) {

	req := &user.IsAuthorizedRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	permissionId := vars["permissionId"]

	if permissionId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`missing permission id`))
		return
	}

	if req == nil {
		req = &user.IsAuthorizedRequest{}
	}

	usr, err := s.isAuthorized(r, permissionId, req.GrantedSlugs, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	bs, _ := json.Marshal(&user.IsAuthorizedResponse{
		Authorized: true,
		User:       usr,
	})

	w.Write(bs)
}

func (s *UserService) isAuthorized(r *http.Request, permissionId string,
	grantedSlugs, contactsGranted []string) (*user.User, error) {
	usr, err := s.getUserFromRequest(r)
	if err != nil {
		return nil, err
	}

	slugGrant := false
	for _, v := range grantedSlugs {
		if usr.Slug == v {
			slugGrant = true
		}
	}
	if slugGrant {
		return usr, nil
	}
	roleLinks, err := s.userRoleLinksStore.Query(
		datastore.Equals(datastore.Field("userId"), usr.Id),
	).Find()
	if err != nil {
		return nil, err
	}
	roleIds := []string{}
	for _, role := range roleLinks {
		roleIds = append(roleIds, role.(*user.UserRoleLink).RoleId)
	}

	roleIdAnys := []any{}
	for _, roleId := range roleIds {
		roleIdAnys = append(roleIdAnys, roleId)
	}
	permissionLinks, err := s.permissionRoleLinksStore.Query(
		datastore.IsInList(datastore.Field("roleId"), roleIdAnys...),
	).Find()
	if err != nil {
		return nil, err
	}

	for _, permissionLink := range permissionLinks {
		if permissionLink.(*user.PermissionRoleLink).PermissionId == permissionId {
			return usr, nil
		}

	}

	// check grants

	// @todo investigate why this doesn't work
	//
	// _, exists, err := s.grantsStore.Query(
	// 	datastore.Equals([]string{"permissionId"}, permissionId),
	// 	datastore.IsInList([]string{"slugs"}, usr.Slug),
	// ).FindOne()

	grantIs, err := s.grantsStore.Query(
		datastore.Equals([]string{"permissionId"}, permissionId),
	).Find()
	if err != nil {
		return nil, err
	}

	exists := false
	for _, grantI := range grantIs {
		if exists {
			break
		}
		grant := grantI.(*user.Grant)
		for _, slug := range grant.Slugs {
			if slug == usr.Slug {
				exists = true
				break
			}
		}
	}

	if exists {
		return usr, nil
	}

	return nil, errors.New("unauthorized")
}

func (s *UserService) getRoleIdsByUserId(userId string) ([]string, error) {
	roleLinks, err := s.userRoleLinksStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
	).Find()
	if err != nil {
		return nil, err
	}
	roleIds := []string{}
	for _, role := range roleLinks {
		roleIds = append(roleIds, role.(*user.UserRoleLink).RoleId)
	}

	return roleIds, nil
}

func (s *UserService) getUserFromRequest(r *http.Request) (*user.User, error) {
	authHeader := r.Header.Get("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	if authHeader == "" || authHeader == "Bearer" {
		return nil, fmt.Errorf("no auth header")
	}

	tokenI, found, err := s.authTokensStore.Query(
		datastore.Equals(datastore.Field("token"), authHeader),
	).FindOne()
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("token not found")
	}
	token := tokenI.(*user.AuthToken)

	userI, found, err := s.usersStore.Query(
		datastore.Id(token.UserId),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		logger.Error("Token refers to nonexistent user",
			slog.String("userId", token.UserId),
			slog.String("tokenId", token.Id),
		)
		return nil, errors.New("token user does not exist")
	}
	user := userI.(*user.User)
	return user, nil
}

func (s *UserService) GetUserFromRequest(
	request *http.Request,
) (*user.User, bool, error) {
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		return nil, false, nil
	}
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	tokenI, found, err := s.authTokensStore.Query(
		datastore.Equals(datastore.Field("token"), authHeader),
	).FindOne()
	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, errors.New("unauthorized")
	}
	token := tokenI.(*user.AuthToken)

	userI, found, err := s.usersStore.Query(
		datastore.Id(token.UserId),
	).FindOne()
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}

	user := userI.(*user.User)
	return user, found, nil
}
