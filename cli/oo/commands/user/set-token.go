package user

import (
	"fmt"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// SetToken
func SetToken(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing token")
	}
	token := args[0]

	conf, err := util.LoadConfig()
	if err != nil {
		return errors.Wrap(err, "failed to load config")
	}

	env, ok := conf.Environments[conf.SelectedEnvironment]
	if !ok {
		return fmt.Errorf("failed to find selected env '%s'", conf.SelectedEnvironment)
	}

	if env.SelectedUser == "" {
		return fmt.Errorf("no selected user in env '%s'", conf.SelectedEnvironment)
	}

	usr, ok := env.Users[env.SelectedUser]
	if !ok {
		return fmt.Errorf("cannot find user '%s' in env '%s'", env.SelectedUser, conf.SelectedEnvironment)
	}

	// get server public key
	cf := client.NewApiClientFactory(env.URL)
	publicKeyRsp, _, err := cf.Client().
		UserSvcAPI.
		GetPublicKey(cmd.Context()).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to get public key")
	}

	trsp, _, err := cf.Client(client.WithToken(token)).
		UserSvcAPI.RefreshToken(cmd.Context()).
		Execute()

	if err != nil {
		return errors.Wrap(err, "failed to refresh token")
	}

	token = trsp.Token.Token

	claims, err := auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, token)
	if err != nil {
		return errors.Wrap(err, "failed to decode token")
	}

	rsp, _, err := cf.Client(client.WithToken(token)).
		UserSvcAPI.ListApps(cmd.Context()).Body(
		openapi.UserSvcListAppsRequest{
			Ids: []string{claims.AppId},
		},
	).Execute()

	appHost := rsp.Apps[0].Host

	if usr.TokensByAppHost == nil {
		usr.TokensByAppHost = map[string]string{}
	}

	usr.TokensByAppHost[appHost] = token
	usr.SelectedAppHost = appHost

	if err := util.SaveConfig(conf); err != nil {
		return errors.Wrap(err, "failed to save config")
	}

	fmt.Printf("Token stored for app host '%s'\n", appHost)

	return nil
}
