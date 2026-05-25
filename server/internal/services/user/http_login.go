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
	"fmt"
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

	if request.AppHost == "" {
		endpoint.WriteErr(w, http.StatusBadRequest, errors.New("AppHost missing"))
		return
	}

	if request.Contact.Id != "" {
		if isEmail(request.Contact.Id) {
			request.Contact.Id = normalizeEmail(request.Contact.Id)
		}
	}

	app, err := s.getOrCreateApp(request.AppHost)
	if err != nil {
		logger.Error(
			"Failed to get or create app",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	token, err := s.login(app.Id, &request, false)
	if err != nil {
		switch err.Error() {
		case "unauthorized":
			logger.Error(
				"Invalid password",
				slog.Any("slug", request.Slug),
				slog.Any("contact", request.Contact),
			)
			endpoint.WriteString(w, http.StatusUnauthorized, "Invalid Password")
		case errTOTPRequired.Error():
			endpoint.WriteString(w, http.StatusUnauthorized, "TOTP code required")
		case errTOTPInvalid.Error():
			endpoint.WriteString(w, http.StatusUnauthorized, "Invalid TOTP code")
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

	// Publish user login event (non-blocking)
	if s.options.PubSub != nil {
		evt := userLifecycleEventPayload(app, token, request.Slug, time.Now().UTC())
		payload, _ := json.Marshal(evt)
		if _, perr := s.options.PubSub.Publish(r.Context(), "user.login", payload); perr != nil {
			logger.Error("Failed to publish login event", slog.Any("error", perr))
		}
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
	appId string,
	request *user.LoginRequest,
	otpAlreadyVerified bool,
) (*user.Token, error) {

	var usr *user.User
	totpAlreadyVerified := false

	if request.Device == "" {
		request.Device = unknownDevice
	}

	if request.Contact.Id != "" {
		contactI, found, err := s.contactStore.Query(
			datastore.Equals(datastore.Field("id"), request.Contact.Id),
		).FindOne()
		if err != nil {
			return nil, err
		}

		if s.options.VerifyContacts {
			if found {
				userI, found, err := s.userStore.Query(
					datastore.Equals(datastore.Field("id"), contactI.(*user.Contact).UserId),
				).FindOne()
				if err != nil {
					return nil, errors.Wrap(err, "error querying user by contact")
				}

				if !found {
					return nil, fmt.Errorf("user not found by contact id '%s'", request.Contact.Id)
				}

				usr = userI.(*user.User)
				if err := s.verifyLoginTOTP(usr.Id, request.TOTPCode); err != nil {
					return nil, err
				}
				totpAlreadyVerified = true
			}

			if !otpAlreadyVerified {
				err = s.verifyContactOTP(&request.Contact)
				if err != nil {
					return nil, errors.Wrap(err, "login contact verification failed")
				}
			}

			if found {
				return s.issueToken(appId, usr, request.Device)
			} else {
				if request.Slug == "" {
					return nil, fmt.Errorf("slug is missing")
				}

				newUser := &user.UserInput{
					Id:   sdk.Id("usr"),
					Slug: request.Slug,
				}

				now := time.Now()

				err = s.createUserWithoutVerification(
					appId,
					newUser,
					[]user.Contact{
						{
							Id:        request.Contact.Id,
							CreatedAt: now,
							UpdatedAt: now,
							Platform:  request.Contact.Platform,
							Verified:  true,
						},
					},
					request.Password,
					nil,
				)

				if err != nil {
					logger.Error("Error when login-registering",
						slog.String("slug", request.Slug),
					)
					return nil, errors.Wrap(err, "error creating user when login-registering")
				}

				userI, found, err := s.userStore.Query(
					datastore.Equals(datastore.Field("slug"), request.Slug),
				).FindOne()
				if err != nil {
					return nil, fmt.Errorf("error querying user by slug '%s': %v", request.Slug, err)
				}
				if !found {
					return nil, fmt.Errorf("user not found by slug '%s'", request.Slug)
				}

				usr = userI.(*user.User)

				return s.issueToken(appId, usr, request.Device)
			}
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
	} else if request.Slug != "" {
		userI, found, err := s.userStore.Query(
			datastore.Equals(datastore.Field("slug"), request.Slug),
		).FindOne()
		if err != nil {
			return nil, fmt.Errorf("error querying user by slug '%s': %v", request.Slug, err)
		}
		if !found {
			return nil, fmt.Errorf("user not found by slug '%s'", request.Slug)
		}

		usr = userI.(*user.User)
	} else {
		return nil, errors.New("slug or contact required")
	}

	switch {
	case request.Password != "":
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
	case otpAlreadyVerified:
		// If OTP is already verified, we can proceed with login
		break
	case request.Contact.OtpId != "" && request.Contact.OtpCode != "":
		if err := s.verifyLoginTOTP(usr.Id, request.TOTPCode); err != nil {
			return nil, err
		}
		totpAlreadyVerified = true

		err := s.verifyContactOTP(&request.Contact)
		if err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("password or otp required")
	}

	if !totpAlreadyVerified {
		if err := s.verifyLoginTOTP(usr.Id, request.TOTPCode); err != nil {
			return nil, err
		}
	}

	return s.issueToken(appId, usr, request.Device)
}

func (s *UserService) issueToken(
	appId string,
	usr *user.User,
	device string,
) (*user.Token, error) {
	if device == "" {
		device = unknownDevice
	}

	releaseLock, err := s.acquireRefreshTokenLock(
		context.Background(),
		refreshTokenLockKeyForParts(appId, usr.Id, device),
	)
	if err != nil {
		return nil, err
	}
	lockHeld := true
	defer func() {
		if lockHeld {
			releaseLock()
		}
	}()

	// Let's see if there is an active token we can reuse
	tokenI, found, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("appId"), appId),
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
				lockHeld = false
				releaseLock()

				tok, err := s.refreshToken(context.Background(), tok.Token)
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

		err = s.inactivateToken(appId, tok.Id)
		if err != nil {
			return nil, errors.Wrap(err, "could not inactivate token")
		}
	}

	token, err := s.generateAuthToken(
		appId,
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
	appId string,
	u *user.User,
	device string,
) (*user.Token, error) {
	roles, err := s.getRolesByUserId(appId, u.Id)
	if err != nil {
		return nil, errors.Wrap(err, "error listing roles")
	}

	activeOrganizationId, err := s.getActiveOrganizationId(appId, u.Id, device)
	if err != nil {
		return nil, errors.Wrap(err, "error reading active organization")
	}

	roles = filterRolesForActiveOrganization(roles, activeOrganizationId)

	return s.buildAuthToken(appId, u, device, roles, activeOrganizationId)
}

func (s *UserService) issueTokenForActiveOrganization(
	appId string,
	usr *user.User,
	device string,
	activeOrganizationId string,
	activeOrganizationRoles []string,
) (*user.Token, error) {
	if strings.TrimSpace(activeOrganizationId) == "" {
		return nil, errors.New("active organization id is required")
	}
	if device == "" {
		device = unknownDevice
	}

	releaseLock, err := s.acquireRefreshTokenLock(
		context.Background(),
		refreshTokenLockKeyForParts(appId, usr.Id, device),
	)
	if err != nil {
		return nil, err
	}
	defer releaseLock()

	roles, err := s.getRolesByUserId(appId, usr.Id)
	if err != nil {
		return nil, errors.Wrap(err, "error listing roles")
	}

	roles = appendUniqueRoles(
		roles,
		append([]string{orgMemberRole(activeOrganizationId, usr.Id)}, activeOrganizationRoles...)...,
	)
	roles = filterRolesForActiveOrganization(roles, activeOrganizationId)

	if err := s.setActivation(appId, usr.Id, device, activeOrganizationId); err != nil {
		return nil, errors.Wrap(err, "failed to update activation")
	}
	if err := s.inactivateTokens(appId, usr.Id); err != nil {
		return nil, errors.Wrap(err, "failed to inactivate tokens")
	}

	token, err := s.buildAuthToken(appId, usr, device, roles, activeOrganizationId)
	if err != nil {
		return nil, err
	}
	if err := s.tokenStore.Create(token); err != nil {
		return nil, errors.Wrap(err, "error creating token")
	}

	return token, nil
}

func (s *UserService) buildAuthToken(
	appId string,
	u *user.User,
	device string,
	roles []string,
	activeOrganizationId string,
) (*user.Token, error) {
	_, token, err := s.generateJWT(
		appId, u, roles, activeOrganizationId,
		s.privateKey, device)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	app, found, err := s.appByID(appId)
	if err != nil {
		return nil, fmt.Errorf("error finding app by id '%s': %v", appId, err)
	}
	if !found {
		return nil, fmt.Errorf("app not found by id '%s'", appId)
	}

	app, err = s.getOrCreateApp(app.Host)
	if err != nil {
		return nil, errors.Wrap(err, "error getting or creating app")
	}

	id := sdk.Id("tok")

	internalId, err := sdk.InternalId(appId, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create internal id")
	}

	return &user.Token{
		InternalId: internalId,
		Id:         id,
		AppId:      appId,
		App:        app,
		UserId:     u.Id,
		Token:      token,
		Device:     device,
		Active:     true,
		ExpiresAt:  now.Add(s.options.TokenExpiration),
		CreatedAt:  now,
	}, nil
}

func appendUniqueRoles(roles []string, extras ...string) []string {
	seen := map[string]struct{}{}
	merged := make([]string, 0, len(roles)+len(extras))
	for _, role := range append(roles, extras...) {
		role = strings.TrimSpace(role)
		if role == "" {
			continue
		}
		if _, ok := seen[role]; ok {
			continue
		}

		seen[role] = struct{}{}
		merged = append(merged, role)
	}
	return merged
}

func filterRolesForActiveOrganization(roles []string, activeOrganizationId string) []string {
	filtered := make([]string, 0, len(roles))

	for _, role := range roles {
		if !strings.HasPrefix(role, "user-svc:org:{") {
			filtered = append(filtered, role)
			continue
		}

		if activeOrganizationId == "" {
			continue
		}

		if strings.HasPrefix(role, fmt.Sprintf("user-svc:org:{%s}:", activeOrganizationId)) {
			filtered = append(filtered, role)
		}
	}

	return filtered
}

func (s *UserService) getOrCreateApp(host string) (*user.App, error) {
	app, found, err := s.appByHost(host)
	if err != nil {
		return nil, errors.Wrap(err, "error querying app by host")
	}
	if found {
		return app, nil
	}

	app = &user.App{
		Id:   sdk.Id("app"),
		Host: host,
	}

	err = s.appStore.Upsert(app)
	if err != nil {
		return nil, errors.Wrap(err, "error creating app")
	}

	s.cacheApp(app)
	return app, nil
}
