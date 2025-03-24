/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package sdk

import (
	"context"

	"github.com/1backend/1backend/sdk/go/datastore"

	client "github.com/1backend/1backend/clients/go"
)

func RegisterServiceAccount(userService client.UserSvcAPI, serviceSlug, serviceName string, store datastore.DataStore) (*client.UserSvcAuthToken, error) {
	ctx := context.Background()

	res, err := store.Query().Find()
	if err != nil {
		return nil, err
	}

	slug := serviceSlug
	pw := ""

	if len(res) > 0 {
		cred := res[0].(*Credential)
		slug = cred.Slug
		pw = cred.Password
	} else {
		pw = Id("cred")
		err = store.Upsert(&Credential{
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

	loginRsp, _, err := userService.Login(ctx).Body(client.UserSvcLoginRequest{
		Slug:     client.PtrString(slug),
		Password: client.PtrString(pw),
	}).Execute()

	if err != nil {
		_, _, err = userService.Register(ctx).Body(client.UserSvcRegisterRequest{
			Slug:     client.PtrString(slug),
			Name:     client.PtrString(serviceName),
			Password: client.PtrString(pw),
		}).Execute()
		if err != nil {
			return nil, err
		}

		loginRsp, _, err = userService.Login(ctx).Body(client.UserSvcLoginRequest{
			Slug:     client.PtrString(slug),
			Password: client.PtrString(pw),
		}).Execute()
		if err != nil {
			return nil, err
		}
	}

	return loginRsp.Token, nil
}

func RegisterUserAccount(userService client.UserSvcAPI, slug, password, username string) (*client.UserSvcAuthToken, error) {
	_, _, err := userService.Register(context.Background()).Body(client.UserSvcRegisterRequest{
		Slug:     client.PtrString(slug),
		Password: client.PtrString(password),
		Name:     client.PtrString(username),
	}).Execute()

	if err != nil {
		return nil, err
	}

	loginRsp, _, err := userService.Login(context.Background()).Body(client.UserSvcLoginRequest{
		Slug:     client.PtrString(slug),
		Password: client.PtrString(password),
	}).Execute()
	if err != nil {
		return nil, err
	}

	return loginRsp.Token, nil
}
