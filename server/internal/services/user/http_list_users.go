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
	"fmt"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID listUsers
// @Summary List Users
// @Description Fetches a list of users with optional query filters and pagination.
// @Description Requires the `user-svc:user:view` permission that only admins have by default for unscoped list requests.
// @Description Requests scoped by `contactId` or `contactIds` do not require the broad user-view permission.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ListUsersRequest false "List Users Request"
// @Success 200 {object} user.ListUsersResponse "List of users retrieved successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/users [post]
func (s *UserService) ListUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	req := user.ListUsersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	if !hasContactIdFilter(&req) {
		_, hasPermission, _, err := s.hasPermission(r, user.PermissionUserView)
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
	}

	users, count, err := s.listUsers(&req)
	if err != nil {
		logger.Error(
			"Failed to list users",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(user.ListUsersResponse{
		Users: users,
		Count: count,
	})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (s *UserService) listUsers(
	request *user.ListUsersRequest,
) ([]*user.UserRecord, int64, error) {
	filters := []datastore.Filter{}

	if request.Search != "" {
		ors := []datastore.Filter{
			datastore.Equals(
				[]string{"slug"}, request.Search,
			),
			datastore.Equals(
				[]string{"name"}, request.Search,
			),
		}

		contact, found, err := s.contactStore.Query(
			datastore.Id(request.Search),
		).FindOne()
		if err != nil {
			return nil, 0, errors.Wrap(err, "error getting contact")
		}

		if found {
			ors = append(ors, datastore.Equals(
				[]string{"id"}, contact.(*user.Contact).UserId,
			))
		}

		filters = append(filters, datastore.Or(ors...))
	}

	if request.Ids != nil {
		ids := []any{}
		for _, id := range request.Ids {
			ids = append(ids, id)
		}
		filters = append(filters, datastore.IsInList(
			[]string{"id"}, ids...,
		))
	}

	contactIds := uniqueContactIds(request)
	if len(contactIds) > 0 {
		contactIdAnys := make([]any, 0, len(contactIds))
		for _, contactId := range contactIds {
			contactIdAnys = append(contactIdAnys, contactId)
		}

		contactIs, err := s.contactStore.Query(
			datastore.IsInList(datastore.Field("id"), contactIdAnys...),
		).Find()
		if err != nil {
			return nil, 0, errors.Wrap(err, "error getting contact")
		}

		foundContactIds := map[string]struct{}{}
		userIds := make([]any, 0, len(contactIs))
		seenUserIds := map[string]struct{}{}
		for _, contactI := range contactIs {
			contact := contactI.(*user.Contact)
			foundContactIds[contact.Id] = struct{}{}
			if _, ok := seenUserIds[contact.UserId]; ok {
				continue
			}
			seenUserIds[contact.UserId] = struct{}{}
			userIds = append(userIds, contact.UserId)
		}

		for _, contactId := range contactIds {
			if _, ok := foundContactIds[contactId]; !ok {
				return nil, 0, fmt.Errorf("cannot find contact with id '%v' when querying users", contactId)
			}
		}

		filters = append(filters, datastore.IsInList(
			[]string{"id"}, userIds...,
		))
	}

	q := s.userStore.Query(
		filters...,
	)

	if request.OrderBy == "" {
		request.OrderBy = user.ListUsersOrderByCreatedAt
	}

	if request.Order == "" {
		request.Order = user.OrderDirectionDesc
	}

	q = q.OrderBy(
		datastore.OrderByField(
			string(request.OrderBy),
			request.Order == user.OrderDirectionDesc),
	)

	if !request.AfterTime.IsZero() {
		q = q.After(request.AfterTime)
	}

	if request.Limit != 0 {
		q = q.Limit(int64(request.Limit))
	} else {
		q = q.Limit(20)
	}

	res, err := q.Find()
	if err != nil {
		return nil, 0, err
	}

	userIds := make([]string, 0, len(res))
	for _, v := range res {
		userIds = append(userIds, v.(*user.User).Id)
	}
	totpEnabledByUserId, err := s.enabledTOTPByUserId(userIds)
	if err != nil {
		return nil, 0, err
	}

	var count int64
	if request.Count {
		var err error
		count, err = q.Count()
		if err != nil {
			return nil, 0, err
		}
	}

	users := []*user.UserRecord{}
	for _, v := range res {
		usr := v.(*user.User)

		// roles, err := s.getRolesByUserId(usr.Id)
		// if err != nil {
		// 	return nil, 0, err
		// }

		contactIds, err := s.getContactIdsByUserId(usr.Id)
		if err != nil {
			return nil, 0, err
		}

		users = append(users, &user.UserRecord{
			Id:        usr.Id,
			Slug:      usr.Slug,
			Name:      usr.Name,
			CreatedAt: usr.CreatedAt,
			UpdatedAt: usr.UpdatedAt,
			// Roles:      roles,
			ContactIds:  contactIds,
			TOTPEnabled: totpEnabledByUserId[usr.Id],
		})
	}

	return users, count, nil
}

func (s *UserService) enabledTOTPByUserId(userIds []string) (map[string]bool, error) {
	enabledByUserId := map[string]bool{}

	userIdAnys := make([]any, 0, len(userIds))
	seen := map[string]struct{}{}
	for _, userId := range userIds {
		if userId == "" {
			continue
		}
		if _, ok := seen[userId]; ok {
			continue
		}
		seen[userId] = struct{}{}
		userIdAnys = append(userIdAnys, userId)
	}
	if len(userIdAnys) == 0 {
		return enabledByUserId, nil
	}

	totpIs, err := s.totpStore.Query(
		datastore.IsInList(datastore.Field("userId"), userIdAnys...),
		datastore.Equals(datastore.Field("enabled"), true),
	).Find()
	if err != nil {
		return nil, errors.Wrap(err, "error querying enabled TOTP records")
	}

	for _, totpI := range totpIs {
		totp := totpI.(*user.TOTP)
		enabledByUserId[totp.UserId] = true
	}

	return enabledByUserId, nil
}

func hasContactIdFilter(request *user.ListUsersRequest) bool {
	if request.ContactId != "" {
		return true
	}

	for _, contactId := range request.ContactIds {
		if contactId != "" {
			return true
		}
	}

	return false
}

func uniqueContactIds(request *user.ListUsersRequest) []string {
	contactIds := []string{}
	seen := map[string]struct{}{}

	for _, contactId := range append([]string{request.ContactId}, request.ContactIds...) {
		if contactId == "" {
			continue
		}
		if _, ok := seen[contactId]; ok {
			continue
		}
		seen[contactId] = struct{}{}
		contactIds = append(contactIds, contactId)
	}

	return contactIds
}
