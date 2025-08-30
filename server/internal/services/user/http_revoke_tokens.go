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

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID revokeTokens
// @Summary Revoke Tokens
// @Description Revoke tokens in one of the following scenarios:
// @Description - For the current user.
// @Description - For another user (see `userId` field), if permitted (`user-svc:token:revoke` permission, typically by admins).
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.RevokeTokensRequest false "Revoke Tokens Request"
// @Success 200 {object} user.RevokeTokensResponse
// @Failure 400 {object} user.ErrorResponse "Token Missing"
// @Failure 400 {object} user.ErrorResponse "Mutually Exclusive Parameters"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/tokens [delete]
func (s *UserService) RevokeTokens(w http.ResponseWriter, r *http.Request) {
	_, hasTokenRevokePermission, claims, err := s.hasPermission(
		r,
		user.PermissionTokenRevoke,
	)
	if err != nil {
		logger.Error(
			"Failed to check permission",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	request := user.RevokeTokensRequest{}
	if r.ContentLength != 0 {
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			logger.Error(
				"Failed to decode request",
				slog.Any("error", err),
			)
			endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
			return
		}
		defer r.Body.Close()
	}

	if !hasTokenRevokePermission && request.UserId != "" {
		endpoint.Unauthorized(w)
		return
	}

	if !hasTokenRevokePermission {
		request.UserId = claims.UserId
	}

	if request.UserId == "" && !request.AllTokens ||
		request.UserId != "" && request.AllTokens {
		endpoint.WriteString(w, http.StatusBadRequest, "Mutually Exclusive Parameters")
		return
	}

	err = s.revokeTokens(&request)
	if err != nil {
		logger.Error(
			"Failed to Revoke Tokens",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	rsp := user.RevokeTokensResponse{}

	endpoint.WriteJSON(w, http.StatusOK, rsp)
}

func (s *UserService) revokeTokens(request *user.RevokeTokensRequest) error {
	filters := []datastore.Filter{}

	if request.UserId != "" {
		filters = append(filters, datastore.Equals(datastore.Field("userId"), request.UserId))
	}

	if request.Device != "" {
		filters = append(filters, datastore.Equals(datastore.Field("device"), request.Device))
	}

	err := s.tokenStore.Query(
		filters...,
	).Delete()
	if err != nil {
		return err
	}

	return nil
}
