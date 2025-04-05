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

	_, isAuthorized, err := s.hasPermission(r, user.PermissionUserView, nil, nil)
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

	req := user.ListUsersRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	if req.Query == nil {
		req.Query = &datastore.Query{}
	}
	if req.Query.Limit == 0 {
		req.Query.Limit = 20
	}

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
	filters = append(filters, request.Query.Filters...)

	if request.UserId != "" {
		filters = append(filters, datastore.Equals(
			[]string{"id"}, request.UserId,
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
	).Limit(request.Query.Limit)

	if len(request.Query.OrderBys) > 0 {
		q = q.OrderBy(request.Query.OrderBys...)
	} else {
		q = q.OrderBy(datastore.OrderByField("createdAt", true))
	}

	if request.Query.JSONAfter != "" {
		v := []any{}
		err := json.Unmarshal([]byte(request.Query.JSONAfter), &v)
		if err != nil {
			return nil, 0, err
		}
		q = q.After(v...)
	}

	res, err := q.Find()
	if err != nil {
		return nil, 0, err
	}

	var count int64
	if request.Query.Count {
		var err error
		count, err = q.Count()
		if err != nil {
			return nil, 0, err
		}
	}

	users := []*user.UserRecord{}
	for _, v := range res {
		usr := v.(*user.User)

		roles, err := s.getRolesByUserId(usr.Id)
		if err != nil {
			return nil, 0, err
		}

		contactIds, err := s.getContactIdsByUserId(usr.Id)
		if err != nil {
			return nil, 0, err
		}

		users = append(users, &user.UserRecord{
			Id:         usr.Id,
			Slug:       usr.Slug,
			CreatedAt:  usr.CreatedAt,
			UpdatedAt:  usr.UpdatedAt,
			Roles:      roles,
			ContactIds: contactIds,
		})
	}

	return users, count, nil
}
