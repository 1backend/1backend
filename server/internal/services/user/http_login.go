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
	"net/http"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
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
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Router /user-svc/login [post]
func (s *UserService) Login(w http.ResponseWriter, r *http.Request) {
	request := user.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	token, err := s.login(&request)
	if err != nil {
		switch err.Error() {
		case "unauthorized":
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`Invalid password`))
		case "not found":
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`Not found`))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	bs, _ := json.Marshal(user.LoginResponse{
		Token: token,
	})
	w.Write(bs)
}

func (s *UserService) login(
	request *user.LoginRequest,
) (*user.AuthToken, error) {

	var usr *user.User

	if request.Slug != "" {
		userI, found, err := s.usersStore.Query(
			datastore.Equals(datastore.Field("slug"), request.Slug),
		).FindOne()
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, errors.New("not found")
		}

		usr = userI.(*user.User)
	} else if request.Contact != "" {
		contactI, found, err := s.contactsStore.Query(
			datastore.Equals(datastore.Field("id"), request.Contact),
		).FindOne()
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, errors.New("not found")
		}

		userI, found, err := s.usersStore.Query(
			datastore.Equals(datastore.Field("id"), contactI.(*user.Contact).UserId),
		).FindOne()
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, errors.New("not found")
		}

		usr = userI.(*user.User)
	} else {
		return nil, errors.New("slug or contact required")
	}

	passwordIs, err := s.passwordsStore.Query(
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

	// Let's see if there is an active token we can reuse
	tokenI, found, err := s.authTokensStore.Query(
		datastore.Equals(datastore.Field("userId"), usr.Id),
		datastore.Equals(datastore.Field("active"), true),
	).FindOne()
	if err != nil {
		return nil, err
	}

	if found {
		tok := tokenI.(*user.AuthToken)
		functional, err := s.isFunctional(tok.Token)
		if err != nil {
			return nil, errors.Wrap(err, "error checking token functionality")
		}

		if functional {
			return tok, nil
		}

		err = s.inactivateToken(tok.Id)
		if err != nil {
			return nil, errors.Wrap(err, "could not inactivate token")
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

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Changing the struct tags on the `auth.Claims` struct or other
// unexpected shenanigans might cause tokens to become invalid.
func (s *UserService) isFunctional(token string) (bool, error) {
	authr := auth.AuthorizerImpl{}
	claims, err := authr.ParseJWT(s.publicKeyPem, token)
	if err != nil {
		return false, err
	}

	if claims.UserId == "" || len(claims.Roles) == 0 || claims.Slug == "" {
		return false, nil
	}

	return true, nil
}

func (s *UserService) generateAuthToken(
	u *user.User,
) (*user.AuthToken, error) {
	roles, err := s.getRolesByUserId(u.Id)
	if err != nil {
		return nil, errors.Wrap(err, "error listing roles")
	}
	if len(roles) == 0 {
		return nil, errors.New("no roles found for user")
	}
	_, activeOrganizationId, err := s.getUserOrganizations(u.Id)
	if err != nil {
		return nil, errors.Wrap(err, "error listing organizations")
	}

	token, err := generateJWT(u, roles, activeOrganizationId, s.privateKey)
	if err != nil {
		return nil, err
	}

	return &user.AuthToken{
		Id:        sdk.Id("tok"),
		UserId:    u.Id,
		Token:     token,
		Active:    true,
		CreatedAt: time.Now(),
	}, nil
}
