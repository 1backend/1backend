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
	"errors"
	"net/http"
	"time"

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	user "github.com/openorch/openorch/server/internal/services/user/types"
	"golang.org/x/crypto/bcrypt"
)

// @ID register
// @Summary Register
// @Description Register a new user with a name, email, and password.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.RegisterRequest true "Register Request"
// @Success 200 {object} user.RegisterResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Router /user-svc/register [post]
func (s *UserService) Register(w http.ResponseWriter, r *http.Request) {
	req := user.RegisterRequest{}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	newUser := &user.User{
		Name: req.Name,
		Slug: req.Slug,
	}

	if req.Contact.Value != "" {
		now := time.Now()
		req.Contact.CreatedAt = now
		req.Contact.UpdatedAt = now
		newUser.Contacts = []user.Contact{
			req.Contact,
		}
	}

	err = s.createUser(newUser, req.Password, []string{user.RoleUser.Id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	token, err := s.login(req.Slug, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.RegisterResponse{
		Token: token,
	})
	w.Write(bs)
}

func validateUserRegistration(req user.RegisterRequest) error {
	if req.Name == "" {
		return errors.New("name missing")
	}
	if req.Contact.Id == "" {
		return errors.New("email missing")
	}
	if req.Slug == "" {
		req.Slug = req.Contact.Id
	}
	if req.Password == "" {
		return errors.New("password missing")
	}

	return nil
}

func (s *UserService) register(
	slug, password, name string,
	roleIds []string,
) (*user.AuthToken, error) {
	_, alreadyExists, err := s.usersStore.Query(
		datastore.Equals(datastore.Field("slug"), slug),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if alreadyExists {
		return nil, errors.New("slug already exists")
	}

	passwordHash, err := s.hashPassword(password)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	user := &user.User{
		Id:           sdk.Id("usr"),
		CreatedAt:    now,
		UpdatedAt:    now,
		Name:         name,
		Slug:         slug,
		PasswordHash: passwordHash,
	}

	err = s.usersStore.Create(user)
	if err != nil {
		return nil, err
	}

	for _, roleId := range roleIds {
		err = s.addRoleToUser(user.Id, roleId)
		if err != nil {
			return nil, err
		}
	}

	token, err := s.generateAuthToken(user)
	if err != nil {
		return nil, err
	}

	return token, s.authTokensStore.Create(token)
}

func (s *UserService) hashPassword(password string) (string, error) {
	cost := bcrypt.DefaultCost
	if s.isTest {
		cost = bcrypt.MinCost
	}

	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		cost,
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
