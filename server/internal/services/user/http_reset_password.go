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
	"time"

	"github.com/gorilla/mux"
	"github.com/openorch/openorch/sdk/go/datastore"
	user "github.com/openorch/openorch/server/internal/services/user/types"
	usertypes "github.com/openorch/openorch/server/internal/services/user/types"
)

// @ID resetPassword
// @Summary Reset Password
// @Description Allows an administrator to change a user's password.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param body body user.ResetPasswordRequest true "Change Password Request"
// @Success 200 {object} user.ResetPasswordResponse "Password changed successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/{userId}/reset-password [post]
func (s *UserService) ResetPassword(
	w http.ResponseWriter,
	r *http.Request,
) {

	_, err := s.isAuthorized(r, user.PermissionUserPasswordChange.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	vars := mux.Vars(r)
	userId := vars["userId"]

	req := user.ResetPasswordRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = s.resetPassword(userId, req.NewPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.ResetPasswordResponse{})
	w.Write(bs)
}

func (s *UserService) resetPassword(userId, newPassword string) error {
	q := s.usersStore.Query(
		datastore.Id(userId),
	)
	userI, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	newPasswordHash, err := s.hashPassword(newPassword)
	if err != nil {
		return err
	}

	user.PasswordHash = newPasswordHash
	user.UpdatedAt = time.Now()

	return q.Update(user)
}
