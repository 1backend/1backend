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

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// @ID saveUser
// @Summary Save User
// @Description Save user information based on the provided user ID.
// @Description Intended for admins. Requires the `user-svc:user:edit` permission.
// @Description For a user to edit their own profile, see `saveSelf`.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param body body user.SaveUserRequest true "Save Profile Request"
// @Success 200 {object} user.SaveUserResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/user/{userId} [put]
func (s *UserService) SaveUser(w http.ResponseWriter, r *http.Request) {

	_, hasPermission, _, err := s.hasPermission(r, user.PermissionUserEdit)
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

	req := user.SaveUserRequest{}
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

	err = s.saveUser(userId, &req)
	if err != nil {
		logger.Error(
			"Failed to save user",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.SaveUserResponse{})
}

func (s *UserService) saveUser(
	userId string,
	request *user.SaveUserRequest,
) error {
	query := s.usersStore.Query(
		datastore.Equals(datastore.Field("id"), userId),
	)

	userI, found, err := query.FindOne()
	if err != nil {
		return err
	}

	if !found {
		return errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.ThumbnailFileId != "" {
		user.ThumbnailFileId = request.ThumbnailFileId
	}

	user.UpdatedAt = time.Now()

	query.Update(user)

	return nil
}
