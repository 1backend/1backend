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

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID createApiKey
// @Summary Create API Key
// @Description Creates an opaque bearer API key for the current user.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.CreateApiKeyRequest false "Create API Key Request"
// @Success 200 {object} user.CreateApiKeyResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/api-key [post]
func (s *UserService) CreateApiKey(w http.ResponseWriter, r *http.Request) {
	_, claims, err := s.getUserFromRequest(r)
	if err != nil {
		writeAuthOrInternalError(w, err, "Failed to get user from request")
		return
	}

	request := user.CreateApiKeyRequest{}
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

	response, err := s.createApiKey(claims, &request)
	if err != nil {
		if isUnauthorizedRequestError(err) {
			endpoint.Unauthorized(w)
			return
		}
		if isBadApiKeyRequest(err) {
			endpoint.WriteString(w, http.StatusBadRequest, err.Error())
			return
		}

		logger.Error(
			"Failed to create API key",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, response)
}

func (s *UserService) createApiKey(
	claims *auth.Claims,
	request *user.CreateApiKeyRequest,
) (*user.CreateApiKeyResponse, error) {
	if claims == nil {
		return nil, errors.New("claims are required")
	}
	if request == nil {
		return nil, errors.New("request is required")
	}

	now := time.Now().UTC()
	if request.ExpiresAt != nil && !request.ExpiresAt.After(now) {
		return nil, errors.New("expiresAt must be in the future")
	}

	appId, err := s.apiKeyRequestAppId(claims.AppId, request.AppId, request.AppHost)
	if err != nil {
		return nil, err
	}

	userId := request.UserId
	if userId == "" {
		userId = claims.UserId
	}
	_, found, err := s.userById(userId)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.Errorf("user not found by id '%s'", userId)
	}

	activeOrganizationId := ""
	if request.ActiveOrganizationId != nil {
		activeOrganizationId = *request.ActiveOrganizationId
	} else if appId == claims.AppId && userId == claims.UserId {
		activeOrganizationId = claims.ActiveOrganizationId
	}
	if err := s.validateApiKeyActiveOrganization(appId, userId, activeOrganizationId); err != nil {
		return nil, err
	}

	id := sdk.Id("apik")
	secret, err := newAPIKeySecret()
	if err != nil {
		return nil, err
	}
	key := formatAPIKey(id, secret)

	internalId, err := sdk.InternalId(appId, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create api key internal id")
	}

	apiKey := &user.ApiKey{
		InternalId:           internalId,
		Id:                   id,
		AppId:                appId,
		UserId:               userId,
		Name:                 request.Name,
		Prefix:               apiKeyDisplayPrefix(key),
		SecretHash:           apiKeySecretHash(secret),
		ActiveOrganizationId: activeOrganizationId,
		CreatedAt:            now,
		UpdatedAt:            now,
		ExpiresAt:            request.ExpiresAt,
	}

	if err := s.apiKeyStore.Create(apiKey); err != nil {
		return nil, err
	}

	return &user.CreateApiKeyResponse{
		ApiKey: apiKeyView(apiKey),
		Key:    key,
	}, nil
}
