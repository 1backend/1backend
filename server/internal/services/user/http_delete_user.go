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
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/gorilla/mux"
)

// @ID deleteUser
// @Summary Delete a User
// @Description Delete a user based on the user ID.
// @Tags User Svc
// @Accept  json
// @Produce  json
// @Param   userId     path    string  true  "User ID"
// @Success 200 {object} user.DeleteUserResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/user/{userId} [delete]
func (s *UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {

	usr, hasPermission, err := s.hasPermission(r, user.PermissionUserDelete, nil, nil)
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

	callerUserId := usr.Id
	isAdmin, err := s.isAdmin(callerUserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAdmin {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	userId := mux.Vars(r)["userId"]

	err = s.deleteUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.DeleteUserResponse{})
	w.Write(bs)
}

func (s *UserService) deleteUser(userId string) error {
	if userId == "" {
		return errors.New("no user id")
	}
	q := s.usersStore.Query(
		datastore.Id(userId),
	)
	_, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}

	isAdminUser, err := s.isAdmin(userId)
	if err != nil {
		return err
	}

	if isAdminUser {
		adminUsers, err := s.enrollsStore.Query(
			datastore.Equals(datastore.Field("role"), usertypes.RoleAdmin),
		).Find()
		if err != nil {
			return err
		}
		if len(adminUsers) == 0 {
			return errors.New("cannot detect number of admin users")
		}
		if len(adminUsers) == 1 {
			return errors.New("Cannot delete last admin user")
		}
	}

	return q.Delete()
}

func (s *UserService) isAdmin(userId string) (bool, error) {
	_, isAdminUser, err := s.enrollsStore.Query(
		datastore.Equals([]string{"userId"}, userId),
		datastore.Equals([]string{"role"}, usertypes.RoleAdmin),
	).FindOne()
	if err != nil {
		return false, err
	}

	return isAdminUser, nil
}
