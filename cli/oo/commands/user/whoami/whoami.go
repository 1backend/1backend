package whoami

import (
	"fmt"

	"github.com/openorch/openorch/cli/oo/config"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"gopkg.in/yaml.v2"
)

// Whoami
func Whoami(cmd *cobra.Command, args []string) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("Failed to load config: %w", err)
	}

	if conf.Environments == nil {
		return fmt.Errorf("No environments")
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
			"No selected user in env '%v'",
			conf.SelectedEnvironment,
		)
	}

	usr, ok := env.Users[env.SelectedUser]
	if !ok {
		return fmt.Errorf(
			"Cannot find user '%v' in env '%v'",
			env.SelectedUser,
			conf.SelectedEnvironment,
		)
	}

	cf := sdk.NewApiClientFactory(env.URL)

	publicKeyRsp, _, err := cf.Client().UserSvcAPI.GetPublicKey(cmd.Context()).Execute()
	if err != nil {
		return errors.Wrap(err, "failed to get public key")
	}

	claims, err := sdk.AuthorizerImpl{}.ParseJWT(*publicKeyRsp.PublicKey, usr.Token)
	if err != nil {
		return errors.Wrap(err, "failed to decode JWT")
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
