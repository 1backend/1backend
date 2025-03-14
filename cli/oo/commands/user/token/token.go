package token

import (
	"fmt"

	"github.com/openorch/openorch/cli/oo/config"
	"github.com/spf13/cobra"
)

// Token
func Token(cmd *cobra.Command, args []string) error {
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

	fmt.Println(usr.Token)

	return nil
}
