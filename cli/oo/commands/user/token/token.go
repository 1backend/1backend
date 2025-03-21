package token

import (
	"fmt"

	"github.com/1backend/1backend/cli/oo/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Token
func Token(cmd *cobra.Command, args []string) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return errors.Wrap(err, "failed to load config")
	}

	if conf.Environments == nil {
		return errors.New("no environments")
	}
	env, ok := conf.Environments[conf.SelectedEnvironment]
	if !ok {
		return fmt.Errorf(
			"failed to find selected env '%s'",
			conf.SelectedEnvironment,
		)
	}

	if env.SelectedUser == "" {
		return fmt.Errorf(
			"no selected user in env '%v'",
			conf.SelectedEnvironment,
		)
	}

	usr, ok := env.Users[env.SelectedUser]
	if !ok {
		return fmt.Errorf(
			"cannot find user '%v' in env '%v'",
			env.SelectedUser,
			conf.SelectedEnvironment,
		)
	}

	fmt.Println(usr.Token)

	return nil
}
