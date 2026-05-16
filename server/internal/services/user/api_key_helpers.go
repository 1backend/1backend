/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"log/slog"
	"net/http"
	"strings"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

const (
	apiKeyBearerPrefix = "obk_"
	apiKeyDevicePrefix = "api-key:"
)

func (s *UserService) apiKeyFromBearer(rawKey string) (*user.ApiKey, error) {
	id, secret, err := parseAPIKey(rawKey)
	if err != nil {
		return nil, errors.New("unauthorized: invalid api key")
	}

	apiKeyI, found, err := s.apiKeyStore.Query(
		datastore.Id(id),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("unauthorized: invalid api key")
	}

	apiKey, ok := apiKeyI.(*user.ApiKey)
	if !ok {
		return nil, errors.Errorf("type mismatch: expected *user.ApiKey, got %T", apiKeyI)
	}
	if apiKey.SecretHash == "" {
		return nil, errors.New("unauthorized: invalid api key")
	}

	providedHash := apiKeySecretHash(secret)
	if subtle.ConstantTimeCompare([]byte(apiKey.SecretHash), []byte(providedHash)) != 1 {
		return nil, errors.New("unauthorized: invalid api key")
	}

	return apiKey, nil
}

func (s *UserService) apiKeyRequestAppId(
	defaultAppId string,
	appId string,
	appHost string,
) (string, error) {
	if appId != "" && appHost != "" {
		return "", errors.New("appHost and appId are mutually exclusive")
	}

	if appId != "" {
		_, found, err := s.appByID(appId)
		if err != nil {
			return "", err
		}
		if !found {
			return "", errors.Errorf("app not found by id '%s'", appId)
		}
		return appId, nil
	}

	if appHost != "" {
		app, err := s.getOrCreateApp(appHost)
		if err != nil {
			return "", err
		}
		return app.Id, nil
	}

	return defaultAppId, nil
}

func (s *UserService) userById(userId string) (*user.User, bool, error) {
	if userId == "" {
		return nil, false, nil
	}

	userI, found, err := s.userStore.Query(
		datastore.Id(userId),
	).FindOne()
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}

	usr, ok := userI.(*user.User)
	if !ok {
		return nil, false, errors.Errorf("type mismatch: expected *user.User, got %T", userI)
	}

	return usr, true, nil
}

func (s *UserService) validateApiKeyActiveOrganization(
	appId string,
	userId string,
	activeOrganizationId string,
) error {
	if activeOrganizationId == "" {
		return nil
	}

	membership, found, err := s.findMembership(appId, userId, activeOrganizationId)
	if err != nil {
		return err
	}
	if !found || membership.Status != user.MembershipStatusAccepted {
		return errors.New("unauthorized: active organization not available")
	}

	return nil
}

func apiKeyView(apiKey *user.ApiKey) user.ApiKeyView {
	return user.ApiKeyView{
		Id:                   apiKey.Id,
		AppId:                apiKey.AppId,
		UserId:               apiKey.UserId,
		Name:                 apiKey.Name,
		Prefix:               apiKey.Prefix,
		ActiveOrganizationId: apiKey.ActiveOrganizationId,
		CreatedAt:            apiKey.CreatedAt,
		UpdatedAt:            apiKey.UpdatedAt,
		LastUsedAt:           apiKey.LastUsedAt,
		ExpiresAt:            apiKey.ExpiresAt,
		RevokedAt:            apiKey.RevokedAt,
	}
}

func apiKeyDevice(apiKeyId string) string {
	return apiKeyDevicePrefix + apiKeyId
}

func canManageAPIKeysOnBehalf(claims *auth.Claims) bool {
	if claims == nil {
		return false
	}

	for _, role := range claims.Roles {
		if role == user.RoleAdmin {
			return true
		}
	}

	return false
}

func formatAPIKey(id string, secret string) string {
	return apiKeyBearerPrefix + id + "_" + secret
}

func parseAPIKey(rawKey string) (string, string, error) {
	trimmed := strings.TrimSpace(rawKey)
	trimmed = strings.TrimPrefix(trimmed, "Bearer ")
	if !strings.HasPrefix(trimmed, apiKeyBearerPrefix) {
		return "", "", errors.New("missing api key prefix")
	}

	parts := strings.SplitN(strings.TrimPrefix(trimmed, apiKeyBearerPrefix), "_", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", errors.New("invalid api key")
	}

	return parts[0], parts[1], nil
}

func newAPIKeySecret() (string, error) {
	var secret [32]byte
	if _, err := rand.Read(secret[:]); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(secret[:]), nil
}

func apiKeySecretHash(secret string) string {
	sum := sha256.Sum256([]byte(secret))
	return hex.EncodeToString(sum[:])
}

func apiKeyDisplayPrefix(rawKey string) string {
	if len(rawKey) <= 24 {
		return rawKey
	}
	return rawKey[:24]
}

func isBadApiKeyRequest(err error) bool {
	if err == nil {
		return false
	}

	msg := err.Error()
	return strings.Contains(msg, "appHost and appId are mutually exclusive") ||
		strings.Contains(msg, "app not found") ||
		strings.Contains(msg, "expiresAt must be in the future") ||
		strings.Contains(msg, "id or ids are required") ||
		strings.Contains(msg, "request is required") ||
		strings.Contains(msg, "user not found")
}

func writeAuthOrInternalError(w http.ResponseWriter, err error, message string) {
	if isUnauthorizedRequestError(err) ||
		err != nil && strings.HasPrefix(err.Error(), "unauthorized:") {
		endpoint.Unauthorized(w)
		return
	}

	logger.Error(
		message,
		slog.Any("error", err),
	)
	endpoint.InternalServerError(w)
}
