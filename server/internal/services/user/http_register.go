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
	"regexp"
	"time"

	"github.com/pkg/errors"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"golang.org/x/crypto/bcrypt"
)

var SlugRegexp = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

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

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		endpoint.WriteErr(w, http.StatusBadRequest, errors.New(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	if req.Password == "" {
		endpoint.WriteErr(w, http.StatusBadRequest, errors.New(`Password missing`))
		return
	}

	if req.Slug == "" {
		endpoint.WriteErr(w, http.StatusBadRequest, errors.New(`Slug missing`))
		return
	}

	if !SlugRegexp.MatchString(req.Slug) {
		endpoint.WriteErr(w, http.StatusBadRequest,
			errors.New(`Slug must be lowercase and can only contain letters, numbers, and dashes`),
		)
		return
	}

	if req.App == "" {
		req.App = user.DefaultApp
	}

	if req.Device == "" {
		req.Device = unknownDevice
	}

	newUser := &user.UserInput{
		Name: req.Name,
		Slug: req.Slug,
	}

	var contacts []user.Contact
	now := time.Now()
	if req.Contact.Id != "" {
		contacts = append(contacts, user.Contact{
			Id:        req.Contact.Id,
			CreatedAt: now,
			UpdatedAt: now,
			Platform:  req.Contact.Platform,
			Handle:    req.Contact.Handle,
		})
	}
	err = s.createUser(
		req.App,
		newUser,
		contacts,
		req.Password,
		nil,
	)
	if err != nil {
		logger.Error(
			"Failed to create user",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	token, err := s.login(&user.LoginRequest{
		App:      req.App,
		Slug:     req.Slug,
		Password: req.Password,
		Device:   req.Device,
	})
	if err != nil {
		logger.Error(
			"Failed to login after registration",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(user.RegisterResponse{
		Token: token,
	})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (s *UserService) register(
	app string,
	slug,
	password,
	name string,
	roles []string,
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
	usr := &user.User{
		Id:        sdk.Id("usr"),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Slug:      slug,
	}

	err = s.passwordsStore.Upsert(&user.Password{
		Id:           sdk.Id("pw"),
		PasswordHash: passwordHash,
		UserId:       usr.Id,
		CreatedAt:    now,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to save password")
	}

	err = s.usersStore.Create(usr)
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		err = s.assignRole(app, usr.Id, role)
		if err != nil {
			return nil, err
		}
	}

	token, err := s.generateAuthToken(
		app,
		usr,
		unknownDevice,
	)
	if err != nil {
		return nil, err
	}

	return token, s.authTokensStore.Create(token)
}

func (s *UserService) hashPassword(password string) (string, error) {
	cost := bcrypt.DefaultCost
	if s.options.Test {
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
