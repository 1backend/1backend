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
	"log/slog"
	"net/http"
	"time"

	"github.com/pkg/errors"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/gorilla/mux"
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

	_, hasPermission, _, err := s.hasPermission(r, user.PermissionUserPasswordReset)
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

	vars := mux.Vars(r)
	userId := vars["userId"]

	req := user.ResetPasswordRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	err = s.resetPassword(userId, req.NewPassword)
	if err != nil {
		logger.Error(
			"Failed to reset password",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.ResetPasswordResponse{})
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

	passwordIs, err := s.passwordsStore.Query(
		datastore.Equals(
			datastore.Field("userId"), user.Id),
	).OrderBy(
		datastore.OrderByField("createdAt", true),
	).Limit(1).Find()
	if err != nil {
		return errors.Wrap(err, "failed to get password")
	}
	if len(passwordIs) == 0 {
		return errors.New("user password not found to update")
	}

	err = s.passwordsStore.Upsert(&usertypes.Password{
		Id:           sdk.Id("pw"),
		PasswordHash: newPasswordHash,
		UserId:       user.Id,
		CreatedAt:    time.Now(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save password")
	}

	return nil
}
