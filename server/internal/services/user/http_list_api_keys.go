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

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID listApiKeys
// @Summary List API Keys
// @Description Lists API key metadata without returning raw secrets.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ListApiKeysRequest false "List API Keys Request"
// @Success 200 {object} user.ListApiKeysResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/api-keys [post]
func (s *UserService) ListApiKeys(w http.ResponseWriter, r *http.Request) {
	_, claims, err := s.getUserFromRequest(r)
	if err != nil {
		writeAuthOrInternalError(w, err, "Failed to get user from request")
		return
	}

	request := user.ListApiKeysRequest{}
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

	apiKeys, err := s.listApiKeys(claims, &request)
	if err != nil {
		if isBadApiKeyRequest(err) {
			endpoint.WriteString(w, http.StatusBadRequest, err.Error())
			return
		}

		logger.Error(
			"Failed to list API keys",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.ListApiKeysResponse{
		ApiKeys: apiKeys,
	})
}

func (s *UserService) listApiKeys(
	claims *auth.Claims,
	request *user.ListApiKeysRequest,
) ([]user.ApiKeyView, error) {
	if claims == nil {
		return nil, errors.New("claims are required")
	}
	if request == nil {
		return nil, errors.New("request is required")
	}

	appId, err := s.apiKeyRequestAppId(claims.AppId, request.AppId, request.AppHost)
	if err != nil {
		return nil, err
	}

	userId := request.UserId
	if userId == "" {
		userId = claims.UserId
	}

	filters := []datastore.Filter{
		datastore.Equals(datastore.Field("appId"), appId),
		datastore.Equals(datastore.Field("userId"), userId),
	}

	if !request.IncludeRevoked {
		filters = append(filters, datastore.Equals(datastore.Field("revokedAt"), nil))
	}

	ids := make([]any, 0, len(request.Ids)+1)
	if request.Id != "" {
		ids = append(ids, request.Id)
	}
	for _, id := range request.Ids {
		ids = append(ids, id)
	}
	if len(ids) > 0 {
		filters = append(filters, datastore.IsInList(datastore.Field("id"), ids...))
	}

	apiKeyIs, err := s.apiKeyStore.Query(filters...).
		OrderBy(datastore.OrderByField("createdAt", true)).
		Find()
	if err != nil {
		return nil, err
	}

	apiKeys := make([]user.ApiKeyView, 0, len(apiKeyIs))
	for _, apiKeyI := range apiKeyIs {
		apiKey, ok := apiKeyI.(*user.ApiKey)
		if !ok {
			return nil, errors.Errorf("type mismatch: expected *user.ApiKey, got %T", apiKeyI)
		}
		apiKeys = append(apiKeys, apiKeyView(apiKey))
	}

	return apiKeys, nil
}
