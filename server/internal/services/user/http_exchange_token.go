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

// @ID exchangeToken
// @Summary Exchange Token
// @Description Exchange an existing token for a new token scoped to a different app (namespace).
// @Description The new token represents the same user but contains roles specific to the target app.
// @Description
// @Description The original token remains valid.
// @Description The minted token is not stored and cannot be refreshed (and will have the same expiration duration as normal tokens), unlike tokens acquired via login.
// @Description
// @Description For now, token exchange is designed to be in situ — the User Svc must be contacted at exchange time.
// @Description This introduces a stateful dependency on the User Svc, but simplifies things until broader use cases emerge.
// @Tags User Svc
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ExchangeTokenRequest true "ExchangeToken Request"
// @Success 200 {object} user.ExchangeTokenResponse "ExchangeToken successful"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 404 {object} user.ErrorResponse "User Not Found"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Router /user-svc/token/exchange [put]
func (s *UserService) ExchangeToken(w http.ResponseWriter, r *http.Request) {
	request := user.ExchangeTokenRequest{}
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

	if request.NewAppHost == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "New app is required")
	}

	claims, err := s.options.Authorizer.ParseJWTFromRequest(s.publicKeyPem, r)
	if err != nil {
		logger.Error(
			"Failed to parse JWT from request",
			slog.Any("error", err),
			slog.Any("request", request),
		)
		endpoint.InternalServerError(w)
	}

	token, err := s.exchangeToken(claims, &request)
	if err != nil {
		logger.Error(
			"Failed to exchange token",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(user.ExchangeTokenResponse{
		Token: token,
	})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (s *UserService) exchangeToken(
	claims *auth.Claims,
	request *user.ExchangeTokenRequest,
) (*user.Token, error) {
	userI, found, err := s.userStore.Query(
		datastore.Id(claims.UserId),
	).FindOne()
	if err != nil {
		return nil, errors.Wrap(err, "error querying user")
	}
	if !found {
		return nil, errors.New("user not found")
	}
	usr, ok := userI.(*user.User)
	if !ok {
		return nil, errors.Errorf("type mismatch: expected *user.User, got %T", userI)
	}

	if request.NewDevice == "" {
		request.NewDevice = claims.Device
	}

	app, err := s.getOrCreateApp(request.NewAppHost)
	if err != nil {
		return nil, errors.Wrap(err, "error getting or creating app")
	}

	return s.issueToken(app.GetId(), usr, request.NewDevice)
}
