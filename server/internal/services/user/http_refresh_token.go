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
			"Failed to login",
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

	err = s.authTokensStore.Create(token)
	if err != nil {
		return nil, errors.Wrap(err, "error creating token")
	}

	return token, nil
}
