/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package boot

import (
	"context"
	"strings"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"

	onebackendapi "github.com/1backend/1backend/clients/go"
)

func RegisterServiceAccount(userService onebackendapi.UserSvcAPI, serviceSlug, serviceName string, store datastore.DataStore) (*onebackendapi.UserSvcAuthToken, error) {
	ctx := context.Background()

	res, err := store.Query().Find()
	if err != nil {
		return nil, err
	}

	slug := serviceSlug
	pw := ""

	if len(res) > 0 {
		cred := res[0].(*auth.Credential)
		slug = cred.Slug
		pw = cred.Password
	} else {
		pw = sdk.Id("cred")
		err = store.Upsert(&auth.Credential{
			Slug:     slug,
			Password: pw,
		})
		if err != nil {
			return nil, err
		}
		err = store.Refresh()
		if err != nil {
			return nil, err
		}
	}

	loginRsp, _, err := userService.Login(ctx).Body(onebackendapi.UserSvcLoginRequest{
		Slug:     onebackendapi.PtrString(slug),
		Password: onebackendapi.PtrString(pw),
	}).Execute()

	if err != nil {
		_, _, err = userService.Register(ctx).Body(onebackendapi.UserSvcRegisterRequest{
			Slug:     slug,
			Name:     onebackendapi.PtrString(serviceName),
			Password: onebackendapi.PtrString(pw),
		}).Execute()
		if err != nil {
			return nil, err
		}

		loginRsp, _, err = userService.Login(ctx).Body(onebackendapi.UserSvcLoginRequest{
			Slug:     onebackendapi.PtrString(slug),
			Password: onebackendapi.PtrString(pw),
		}).Execute()
		if err != nil {
			return nil, err
		}
	}

	return loginRsp.Token, nil
}

func RegisterUserAccount(userService onebackendapi.UserSvcAPI, slug, password, username string) (*onebackendapi.UserSvcAuthToken, error) {
	_, _, err := userService.Register(context.Background()).Body(onebackendapi.UserSvcRegisterRequest{
		Slug:     slug,
		Password: onebackendapi.PtrString(password),
		Name:     onebackendapi.PtrString(username),
	}).Execute()

	if err != nil {
		return nil, err
	}

	loginRsp, _, err := userService.Login(context.Background()).Body(onebackendapi.UserSvcLoginRequest{
		Slug:     onebackendapi.PtrString(slug),
		Password: onebackendapi.PtrString(password),
	}).Execute()
	if err != nil {
		return nil, err
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
