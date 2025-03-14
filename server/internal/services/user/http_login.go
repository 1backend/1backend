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

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	user "github.com/openorch/openorch/server/internal/services/user/types"
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

	req := user.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	token, err := s.login(req.Slug, req.Password)
	if err != nil {
		switch err.Error() {
		case "unauthorized":
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`Invalid password`))
		case "slug not found":
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`Slug not found`))
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
	slug, password string,
) (*user.AuthToken, error) {
	userI, found, err := s.usersStore.Query(
		datastore.Equals(datastore.Field("slug"), slug),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("slug not found")
	}
	u := userI.(*user.User)

	if !checkPasswordHash(password, u.PasswordHash) {
		return nil, errors.New("unauthorized")
	}

	// Let's see if there is an active token we can reuse
	tokenI, found, err := s.authTokensStore.Query(
		datastore.Equals(datastore.Field("userId"), u.Id),
		datastore.Equals(datastore.Field("active"), true),
	).FindOne()
	if err != nil {
		return nil, err
	}

	if found {
		return tokenI.(*user.AuthToken), nil
	}

	token, err := s.generateAuthToken(u)
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

func (s *UserService) generateAuthToken(
	u *user.User,
) (*user.AuthToken, error) {
	roleLinks, err := s.userRoleLinksStore.Query(
		datastore.Equals(datastore.Field("userId"), u.Id),
	).Find()
	if err != nil {
		return nil, err
	}
	roleIds := []string{}
	for _, roleLink := range roleLinks {
		roleIds = append(roleIds, roleLink.(*user.UserRoleLink).RoleId)
	}

	token, err := generateJWT(u, roleIds, s.privateKey)
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
