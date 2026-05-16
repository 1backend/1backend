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
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID exchangeApiKey
// @Summary Exchange API Key
// @Description Exchanges an opaque API key bearer into a normal short-lived JWT.
// @Tags User Svc
// @Accept json
// @Produce json
// @Success 200 {object} user.ExchangeApiKeyResponse
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/api-key/exchange [post]
func (s *UserService) ExchangeApiKey(w http.ResponseWriter, r *http.Request) {
	token, err := s.exchangeApiKeyFromRequest(r)
	if err != nil {
		writeAuthOrInternalError(w, err, "Failed to exchange API key")
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.ExchangeApiKeyResponse{
		Token: token,
	})
}

func (s *UserService) exchangeApiKeyFromRequest(r *http.Request) (*user.Token, error) {
	rawKey, exists := s.options.Authorizer.TokenFromRequest(r)
	if !exists {
		return nil, errors.New("unauthorized: missing api key")
	}

	apiKey, err := s.apiKeyFromBearer(rawKey)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	if apiKey.RevokedAt != nil {
		return nil, errors.New("unauthorized: api key revoked")
	}
	if apiKey.ExpiresAt != nil && !apiKey.ExpiresAt.After(now) {
		return nil, errors.New("unauthorized: api key expired")
	}
	if err := s.validateApiKeyActiveOrganization(
		apiKey.AppId,
		apiKey.UserId,
		apiKey.ActiveOrganizationId,
	); err != nil {
		return nil, err
	}

	usr, found, err := s.userById(apiKey.UserId)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("unauthorized: user not found")
	}

	device := apiKeyDevice(apiKey.Id)
	if err := s.setActivation(
		apiKey.AppId,
		apiKey.UserId,
		device,
		apiKey.ActiveOrganizationId,
	); err != nil {
		return nil, err
	}

	token, err := s.issueToken(
		apiKey.AppId,
		usr,
		device,
	)
	if err != nil {
		return nil, err
	}

	usedAt := time.Now().UTC()
	if err := s.apiKeyStore.Query(datastore.Id(apiKey.Id)).UpdateFields(map[string]interface{}{
		"lastUsedAt": usedAt,
		"updatedAt":  usedAt,
	}); err != nil {
		logger.Error(
			"Failed to update API key last used timestamp",
			slog.Any("error", err),
			slog.String("apiKeyId", apiKey.Id),
		)
	}

	return token, nil
}
