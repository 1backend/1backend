package user

import (
	"fmt"

	"github.com/1backend/1backend/cli/oo/config"
	"github.com/1backend/1backend/cli/oo/types"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"gopkg.in/yaml.v2"
)

// Whoami
func Whoami(cmd *cobra.Command, args []string, all bool) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if conf.Environments == nil {
		return fmt.Errorf("no environments")
	}
	env, ok := conf.Environments[conf.SelectedEnvironment]
	if !ok {
		return fmt.Errorf(
			"failed to find selected env: %s",
			conf.SelectedEnvironment,
		)
	}

	if env.SelectedUser == "" {
		return fmt.Errorf(
			"no selected user in env '%v'",
			conf.SelectedEnvironment,
		)
	}

	cf := client.NewApiClientFactory(env.URL)
	publicKeyRsp, _, err := cf.Client().UserSvcAPI.GetPublicKey(cmd.Context()).Execute()
	if err != nil {
		return errors.Wrap(err, "failed to get public key")
	}

	selectedUser, ok := env.Users[env.SelectedUser]
	if !ok {
		return fmt.Errorf(
			"cannot find user '%v' in env '%v'",
			env.SelectedUser,
			conf.SelectedEnvironment,
		)
	}
	if all {
		fmt.Println("# Selected user: " + selectedUser.Slug)
	}
	err = displayUser(cmd, publicKeyRsp.PublicKey, true, selectedUser)
	if err != nil {
		return errors.Wrap(err, "failed to display selected user")
	}

	if all {
		for _, usr := range env.Users {
			if usr.Slug == env.SelectedUser {
				continue
			}

			fmt.Println("")
			err = displayUser(cmd, publicKeyRsp.PublicKey, false, usr)
			if err != nil {
				return errors.Wrap(err, "failed to display user")
			}

		}
	}

	return nil
}

func displayUser(
	cmd *cobra.Command,
	publicKey string,
	active bool,
	usr *types.User,
) error {
	claims, err := auth.AuthorizerImpl{}.ParseJWT(publicKey, usr.Token)
	if err != nil {
		return errors.Wrap(err, "failed to decode JWT. it is possible that the public key of the server has changed. try logging in again")
	}

	userInfo := UserInfo{
		Id:    claims.UserId,
		Slug:  claims.Slug,
		Roles: claims.RoleIds,
	}

	enc := yaml.NewEncoder(cmd.OutOrStdout())
	if err := enc.Encode(userInfo); err != nil {
		return errors.Wrap(err, "failed to encode user info")
	}

	return nil
}

type UserInfo struct {
	Slug  string   `json:"slug"`
	Id    string   `json:"id"`
	Roles []string `json:"roles"`
}
