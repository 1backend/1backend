/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"sort"
	"time"

	"github.com/1backend/1backend/sdk/go/auth"
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
	stringToken, exists := s.options.Authorizer.TokenFromRequest(r)
	if !exists {
		endpoint.Unauthorized(w)
		return
	}

	token, err := s.refreshToken(r.Context(), stringToken)
	if err != nil {
		args := append(
			attrsToArgs(auth.TokenDebugAttrs(stringToken)),
			slog.Any("error", err),
		)
		logger.Error("Failed to refresh token", args...)
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
	ctx context.Context,
	stringToken string,
) (*user.Token, error) {
	cacheKey := generateCacheKey(stringToken)

	// Fast Path: Check cache without any locking (handles 99% of traffic)
	if cachedToken, found := s.cachedReplacementToken(cacheKey); found {
		return cachedToken, nil
	}

	val, err, _ := s.refreshGroup.Do(cacheKey, func() (interface{}, error) {
		// THE DOUBLE CHECK:
		// Requests 2, 3, 4, and 5 wait at the door. When Request 1 finishes
		// and sets the cache, these requests enter one by one (or share the result).
		// We check the cache again to see if Request 1 already did the work.
		if cachedToken, found := s.cachedReplacementToken(cacheKey); found {
			return cachedToken, nil
		}

		releaseEntryLock, err := s.acquireRefreshTokenLock(
			ctx,
			"user-svc:refresh-token-entry",
		)
		if err != nil {
			return nil, err
		}
		entryLockHeld := true
		defer func() {
			if entryLockHeld {
				releaseEntryLock()
			}
		}()

		lockKey, err := s.refreshTokenLockKey(stringToken)
		if err != nil {
			return nil, err
		}

		// Mark the exact presented token before waiting on the per-device lock.
		// A concurrent refresh holding that lock can prune old tokens; this keeps
		// an already in-flight refresh from looking abandoned.
		now := time.Now()
		err = s.tokenStore.Query(
			datastore.Equals(datastore.Field("token"), stringToken),
		).UpdateFields(map[string]any{
			"lastRefreshedAt": now,
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to update token")
		}

		releaseEntryLock()
		entryLockHeld = false

		releaseLock, err := s.acquireRefreshTokenLock(ctx, lockKey)
		if err != nil {
			return nil, err
		}
		defer releaseLock()

		if cachedToken, found := s.cachedReplacementToken(cacheKey); found {
			return cachedToken, nil
		}

		tokenToBeRefreshedI, found, err := s.tokenStore.Query(
			datastore.Equals(datastore.Field("token"), stringToken),
		).FindOne()
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, errors.New("token not found")
		}

		tokenToBeRefreshed := tokenToBeRefreshedI.(*user.Token)

		userI, found, err := s.userStore.Query(
			datastore.Equals(datastore.Field("id"), tokenToBeRefreshed.UserId),
		).FindOne()
		if err != nil {
			return nil, errors.Wrap(err, "failed to query user")
		}
		if !found {
			return nil, errors.New("user not found")
		}
		usr := userI.(*user.User)

		activeTokenI, found, err := s.tokenStore.Query(
			datastore.Equals(datastore.Field("appId"), tokenToBeRefreshed.AppId),
			datastore.Equals(datastore.Field("userId"), usr.Id),
			datastore.Equals(datastore.Field("active"), true),
			datastore.Equals(datastore.Field("device"), tokenToBeRefreshed.Device),
		).FindOne()
		if err != nil {
			return nil, err
		}
		var activeToken *user.Token
		if found {
			activeToken = activeTokenI.(*user.Token)
		}

		if tokenToBeRefreshed.Active {
			err = s.inactivateToken(tokenToBeRefreshed.AppId, tokenToBeRefreshed.Id)
			if err != nil {
				return nil, errors.Wrap(err, "failed to inactivate token")
			}
		}

		now = time.Now()

		err = s.tokenStore.Query(
			datastore.Equals(datastore.Field("id"), tokenToBeRefreshed.Id),
		).UpdateFields(map[string]any{
			"lastRefreshedAt": now,
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to update token")
		}

		if activeToken != nil && activeToken.Id != tokenToBeRefreshed.Id {
			err = s.inactivateToken(tokenToBeRefreshed.AppId, activeToken.Id)
			if err != nil {
				return nil, errors.Wrap(err, "failed to inactivate token")
			}
		}

		token, err := s.generateAuthToken(
			tokenToBeRefreshed.AppId,
			usr,
			tokenToBeRefreshed.Device,
		)
		if err != nil {
			return nil, err
		}

		if token.Device == "" {
			token.Device = unknownDevice
		}

		err = s.tokenStore.Create(token)
		if err != nil {
			return nil, errors.Wrap(err, "error creating token")
		}

		// Prune old tokens for the same device
		tokens, err := s.tokenStore.Query(
			datastore.Equals(datastore.Field("appId"), token.AppId),
			datastore.Equals(datastore.Field("userId"), usr.Id),
			datastore.Equals(datastore.Field("device"), tokenToBeRefreshed.Device),
		).Find()
		if err != nil {
			logger.Error("Failed to query tokens for pruning", slog.Any("error", err))
			return token, nil // return token anyway, pruning is best-effort
		}

		// Sort tokens by recent use. LastRefreshedAt is the strongest signal,
		// while CreatedAt keeps newly issued active tokens ahead of old nils.
		sort.Slice(tokens, func(i, j int) bool {
			ti := tokens[i].(*user.Token)
			tj := tokens[j].(*user.Token)

			return tokenPruneTime(ti).After(tokenPruneTime(tj))
		})

		// Keep the new active token and the two most recently used inactive tokens.
		keptInactive := 0
		for _, tokenI := range tokens {
			t := tokenI.(*user.Token)
			if t.Id == token.Id {
				continue
			}
			if !t.Active && keptInactive < 2 {
				keptInactive++
				continue
			}
			err := s.tokenStore.Query(datastore.Id(t.Id)).Delete()
			if err != nil {
				logger.Error("Failed to delete old token",
					slog.String("appId", t.AppId),
					slog.String("tokenId", t.Id),
					slog.Any("error", err),
				)
			}
		}

		s.tokenReplacementCache.SetWithTTL(
			cacheKey,
			token,
			1,
			s.options.TokenExpiration,
		)
		if s.options.Test {
			s.tokenReplacementCache.Wait()
		}

		return token, nil
	})

	if err != nil {
		return nil, err
	}

	return val.(*user.Token), nil
}

func (s *UserService) cachedReplacementToken(
	cacheKey string,
) (*user.Token, bool) {
	cachedToken, found := s.tokenReplacementCache.Get(cacheKey)
	if !found {
		return nil, false
	}

	token, ok := cachedToken.(*user.Token)
	if !ok {
		s.tokenReplacementCache.Del(cacheKey)
		if s.options.Test {
			s.tokenReplacementCache.Wait()
		}
		return nil, false
	}

	_, found, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("id"), token.Id),
		datastore.Equals(datastore.Field("token"), token.Token),
		datastore.Equals(datastore.Field("active"), true),
	).FindOne()
	if err != nil {
		logger.Error(
			"Failed to validate cached replacement token",
			slog.String("tokenId", token.Id),
			slog.Any("error", err),
		)
		return nil, false
	}
	if !found {
		s.tokenReplacementCache.Del(cacheKey)
		if s.options.Test {
			s.tokenReplacementCache.Wait()
		}
		return nil, false
	}

	return token, true
}

func (s *UserService) acquireRefreshTokenLock(
	ctx context.Context,
	key string,
) (func(), error) {
	if s.options.Lock == nil {
		return func() {}, nil
	}

	err := s.options.Lock.Acquire(ctx, key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to acquire refresh token lock")
	}

	return func() {
		if err := s.options.Lock.Release(context.Background(), key); err != nil {
			logger.Error(
				"Failed to release refresh token lock",
				slog.String("lockKey", key),
				slog.Any("error", err),
			)
		}
	}, nil
}

func (s *UserService) refreshTokenLockKey(
	stringToken string,
) (string, error) {
	claims, err := s.options.Authorizer.ParseJWTUnverified(stringToken)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse refresh token claims")
	}
	if claims.AppId == "" || claims.UserId == "" {
		return "", errors.New("refresh token claims missing appId or userId")
	}
	device := claims.Device
	if device == "" {
		device = unknownDevice
	}

	return "user-svc:refresh-token:" + claims.AppId + ":" + claims.UserId + ":" + device, nil
}

func refreshTokenLockKeyForParts(appId, userId, device string) string {
	if device == "" {
		device = unknownDevice
	}

	return "user-svc:refresh-token:" + appId + ":" + userId + ":" + device
}

func tokenPruneTime(token *user.Token) time.Time {
	if token.LastRefreshedAt != nil {
		return *token.LastRefreshedAt
	}
	return token.CreatedAt
}

func attrsToArgs(attrs []slog.Attr) []any {
	if len(attrs) == 0 {
		return nil
	}

	args := make([]any, 0, len(attrs))
	for _, attr := range attrs {
		args = append(args, attr)
	}

	return args
}

func generateCacheKey(token string) string {
	return auth.TokenHash(token)
}
