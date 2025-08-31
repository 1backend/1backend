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
	"github.com/pkg/errors"
)

// @ID saveSelf
// @Summary Save User Profile
// @Description Save user's own profile information.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.SaveSelfRequest true "Save Profile Request"
// @Success 200 {object} user.SaveSelfResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/self [put]
func (s *UserService) SaveSelf(w http.ResponseWriter, r *http.Request) {
	stringToken, exists := s.options.Authorizer.TokenFromRequest(r)
	if !exists {
		endpoint.WriteString(w, http.StatusBadRequest, "Token Missing")
		return
	}

	token, err := s.refreshToken(stringToken)
	if err != nil {
		logger.Error(
			"Failed to refresh token",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)

		return
	}

	usr, err := s.readSelf(token.UserId)
	if err != nil {
		logger.Error(
			"Failed to read self",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	req := user.SaveSelfRequest{}
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

	// cannot change slug for now
	err = s.saveSelf(usr, &req)
	if err != nil {
		logger.Error(
			"Failed to save self",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(user.SaveSelfResponse{})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (s *UserService) saveSelf(
	usr *user.User,
	request *user.SaveSelfRequest,
) error {

	if request.Name != "" {
		usr.Name = request.Name
	}

	if request.ThumbnailFileId != "" {
		usr.ThumbnailFileId = request.ThumbnailFileId
	}

	if request.Labels != nil {
		if usr.Labels == nil {
			usr.Labels = map[string]string{}
		}
		for k, v := range *request.Labels {
			usr.Labels[k] = v
		}
	}

	usr.UpdatedAt = time.Now()

	err := s.userStore.Query(
		datastore.Id(usr.Id),
	).Update(usr)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}
