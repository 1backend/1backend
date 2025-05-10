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

	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID listUsers
// @Summary List Users
// @Description Fetches a list of users with optional query filters and pagination.
// @Description Requires the `user-svc:user:view` permission that only admins have by default.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ListUsersRequest false "List Users Request"
// @Success 200 {object} user.ListUsersResponse "List of users retrieved successfully"
// @Failure 400 {object} user.ErrorResponse "invalid JSON"
// @Failure 401 {object} user.ErrorResponse "unauthorized"
// @Failure 500 {object} user.ErrorResponse "internal server error"
// @Security BearerAuth
// @Router /user-svc/users [post]
func (s *UserService) ListUsers(
	w http.ResponseWriter,
	r *http.Request,
) {

	_, hasPermission, err := s.hasPermission(r, user.PermissionUserView, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !hasPermission {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	req := user.ListUsersRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	users, count, err := s.listUsers(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.ListUsersResponse{
		Users: users,
		Count: count,
	})
	w.Write(bs)
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

		contact, found, err := s.contactsStore.Query(
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

	if request.ContactId != "" {
		contactIs, err := s.contactsStore.Query(
			datastore.Id(request.ContactId),
		).Find()
		if err != nil {
			return nil, 0, errors.Wrap(err, "error getting contact")
		}

		if len(contactIs) == 0 {
			return nil, 0, fmt.Errorf("cannot find contact with id '%v' when querying users", request.ContactId)
		}

		filters = append(filters, datastore.Equals(
			[]string{"id"}, contactIs[0].(*user.Contact).UserId,
		))
	}

	q := s.usersStore.Query(
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
			ContactIds: contactIds,
		})
	}

	return users, count, nil
}
