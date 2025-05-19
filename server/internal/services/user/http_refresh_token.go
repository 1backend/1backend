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
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log/slog"
	"net/http"
	"sort"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID refreshToken
// @Summary Refresh Token
// @Description Refreshes an existing token, including inactive ones.
// @Description The old token becomes inactive (if not already inactive), and a new, active token is issued.
// @Description This allows continued verification of user roles without requiring a new login.
// @Description Inactive tokens are refreshable unless explicitly revoked (no mechanism for this yet).
// @Description Leaked tokens should be handled separately, via a revocation flag or deletion.
// @Tags User Svc
// @Accept json
// @Produce json
// @Success 200 {object} user.RefreshTokenResponse "Refresh Token successful"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Router /user-svc/refresh-token [post]
func (s *UserService) RefreshToken(w http.ResponseWriter, r *http.Request) {
	stringToken, exists := s.authorizer.TokenFromRequest(r)
	if !exists {
		endpoint.Unauthorized(w)
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

	bs, _ := json.Marshal(user.RefreshTokenResponse{
		Token: token,
	})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (s *UserService) refreshToken(
	stringToken string,
) (*user.AuthToken, error) {
	cacheKey := generateCacheKey(stringToken)
	if cachedToken, found := s.tokenReplacementCache.Get(cacheKey); found {
		return cachedToken.(*user.AuthToken), nil
	}

	tokenToBeRefreshedI, found, err := s.authTokensStore.Query(
		datastore.Equals(datastore.Field("token"), stringToken),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("token not found")
	}

	tokenToBeRefreshed := tokenToBeRefreshedI.(*user.AuthToken)

	userI, found, err := s.usersStore.Query(
		datastore.Equals(datastore.Field("id"), tokenToBeRefreshed.UserId),
	).FindOne()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query user")
	}
	if !found {
		return nil, errors.New("user not found")
	}
	usr := userI.(*user.User)

	activeTokenI, found, err := s.authTokensStore.Query(
		datastore.Equals(datastore.Field("userId"), usr.Id),
		datastore.Equals(datastore.Field("active"), true),
		datastore.Equals(datastore.Field("device"), tokenToBeRefreshed.Device),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("active token not found")
	}

	activeToken := activeTokenI.(*user.AuthToken)

	if tokenToBeRefreshed.Active {
		err = s.inactivateToken(tokenToBeRefreshed.Id)
		if err != nil {
			return nil, errors.Wrap(err, "failed to inactivate token")
		}
	}

	now := time.Now()

	err = s.authTokensStore.Query(
		datastore.Equals(datastore.Field("id"), tokenToBeRefreshed.Id),
	).UpdateFields(map[string]any{
		"lastRefreshedAt": now,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to update token")
	}

	if activeToken.Id != tokenToBeRefreshed.Id {
		err = s.inactivateToken(activeToken.Id)
		if err != nil {
			return nil, errors.Wrap(err, "failed to inactivate token")
		}
	}

	token, err := s.generateAuthToken(usr)
	if err != nil {
		return nil, err
	}
	token.Device = tokenToBeRefreshed.Device

	// for backwards compatibility
	if token.Device == "" {
		token.Device = defaultDevice
	}

	err = s.authTokensStore.Create(token)
	if err != nil {
		return nil, errors.Wrap(err, "error creating token")
	}

	// Prune old tokens for the same device
	tokens, err := s.authTokensStore.Query(
		datastore.Equals(datastore.Field("userId"), usr.Id),
		datastore.Equals(datastore.Field("device"), tokenToBeRefreshed.Device),
	).Find()
	if err != nil {
		logger.Error("Failed to query tokens for pruning", slog.Any("error", err))
		return token, nil // return token anyway, pruning is best-effort
	}

	// Sort tokens by LastRefreshedAt descending, treating nil as zero (oldest)
	sort.Slice(tokens, func(i, j int) bool {
		ti := tokens[i].(*user.AuthToken)
		tj := tokens[j].(*user.AuthToken)

		var tiTime, tjTime time.Time
		if ti.LastRefreshedAt != nil {
			tiTime = *ti.LastRefreshedAt
		}
		if tj.LastRefreshedAt != nil {
			tjTime = *tj.LastRefreshedAt
		}
		// Now compare: newer first
		return tiTime.After(tjTime)
	})

	// Keep latest 2, delete the rest
	for i := 2; i < len(tokens); i++ {
		t := tokens[i].(*user.AuthToken)
		if t.Active {
			// We don't want to delete the active token
			continue
		}
		err := s.authTokensStore.Query(datastore.Id(t.Id)).Delete()
		if err != nil {
			logger.Error("Failed to delete old token",
				slog.String("tokenId", t.Id),
				slog.Any("error", err),
			)
		}
	}

	s.tokenReplacementCache.SetWithTTL(
		cacheKey,
		token,
		1,
		s.tokenExpiration,
	)

	return token, nil
}

func generateCacheKey(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
