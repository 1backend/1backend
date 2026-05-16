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

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID revokeApiKeys
// @Summary Revoke API Keys
// @Description Revokes one or more API keys.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.RevokeApiKeysRequest true "Revoke API Keys Request"
// @Success 200 {object} user.RevokeApiKeysResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/api-keys [delete]
func (s *UserService) RevokeApiKeys(w http.ResponseWriter, r *http.Request) {
	_, claims, err := s.getUserFromRequest(r)
	if err != nil {
		writeAuthOrInternalError(w, err, "Failed to get user from request")
		return
	}

	request := user.RevokeApiKeysRequest{}
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

	targetUserId := request.UserId
	if targetUserId == "" {
		targetUserId = claims.UserId
	}
	if targetUserId != claims.UserId && !canManageAPIKeysOnBehalf(claims) {
		endpoint.Unauthorized(w)
		return
	}

	if err := s.revokeApiKeys(claims, &request); err != nil {
		if isBadApiKeyRequest(err) {
			endpoint.WriteString(w, http.StatusBadRequest, err.Error())
			return
		}

		logger.Error(
			"Failed to revoke API keys",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.RevokeApiKeysResponse{})
}

func (s *UserService) revokeApiKeys(
	claims *auth.Claims,
	request *user.RevokeApiKeysRequest,
) error {
	if claims == nil {
		return errors.New("claims are required")
	}
	if request == nil {
		return errors.New("request is required")
	}

	ids := make([]any, 0, len(request.Ids)+1)
	seen := map[string]bool{}
	if request.Id != "" {
		ids = append(ids, request.Id)
		seen[request.Id] = true
	}
	for _, id := range request.Ids {
		if id == "" || seen[id] {
			continue
		}
		ids = append(ids, id)
		seen[id] = true
	}
	if len(ids) == 0 {
		return errors.New("id or ids are required")
	}

	userId := request.UserId
	if userId == "" {
		userId = claims.UserId
	}

	now := time.Now().UTC()
	filters := []datastore.Filter{
		datastore.Equals(datastore.Field("userId"), userId),
		datastore.IsInList(datastore.Field("id"), ids...),
	}

	apiKeyIs, err := s.apiKeyStore.Query(filters...).Find()
	if err != nil {
		return err
	}

	err = s.apiKeyStore.Query(filters...).UpdateFields(map[string]interface{}{
		"revokedAt": now,
		"updatedAt": now,
	})
	if err != nil {
		return err
	}

	for _, apiKeyI := range apiKeyIs {
		apiKey, ok := apiKeyI.(*user.ApiKey)
		if !ok {
			return errors.Errorf("type mismatch: expected *user.ApiKey, got %T", apiKeyI)
		}

		err := s.tokenStore.Query(
			datastore.Equals(datastore.Field("appId"), apiKey.AppId),
			datastore.Equals(datastore.Field("userId"), apiKey.UserId),
			datastore.Equals(datastore.Field("device"), apiKeyDevice(apiKey.Id)),
		).Delete()
		if err != nil {
			return err
		}
	}

	return nil
}
