package sdk

import (
	"context"

	"github.com/openorch/openorch/sdk/go/datastore"

	client "github.com/openorch/openorch/clients/go"
)

func RegisterServiceAccount(userService client.UserSvcAPI, serviceSlug, serviceName string, store datastore.DataStore) (string, error) {
	ctx := context.Background()

	res, err := store.Query().Find()
	if err != nil {
		return "", err
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
			return "", err
		}
		err = store.Refresh()
		if err != nil {
			return "", err
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
			return "", err
		}

		loginRsp, _, err = userService.Login(ctx).Body(client.UserSvcLoginRequest{
			Slug:     client.PtrString(slug),
			Password: client.PtrString(pw),
		}).Execute()
		if err != nil {
			return "", err
		}
	}

	return *loginRsp.Token.Token, nil
}

func RegisterUserAccount(userService client.UserSvcAPI, slug, password, username string) (string, error) {
	_, _, err := userService.Register(context.Background()).Body(client.UserSvcRegisterRequest{
		Slug:     client.PtrString(slug),
		Password: client.PtrString(password),
		Name:     client.PtrString(username),
	}).Execute()

	if err != nil {
		return "", err
	}

	loginRsp, _, err := userService.Login(context.Background()).Body(client.UserSvcLoginRequest{
		Slug:     client.PtrString(slug),
		Password: client.PtrString(password),
	}).Execute()
	if err != nil {
		return "", err
	}

	return *loginRsp.Token.Token, nil
}
