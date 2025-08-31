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
	"strings"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// @ID login
// @Summary Login
// @Description Authenticates a user and returns a token.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.LoginRequest true "Login Request"
// @Success 200 {object} user.LoginResponse "Login successful"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 404 {object} user.ErrorResponse "User Not Found"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Router /user-svc/login [post]
func (s *UserService) Login(w http.ResponseWriter, r *http.Request) {
	request := user.LoginRequest{}
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

	token, err := s.login(&request)
	if err != nil {
		switch err.Error() {
		case "unauthorized":
			logger.Error(
				"Invalid password",
				slog.Any("slug", request.Slug),
				slog.Any("contact", request.Contact),
			)
			endpoint.WriteString(w, http.StatusUnauthorized, "Invalid Password")
		case "not found":
			logger.Error(
				"Cannot find user",
				slog.Any("slug", request.Slug),
				slog.Any("contact", request.Contact),
				slog.Any("error", err),
			)
			endpoint.WriteString(w, http.StatusNotFound, "Not Found")
		default:
			logger.Error(
				"Failed to login",
				slog.Any("error", err),
			)
			endpoint.InternalServerError(w)
		}
		return
	}

	bs, _ := json.Marshal(user.LoginResponse{
		Token: token,
	})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (s *UserService) login(
	request *user.LoginRequest,
) (*user.Token, error) {
	if request.App == "" {
		request.App = "unnamed"
	}

	var usr *user.User

	if request.Slug != "" {
		userI, found, err := s.userStore.Query(
			datastore.Equals(datastore.Field("slug"), request.Slug),
		).FindOne()
		if err != nil {
			return nil, errors.Wrap(err, "error querying user by slug")
		}
		if !found {
			return nil, errors.New("user not found by slug")
		}

		usr = userI.(*user.User)
	} else if request.Contact != "" {
		contactI, found, err := s.contactStore.Query(
			datastore.Equals(datastore.Field("id"), request.Contact),
		).FindOne()
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, errors.New("not found")
		}

		userI, found, err := s.userStore.Query(
			datastore.Equals(datastore.Field("id"), contactI.(*user.Contact).UserId),
		).FindOne()
		if err != nil {
			return nil, errors.Wrap(err, "error querying user by contact")
		}
		if !found {
			return nil, errors.New("user not found by contact")
		}

		usr = userI.(*user.User)
	} else {
		return nil, errors.New("slug or contact required")
	}

	passwordIs, err := s.passwordStore.Query(
		datastore.Equals(
			datastore.Field("userId"), usr.Id),
	).OrderBy(
		datastore.OrderByField("createdAt", true),
	).Limit(1).Find()
	if err != nil {
		return nil, errors.Wrap(err, "error listing passwords for user")
	}
	if len(passwordIs) == 0 {
		return nil, errors.New("password not found for user")
	}

	password := passwordIs[0].(*user.Password)

	if !checkPasswordHash(request.Password, password.PasswordHash) {
		return nil, errors.New("unauthorized")
	}

	if request.Device == "" {
		request.Device = unknownDevice
	}

	return s.issueToken(request.App, usr, request.Device)
}

func (s *UserService) issueToken(
	app string,
	usr *user.User,
	device string,
) (*user.Token, error) {

	// Let's see if there is an active token we can reuse
	tokenI, found, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("app"), app),
		datastore.Equals(datastore.Field("userId"), usr.Id),
		datastore.Equals(datastore.Field("active"), true),
		datastore.Equals(datastore.Field("device"), device),
	).FindOne()
	if err != nil {
		return nil, err
	}

	if found {
		tok := tokenI.(*user.Token)
		functional, err := s.isFunctional(tok.Token)
		if err != nil {
			if strings.Contains(err.Error(), "token is expired") {
				// @todo this feels pretty crufty
				tok, err := s.refreshToken(tok.Token)
				if err != nil {
					return nil, errors.Wrap(err, "error refreshing token")
				}

				return tok, nil
			}
			return nil, errors.Wrap(err, "error checking token functionality")
		}

		if functional {
			return tok, nil
		}

		err = s.inactivateToken(app, tok.Id)
		if err != nil {
			return nil, errors.Wrap(err, "could not inactivate token")
		}
	}

	token, err := s.generateAuthToken(
		app,
		usr,
		device,
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

	return token, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Changing the struct tags on the `auth.Claims` struct or other
// unexpected shenanigans might cause tokens to become invalid.
func (s *UserService) isFunctional(token string) (bool, error) {
	claims, err := s.options.Authorizer.ParseJWT(s.publicKeyPem, token)
	if err != nil {
		return false, err
	}

	if claims.UserId == "" || len(claims.Roles) == 0 || claims.Slug == "" {
		return false, nil
	}

	return true, nil
}

func (s *UserService) generateAuthToken(
	app string,
	u *user.User,
	device string,
) (*user.Token, error) {
	roles, err := s.getRolesByUserId(app, u.Id)
	if err != nil {
		return nil, errors.Wrap(err, "error listing roles")
	}

	_, activeOrganizationId, err := s.getUserOrganizations(app, u.Id)
	if err != nil {
		return nil, errors.Wrap(err, "error listing organizations")
	}

	claims, token, err := s.generateJWT(
		app, u, roles, activeOrganizationId,
		s.privateKey, device)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	id := sdk.Id("tok")

	return &user.Token{
		InternalId: sdk.InternalId(claims.ID, app),
		Id:         id,
		App:        app,
		UserId:     u.Id,
		Token:      token,
		Device:     device,
		Active:     true,
		ExpiresAt:  now.Add(s.options.TokenExpiration),
		CreatedAt:  now,
	}, nil
}
