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

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

var (
	cannotListUnowned = errors.New("cannot list enrolls for unowned role")
)

// @ID listEnrolls
// @Summary List Enrolls
// @Description List enrolls. Role, user ID or contact ID must be specified.
// @Description
// @Description Requires the `user-svc:enroll:view` permission, which by default all users have.
// @Description Caller can only list enrolls of roles they own (unless they are an admin).
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ListEnrollsRequest true "List Enrolls Request"
// @Success 200 {object} user.ListEnrollsResponse "Enrolls listed successfully"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 400 {object} user.ErrorResponse "Role, Contact ID or User ID is Required"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/enrolls [post]
func (s *UserService) ListEnrolls(w http.ResponseWriter, r *http.Request) {

	_, hasPermission, claims, err := s.hasPermission(r, user.PermissionEnrollView)
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
	_, hasUserViewPermission, _, err := s.hasPermission(r, user.PermissionUserView)
	if err != nil {
		logger.Error(
			"Failed to check user view permission",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	req := &user.ListEnrollsRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	if req.Role == "" && req.ContactId == "" && req.UserId == "" {
		logger.Error(
			"Role, contact ID or user ID is required",
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Role, Contact ID or User ID is Required")
		return
	}

	roles, err := s.getRolesByUserId(claims.AppId, claims.UserId)
	if err != nil {
		logger.Error("Failed to list effective roles", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	effectiveClaims := *claims
	effectiveClaims.Roles = roles

	enrolls, err := s.listEnrolls(&effectiveClaims, req, hasUserViewPermission)
	switch {
	case errors.Is(err, cannotListUnowned):
		endpoint.Unauthorized(w)
		return
	case err != nil:
		logger.Error(
			"Failed to list enrolls",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	rsp := user.ListEnrollsResponse{
		Enrolls: enrolls,
	}
	endpoint.WriteJSON(w, http.StatusOK, rsp)
}

func (s *UserService) listEnrolls(
	claims *auth.Claims,
	req *user.ListEnrollsRequest,
	canViewAllUserContactIds bool,
) ([]user.Enroll, error) {

	isAdmin := lo.Contains(claims.Roles, user.RoleAdmin)

	filters := []datastore.Filter{}

	if req.Role != "" {
		if !isAdmin {
			if !auth.OwnsRole(claims, req.Role) {
				return nil, cannotListUnowned
			}
		}
		filters = append(filters, datastore.Equals([]string{"role"}, req.Role))
	}
	if req.ContactId != "" {
		if !isAdmin {
			// @todo let users list for their own contacts
			return nil, errors.New("only admins can list based on contact id")
		}
		filters = append(filters, datastore.Equals([]string{"contactId"}, req.ContactId))
	}
	if req.UserId != "" {
		if !isAdmin && req.UserId != claims.UserId {
			return nil, errors.New("cannot list for other user ids")
		}
		filters = append(filters, datastore.Equals([]string{"userId"}, req.UserId))
	}

	enrollIs, err := s.enrollStore.Query(filters...).Find()
	if err != nil {
		return nil, errors.Wrap(err, "error querying enrolls")
	}

	enrolls := []user.Enroll{}
	for _, enrollI := range enrollIs {
		e := *enrollI.(*user.Enroll)

		if isAdmin || auth.OwnsRole(claims, e.Role) {
			enrolls = append(enrolls, e)
		}
	}

	if err := s.hydrateEnrollContactIds(
		enrolls,
		visibleContactUserIds(claims, enrolls, isAdmin || canViewAllUserContactIds),
	); err != nil {
		return nil, err
	}

	return enrolls, nil
}

func visibleContactUserIds(
	claims *auth.Claims,
	enrolls []user.Enroll,
	canViewAll bool,
) map[string]struct{} {
	userIds := map[string]struct{}{}
	for _, enroll := range enrolls {
		if enroll.UserId == "" {
			continue
		}
		if canViewAll || (claims != nil && enroll.UserId == claims.UserId) {
			userIds[enroll.UserId] = struct{}{}
		}
	}
	return userIds
}

func (s *UserService) hydrateEnrollContactIds(
	enrolls []user.Enroll,
	visibleUserContactIds map[string]struct{},
) error {
	userIdSet := map[string]struct{}{}
	for _, enroll := range enrolls {
		if enroll.UserId == "" {
			continue
		}
		if _, ok := visibleUserContactIds[enroll.UserId]; !ok {
			continue
		}
		userIdSet[enroll.UserId] = struct{}{}
	}

	contactsByUserId := map[string][]string{}
	if len(userIdSet) > 0 {
		userIds := make([]any, 0, len(userIdSet))
		for userId := range userIdSet {
			userIds = append(userIds, userId)
		}

		contactIs, err := s.contactStore.Query(
			datastore.IsInList(datastore.Field("userId"), userIds...),
		).Find()
		if err != nil {
			return errors.Wrap(err, "error querying contacts for enrolls")
		}

		for _, contactI := range contactIs {
			contact := contactI.(*user.Contact)
			contactsByUserId[contact.UserId] = append(contactsByUserId[contact.UserId], contact.Id)
		}
	}

	for i := range enrolls {
		contactIds := []string{}
		seen := map[string]struct{}{}

		if enrolls[i].ContactId != "" {
			contactIds = append(contactIds, enrolls[i].ContactId)
			seen[enrolls[i].ContactId] = struct{}{}
		}

		for _, contactId := range contactsByUserId[enrolls[i].UserId] {
			if _, ok := seen[contactId]; ok {
				continue
			}
			contactIds = append(contactIds, contactId)
		}

		enrolls[i].ContactIds = contactIds
	}

	return nil
}
