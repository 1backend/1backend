/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
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
// @Failure 422 {object} user.ErrorResponse "No Auth Header"
// @Security BearerAuth
// @Router /user-svc/self/has/{permission} [post]
func (s *UserService) HasPermission(
	w http.ResponseWriter,
	r *http.Request,
) {
	vars := mux.Vars(r)
	rawPermission := vars["permission"]

	if rawPermission == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "Missing Permission")
		return
	}

	permission, err := url.PathUnescape(rawPermission)
	if err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid Permission Format")
		return
	}

	authHeader := r.Header.Get("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	if authHeader == "" || authHeader == "Bearer" {
		endpoint.WriteErr(w, http.StatusUnprocessableEntity, errors.New("No Auth Header"))
		return
	}

	usr, hasPermission, claims, err := s.hasPermission(r, permission)
	if err != nil {
		// TODO: Previously this check was only `err != nil`. If a user simply lacked
		// permission (e.g., unauthenticated), `claims` could be nil and the code below
		// would panic with a nil pointer. That panic resulted in a 500 response.
		logger.Error(
			"Failed to check permission",
			slog.String("permission", permission),
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	// TODO: Previously this check was only `err != nil`. If a user simply lacked
	// permission (e.g., unauthenticated), `claims` could be nil and the code below
	// would panic with a nil pointer. That panic resulted in a 500 response.
	if claims == nil {
		endpoint.InternalServerError(w)
		return
	}

	app, found, err := s.appByID(claims.AppId)
	if err != nil {
		logger.Error(
			"Failed to query app",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}
	if !found {
		logger.Error(
			"App from token not found",
			slog.String("appId", claims.AppId),
			slog.String("userId", claims.UserId),
		)
		endpoint.InternalServerError(w)
		return
	}

	rsp := &user.HasPermissionResponse{
		Authorized: hasPermission,
		AppId:      claims.AppId,
		Until:      claims.ExpiresAt.Time,
		User:       *usr,
		App:        *app,
	}

	endpoint.WriteJSON(w, http.StatusOK, rsp)
}

func (s *UserService) hasPermission(
	r *http.Request,
	permission string,
) (*user.User, bool, *auth.Claims, error) {
	usr, claims, err := s.getUserFromRequest(r)
	if err != nil {
		if isExpiredTokenError(err) {
			return usr, false, claims, nil
		}

		return nil, false, claims, errors.Wrap(err, "failed to get user from request")
	}

	roleValues := make([]any, 0, len(claims.Roles))
	for _, role := range claims.Roles {
		roleValues = append(roleValues, role)
	}

	subjectFilters := []datastore.Filter{
		datastore.Equals(datastore.Field("slugs"), claims.Slug),
	}
	if len(roleValues) > 0 {
		subjectFilters = append(subjectFilters, datastore.Intersects(datastore.Field("roles"), roleValues))
	}

	permitIs, err := s.permitStore.Query(
		datastore.Or(
			datastore.Equals(datastore.Field("appId"), claims.AppId),
			datastore.Equals(datastore.Field("appId"), "*"),
		),
		datastore.Or(
			datastore.Equals(datastore.Field("permission"), permission),
			datastore.Equals(datastore.Field("permissions"), permission),
		),
		datastore.Or(subjectFilters...),
	).Find()
	if err != nil {
		return usr, false, claims, err
	}

	if len(permitIs) > 0 {
		return usr, true, claims, nil
	}

	return usr, false, claims, nil
}

func (s *UserService) getRolesByUserId(appId, userId string) ([]string, error) {
	contactIs, err := s.contactStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
	).Find()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query contacts")
	}
	contactIds := []any{}
	for _, contactI := range contactIs {
		contactIds = append(contactIds, contactI.(*user.Contact).Id)
	}

	// An admin should be an admin in all apps
	// `user-svc:admin`s are platform admins, not app admins.
	// This is mostly to avoid bootstrapping issues.

	adminEnrolls, err := s.enrollStore.Query(
		datastore.Equals(datastore.Field("role"), "user-svc:admin"),
		datastore.Equals(datastore.Field("userId"), userId),
	).Find()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query admin enrolls")
	}

	enrolls, err := s.enrollStore.Query(
		datastore.Or(
			datastore.Equals(datastore.Field("appId"), appId),
			datastore.Equals(datastore.Field("appId"), "*"),
		),
		datastore.Or(
			datastore.Equals(datastore.Field("userId"), userId),
			datastore.IsInList(datastore.Field("contactId"), contactIds...),
		),
	).Find()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query enrolls")
	}

	rolesIndex := map[string]struct{}{}
	for _, enroll := range append(enrolls, adminEnrolls...) {
		rolesIndex[enroll.(*user.Enroll).Role] = struct{}{}
	}

	membershipIs, err := s.membershipStore.Query(
		datastore.Equals(datastore.Field("appId"), appId),
		datastore.Equals(datastore.Field("userId"), userId),
		datastore.Equals(datastore.Field("status"), user.MembershipStatusAccepted),
	).Find()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query memberships")
	}
	for _, membershipI := range membershipIs {
		membership := membershipI.(*user.Membership)
		rolesIndex[orgMemberRole(membership.OrganizationId, membership.UserId)] = struct{}{}
		for _, roleId := range membership.Roles {
			rolesIndex[roleId] = struct{}{}
		}
	}

	roles := []string{"user-svc:user"}
	for role := range rolesIndex {
		roles = append(roles, role)
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
	contactIs, err := s.contactStore.Query(
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

func (s *UserService) getUserByContactId(contactId string) (*user.User, bool, error) {
	if contactId == "" {
		return nil, false, fmt.Errorf("contact id is required")
	}

	contactI, found, err := s.contactStore.Query(
		datastore.Equals(datastore.Field("id"), contactId),
	).FindOne()
	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, nil
	}

	contact := contactI.(*user.Contact)

	userI, found, err := s.userStore.Query(
		datastore.Equals(datastore.Field("id"), contact.UserId),
	).FindOne()

	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, fmt.Errorf("database corrupted - contact '%s' found but no corresponding user for user id '%v'",
			contact.Id,
			contact.UserId,
		)
	}

	return userI.(*user.User), true, nil
}

func (s *UserService) getUserFromRequest(r *http.Request) (
	*user.User,
	*auth.Claims,
	error,
) {
	authHeader := r.Header.Get("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	// at this point we should definitely have an auth header but we still check
	if authHeader == "" || authHeader == "Bearer" {
		return nil, nil, fmt.Errorf("no auth header")
	}

	claims, err := s.options.Authorizer.ParseJWTFromRequest(s.publicKeyPem, r)
	tokenErr := error(nil)
	if err != nil {
		tokenErr = errors.Wrap(err, "failed to parse JWT from request")

		if isExpiredTokenError(err) {
			recoveredUsr, recoveredClaims, recoverErr := s.recoverExpiredRequestUser(r)
			if recoverErr == nil {
				return recoveredUsr, recoveredClaims, tokenErr
			}
			if recoveredClaims != nil {
				claims = recoveredClaims
			}
		}
		if claims == nil {
			return nil, nil, tokenErr
		}
	}

	// Do not read the token here:
	// exchanged tokens are not persisted so they won't be found.

	userI, found, err := s.userStore.Query(
		datastore.Id(claims.UserId),
	).FindOne()
	if err != nil {
		return nil, claims, err
	}
	if !found {
		logger.Error("Token refers to nonexistent user",
			slog.String("userId", claims.UserId),
			//slog.String("tokenId", token.Id),
		)
		return nil, claims, errors.New("token user does not exist")
	}

	usr := userI.(*user.User)

	if tokenErr != nil {
		return usr, claims, tokenErr
	}

	return usr, claims, nil
}

func (s *UserService) recoverExpiredRequestUser(r *http.Request) (*user.User, *auth.Claims, error) {
	tokenString, hasToken := s.options.Authorizer.TokenFromRequest(r)
	if !hasToken {
		return nil, nil, fmt.Errorf("no auth header")
	}

	claims, err := s.options.Authorizer.ParseJWTUnverified(tokenString)
	if err != nil {
		return nil, nil, err
	}

	userI, found, err := s.userStore.Query(
		datastore.Id(claims.UserId),
	).FindOne()
	if err != nil {
		return nil, claims, err
	}
	if !found {
		logger.Error("Expired token refers to nonexistent user",
			slog.String("userId", claims.UserId),
		)
		return nil, claims, errors.New("token user does not exist")
	}

	return userI.(*user.User), claims, nil
}

func isExpiredTokenError(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), "token is expired")
}

func isUnauthorizedRequestError(err error) bool {
	if err == nil {
		return false
	}

	msg := err.Error()

	return strings.Contains(msg, "no auth header") ||
		strings.Contains(msg, "failed to parse JWT from request") ||
		strings.Contains(msg, "token user does not exist")
}

func (s *UserService) GetUserFromRequest(
	request *http.Request,
) (*user.User, bool, error) {
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		return nil, false, nil
	}
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	tokenI, found, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("token"), authHeader),
	).FindOne()
	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, errors.New("unauthorized")
	}
	token := tokenI.(*user.Token)

	userI, found, err := s.userStore.Query(
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
