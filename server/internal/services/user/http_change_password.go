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

	"github.com/openorch/openorch/sdk/go/datastore"
	user "github.com/openorch/openorch/server/internal/services/user/types"
	usertypes "github.com/openorch/openorch/server/internal/services/user/types"
)

// @ID changePassword
// @Summary Change User Password
// @Description Allows an authenticated user to change their own password.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ChangePasswordRequest true "Change Password Request"
// @Success 200 {object} user.ChangePasswordResponse "Password changed successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/change-password [post]
func (s *UserService) ChangePassword(w http.ResponseWriter, r *http.Request) {

	_, err := s.isAuthorized(r, user.PermissionUserPasswordChange.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	req := user.ChangePasswordRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = s.changePassword(req.Slug, req.CurrentPassword, req.NewPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.ChangePasswordResponse{})
	w.Write(bs)
}

func (s *UserService) changePassword(
	slug, currentPassword, newPassword string,
) error {
	q := s.usersStore.Query(
		datastore.Equals(datastore.Field("slug"), slug),
	)
	userI, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	if !checkPasswordHash(currentPassword, user.PasswordHash) {
		return errors.New("current password is incorrect")
	}

	newPasswordHash, err := s.hashPassword(newPassword)
	if err != nil {
		return err
	}
	user.PasswordHash = newPasswordHash
	user.UpdatedAt = time.Now()

	return q.Update(user)
}
