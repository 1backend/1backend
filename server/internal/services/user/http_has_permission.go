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

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/gorilla/mux"

	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID hasPermission
// @Summary Has Permission
// @Description Check whether the caller user has a specific permission.
// @Description Ideally, this endpoint should rarely be used, as the JWT token
// @Description already includes all user roles. Caching the `List Permissions` and `List Grants`
// @Description responses allows services to determine user authorization
// @Description without repeatedly calling this endpoint.
// @Desciption This endpoint also checks for grants.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param permission path string true "Permission"
// @Param body body user.HasPermissionRequest false "Is Authorized Request"
// @Success 200 {object} user.HasPermissionResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON or missing permission id"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Security BearerAuth
// @Router /user-svc/self/has/{permission} [post]
func (s *UserService) HasPermission(
	w http.ResponseWriter,
	r *http.Request,
) {

	req := &user.HasPermissionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	permission := vars["permission"]

	if permission == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`missing permission id`))
		return
	}

	if req == nil {
		req = &user.HasPermissionRequest{}
	}

	usr, isAuth, err := s.hasPermission(r, permission, req.GrantedSlugs, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(&user.HasPermissionResponse{
		Authorized: isAuth,
		User:       usr,
	})

	w.Write(bs)
}

func (s *UserService) hasPermission(
	r *http.Request,
	permission string,
	grantedSlugs,
	contactsGranted []string,
) (*user.User, bool, error) {
	usr, err := s.getUserFromRequest(r)
	if err != nil {
		return nil, false, err
	}

	slugGrant := false
	for _, v := range grantedSlugs {
		if usr.Slug == v {
			slugGrant = true
		}
	}
	if slugGrant {
		return usr, true, nil
	}
	invites, err := s.invitesStore.Query(
		datastore.Equals(datastore.Field("userId"), usr.Id),
	).Find()
	if err != nil {

		return nil, false, err
	}

	roleIds := []string{}
	for _, role := range invites {
		roleIds = append(roleIds, role.(*user.Invite).Role)
	}

	roleIdAnys := []any{}
	for _, roleId := range roleIds {
		roleIdAnys = append(roleIdAnys, roleId)
	}
	grants, err := s.grantsStore.Query(
		datastore.Intersects(datastore.Field("roles"), roleIdAnys),
	).Find()
	if err != nil {
		return nil, false, err
	}

	for _, permissionLink := range grants {
		if permissionLink.(*user.Grant).Permission == permission {
			return usr, true, nil
		}

	}

	// check grants

	// @todo investigate why this doesn't work
	//
	// _, exists, err := s.grantsStore.Query(
	// 	datastore.Equals([]string{"permission"}, permissionId),
	// 	datastore.IsInList([]string{"slugs"}, usr.Slug),
	// ).FindOne()

	grantIs, err := s.grantsStore.Query(
		datastore.Equals([]string{"permission"}, permission),
	).Find()
	if err != nil {
		return nil, false, err
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

		for _, userRoleId := range roleIds {
			for _, grantRoleIds := range grant.Roles {

				if userRoleId == grantRoleIds {
					exists = true
					break
				}
			}
			if exists {
				break
			}
		}
	}

	if exists {
		return usr, true, nil
	}

	return nil, false, nil
}

func (s *UserService) getRolesByUserId(userId string) ([]string, error) {
	contactIs, err := s.contactsStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
	).Find()
	if err != nil {
		return nil, err
	}
	contactIds := []any{}
	for _, contactI := range contactIs {
		contactIds = append(contactIds, contactI.(*user.Contact).Id)
	}

	invites, err := s.invitesStore.Query(
		datastore.Or(
			datastore.Equals(datastore.Field("userId"), userId),
			datastore.IsInList(datastore.Field("contactId"), contactIds...),
		),
	).Find()
	if err != nil {
		return nil, err
	}

	roles := []string{}
	for _, invite := range invites {
		roles = append(roles, invite.(*user.Invite).Role)
	}

	return roles, nil
}

func (s *UserService) getContactIdsByUserId(userId string) ([]string, error) {
	contacts, err := s.contactsStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
	).Find()
	if err != nil {
		return nil, err
	}

	contactIds := []string{}
	for _, role := range contacts {
		contactIds = append(contactIds, role.(*user.Contact).Id)
	}

	return contactIds, nil
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
