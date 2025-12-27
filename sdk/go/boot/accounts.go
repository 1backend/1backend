/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package boot

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	onebackendapi "github.com/1backend/1backend/clients/go"
)

// RegisterServiceAccount is used during service startup to ensure the service is
// properly registered and authenticated with the user service.
// It first attempts to load existing credentials from the store.
// If none are found, it generates new ones and attempts to register the service.
// If login fails because the account doesn’t exist (e.g. registration was interrupted),
// it retries registration using the saved credentials.
func RegisterServiceAccount(
	userService onebackendapi.UserSvcAPI,
	serviceSlug,
	serviceName string,
	store datastore.DataStore,
) (*onebackendapi.UserSvcToken, error) {
	ctx := context.Background()

	res, err := store.Query().Find()
	if err != nil {
		return nil, errors.Wrap(err, "error querying credentials")
	}

	if len(res) == 0 {
		pw := uuid.NewString()
		err = store.Upsert(&auth.Credential{
			Slug:     serviceSlug,
			Password: pw,
		})
		if err != nil {
			return nil, errors.Wrap(err, "error upserting service account credential")
		}
		err = store.Refresh()
		if err != nil {
			return nil, errors.Wrap(err, "error refreshing credential store")
		}

		rsp, _, err := userService.Register(ctx).Body(onebackendapi.UserSvcRegisterRequest{
			AppHost:  sdk.DefaultAppHost,
			Slug:     serviceSlug,
			Name:     onebackendapi.PtrString(serviceName),
			Password: onebackendapi.PtrString(pw),
		}).Execute()
		if err != nil {
			return nil, errors.Wrap(err, "error registering service account")
		}

		return rsp.Token, nil
	}

	cred := res[0].(*auth.Credential)

	loginRsp, _, err := userService.Login(ctx).Body(onebackendapi.UserSvcLoginRequest{
		AppHost:  sdk.DefaultAppHost,
		Slug:     onebackendapi.PtrString(serviceSlug),
		Password: onebackendapi.PtrString(cred.Password),
	}).Execute()

	if err != nil {
		// We'll try to register as maybe registration failed or did not
		// happen after saving the credential.
		rsp, _, regErr := userService.Register(ctx).Body(onebackendapi.UserSvcRegisterRequest{
			AppHost:  sdk.DefaultAppHost,
			Slug:     serviceSlug,
			Name:     onebackendapi.PtrString(serviceName),
			Password: onebackendapi.PtrString(cred.Password),
		}).Execute()
		if regErr != nil {
			return nil, fmt.Errorf("failed to recover service account: register error: %w; original login error: %w", regErr, err)
		}

		return rsp.Token, nil
	}

	return loginRsp.Token, nil
}

// RegisterUserAccount is primarily used in tests.
// @todo Move to a test package.
func RegisterUserAccount(
	userService onebackendapi.UserSvcAPI,
	appHost string,
	slug string,
	password string,
	name string,
) (*onebackendapi.UserSvcToken, error) {
	_, _, err := userService.Register(context.Background()).Body(onebackendapi.UserSvcRegisterRequest{
		AppHost:  appHost,
		Slug:     slug,
		Password: onebackendapi.PtrString(password),
		Name:     onebackendapi.PtrString(name),
	}).Execute()

	if err != nil {
		return nil, errors.Wrap(err, "error registering user account")
	}

	loginRsp, _, err := userService.Login(context.Background()).Body(onebackendapi.UserSvcLoginRequest{
		AppHost:  appHost,
		Slug:     onebackendapi.PtrString(slug),
		Password: onebackendapi.PtrString(password),
	}).Execute()
	if err != nil {
		return nil, errors.Wrap(err, "error logging in with user account")
	}

	return loginRsp.Token, nil
}

func ListenAddress(s string) string {
	return strings.Replace(
		strings.Replace(s, "https://", "", -1),
		"http://",
		"",
		-1,
	)
}
