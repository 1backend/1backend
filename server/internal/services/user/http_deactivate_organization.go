/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID deactivateOrganization
// @Summary Deactivate Organization
// @Description Clears the caller user's active organization for the current device and returns a fresh token without an active organization.
// @Tags User Svc
// @Produce json
// @Success 200 {object} user.DeactivateOrganizationResponse
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/organization/deactivate [post]
func (s *UserService) DeactivateOrganization(
	w http.ResponseWriter,
	r *http.Request,
) {
	claims, err := s.options.Authorizer.ParseJWTFromRequest(s.publicKeyPem, r)
	if err != nil {
		logger.Error(
			"Failed to parse JWT",
			slog.Any("error", err),
		)
		endpoint.Unauthorized(w)
		return
	}

	usr, err := s.readSelf(claims.UserId)
	if err != nil {
		logger.Error(
			"Failed to read self",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	token, err := s.deactivateOrganization(claims.AppId, usr, claims.Device)
	if err != nil {
		logger.Error(
			"Failed to deactivate organization",
			slog.Any("error", err),
			slog.String("userId", usr.Id),
		)
		endpoint.InternalServerError(w)
		return
	}

	if token == nil {
		logger.Error(
			"Failed to refresh token after deactivating organization",
			slog.String("userId", usr.Id),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.DeactivateOrganizationResponse{
		Token: *token,
	})
}
