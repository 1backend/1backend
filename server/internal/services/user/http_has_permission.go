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
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/gorilla/mux"

	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID hasPermission
// @Summary Has Permission
// @Description Checks whether the caller has a specific permission.
// @Description Optimized for caching — only the caller and the permission are required.
// @Description To assign a permission to a user or role, use the `Save Permits` endpoint.
// @Description
// @Description This endpoint does not return 401 Unauthorized if access is denied.
// @Description Instead, it always returns 200 OK with `Authorized: false` if the permission is missing.
// @Description The response will still include the caller’s user information if not authorized.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param permission path string true "Permission"
// @Success 200 {object} user.HasPermissionResponse
// @Failure 400 {object} user.ErrorResponse "Missing Permission"
// @Security BearerAuth
// @Router /user-svc/self/has/{permission} [post]
func (s *UserService) HasPermission(
	w http.ResponseWriter,
	r *http.Request,
) {
	vars := mux.Vars(r)
	permission := vars["permission"]

	if permission == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "Missing Permission")
		return
	}

	usr, hasPermission, claims, err := s.hasPermission(r, permission)
	if err != nil {
		logger.Error(
			"Failed to check permission",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	rsp := &user.HasPermissionResponse{
		Authorized: hasPermission,
		Until:      claims.ExpiresAt.Time,
		User:       *usr,
	}

	endpoint.WriteJSON(w, http.StatusOK, rsp)
}

func (s *UserService) hasPermission(
	r *http.Request,
	permission string,
) (*user.User, bool, *auth.Claims, error) {
	usr, claims, err := s.getUserFromRequest(r)
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, false, claims, nil
		}
		return nil, false, claims, errors.Wrap(err, "failed to get user from request")
	}

	enrolls, err := s.enrollsStore.Query(
		datastore.Equals(datastore.Field("userId"), usr.Id),
	).Find()
	if err != nil {
		return usr, false, claims, errors.Wrap(err, "failed to query enrolls")
	}

	roleIds := []string{}
	for _, role := range enrolls {
		roleIds = append(roleIds, role.(*user.Enroll).Role)
	}

	roleIdAnys := []any{}
	for _, roleId := range roleIds {
		roleIdAnys = append(roleIdAnys, roleId)
	}
	permits, err := s.permitsStore.Query(
		datastore.Intersects(datastore.Field("roles"), roleIdAnys),
	).Find()
	if err != nil {
		return usr, false, claims, err
	}

	for _, permissionLink := range permits {
		if permissionLink.(*user.Permit).Permission == permission {
			return usr, true, claims, nil
		}

	}

	// check permits

	// @todo investigate why this doesn't work
	//
	// _, exists, err := s.permitsStore.Query(
	// 	datastore.Equals([]string{"permission"}, permissionId),
	// 	datastore.IsInList([]string{"slugs"}, usr.Slug),
	// ).FindOne()

	permitIs, err := s.permitsStore.Query(
		datastore.Equals([]string{"permission"}, permission),
	).Find()
	if err != nil {
		return usr, false, claims, err
	}

	exists := false
	for _, permitI := range permitIs {
		if exists {
			break
		}
		permit := permitI.(*user.Permit)
		for _, slug := range permit.Slugs {
			if slug == usr.Slug {
				exists = true
				break
			}
		}

		for _, userRoleId := range roleIds {
			for _, permitRoleIds := range permit.Roles {

				if userRoleId == permitRoleIds {
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
		return usr, true, claims, nil
	}

	return usr, false, claims, nil
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

	enrolls, err := s.enrollsStore.Query(
		datastore.Or(
			datastore.Equals(datastore.Field("userId"), userId),
			datastore.IsInList(datastore.Field("contactId"), contactIds...),
		),
	).Find()
	if err != nil {
		return nil, err
	}

	roles := []string{}
	for _, enroll := range enrolls {
		roles = append(roles, enroll.(*user.Enroll).Role)
	}

	return roles, nil
}

func (s *UserService) getContactIdsByUserId(userId string) ([]string, error) {
	contacts, err := s.getContactsByUserId(userId)
	if err != nil {
		return nil, err
	}

	contactIds := []string{}
	for _, contact := range contacts {
		contactIds = append(contactIds, contact.Id)
	}

	return contactIds, nil
}

func (s *UserService) getContactsByUserId(userId string) ([]user.Contact, error) {
	contactIs, err := s.contactsStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
	).Find()
	if err != nil {
		return nil, err
	}

	contacts := []user.Contact{}
	for _, contact := range contactIs {
		contacts = append(contacts, *contact.(*user.Contact))
	}

	return contacts, nil
}

func (s *UserService) getUserFromRequest(r *http.Request) (
	*user.User,
	*auth.Claims,
	error,
) {
	authHeader := r.Header.Get("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	if authHeader == "" || authHeader == "Bearer" {
		return nil, nil, fmt.Errorf("no auth header")
	}

	claims, err := s.options.Authorizer.ParseJWTFromRequest(s.publicKeyPem, r)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse JWT from request")
	}

	tokenI, found, err := s.authTokensStore.Query(
		datastore.Equals(datastore.Field("token"), authHeader),
	).FindOne()
	if err != nil {
		return nil, claims, err
	}

	if !found {
		return nil, claims, errors.New("token not found")
	}

	token := tokenI.(*user.AuthToken)

	userI, found, err := s.usersStore.Query(
		datastore.Id(token.UserId),
	).FindOne()
	if err != nil {
		return nil, claims, err
	}
	if !found {
		logger.Error("Token refers to nonexistent user",
			slog.String("userId", token.UserId),
			slog.String("tokenId", token.Id),
		)
		return nil, claims, errors.New("token user does not exist")
	}

	usr := userI.(*user.User)

	return usr, claims, nil
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
