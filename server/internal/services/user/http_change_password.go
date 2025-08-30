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
	"time"

	"github.com/pkg/errors"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID changePassword
// @Summary Change Password
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

	usr, hasPermission, _, err := s.hasPermission(r, user.PermissionUserPasswordChange)
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

	req := user.ChangePasswordRequest{}
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

	err = s.changePassword(usr.Id, req.CurrentPassword, req.NewPassword)
	if err != nil {
		logger.Error(
			"Failed to change password",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.ChangePasswordResponse{})
}

func (s *UserService) changePassword(
	userId, currentPassword, newPassword string,
) error {
	q := s.passwordStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
	).OrderBy(
		datastore.OrderByField("createdAt", true),
	).Limit(1)

	passwordI, found, err := q.FindOne()
	if err != nil {
		return err
	}
	if !found {
		return errors.New("user not found")
	}
	password := passwordI.(*usertypes.Password)

	if !checkPasswordHash(currentPassword, password.PasswordHash) {
		return errors.New("current password is incorrect")
	}

	newPasswordHash, err := s.hashPassword(newPassword)
	if err != nil {
		return err
	}

	err = s.passwordStore.Upsert(&usertypes.Password{
		Id:           password.Id,
		PasswordHash: newPasswordHash,
		UserId:       password.UserId,
		CreatedAt:    time.Now(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save password")
	}

	return nil
}
