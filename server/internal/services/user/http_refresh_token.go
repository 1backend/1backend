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
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID refreshToken
// @Summary Refresh Token
// @Description Refreshes a token.
// @Tags User Svc
// @Accept json
// @Produce json
// @Success 200 {object} user.RefreshTokenResponse "Refresh Token successful"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Router /user-svc/refresh-token [post]
func (s *UserService) RefreshToken(w http.ResponseWriter, r *http.Request) {
	token, exists := s.authorizer.TokenFromRequest(r)
	if !exists {
		endpoint.Unauthorized(w)
		return
	}

	_, err := s.refreshToken(token)
	if err != nil {
		switch err.Error() {
		case "unauthorized":
			endpoint.WriteString(w, http.StatusUnauthorized, "Invalid Password")
		case "not found":
			endpoint.WriteString(w, http.StatusBadRequest, "Not Found")
		default:
			logger.Error(
				"Failed to login",
				slog.Any("error", err),
			)
			endpoint.InternalServerError(w)
		}
		return
	}

	bs, _ := json.Marshal(user.LoginResponse{})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (s *UserService) refreshToken(
	token string,
) (*user.AuthToken, error) {

	authTokenI, found, err := s.authTokensStore.Query(
		datastore.Equals(datastore.Field("token"), token),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("token not found")
	}

	logger.Info("authTokenI", slog.Any("authTokenI", authTokenI))
	//return token.(*user.AuthToken), nil

	return nil, nil
}
