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
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID createUser
// @Summary Create a New User
// @Description Allows an authenticated administrator to create a new user with specified details.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.CreateUserRequest true "Create User Request"
// @Success 200 {object} user.CreateUserResponse "User created successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 400 {object} user.ErrorResponse "Invalid User"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/user [post]
func (s *UserService) CreateUser(
	w http.ResponseWriter,
	r *http.Request) {

	_, hasPermission, err := s.hasPermission(r, user.PermissionUserCreate)
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

	req := user.CreateUserRequest{}
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

	err = validateUser(req.User, req.Contacts)
	if err != nil {
		logger.Error(
			"Failed to validate user",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid User")
		return
	}

	err = s.createUser(req.User, req.Contacts, req.Password, req.RoleIds)
	if err != nil {
		logger.Error(
			"Failed to create user",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(user.CreateUserResponse{})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func validateUser(u *user.UserInput, contacts []user.Contact) error {
	if u == nil {
		return errors.New("user is required")
	}
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Slug == "" {
		return errors.New("slug is required")
	}
	if len(contacts) > 1 {
		return errors.New("more than one contact")
	}

	return nil
}
