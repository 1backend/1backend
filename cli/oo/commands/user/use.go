package user

import (
	"fmt"

	"github.com/1backend/1backend/cli/oo/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Use [slug]
func Use(cmd *cobra.Command, args []string) error {
	conf, err := util.LoadConfig()
	if err != nil {
		return errors.Wrap(err, "failed to load config")
	}

	slug := args[0]

	if conf.Environments == nil {
		return fmt.Errorf("no environments found")
	}

	env, ok := conf.Environments[conf.SelectedEnvironment]
	if !ok {
		return fmt.Errorf(
			"failed to find selected env: %s",
			conf.SelectedEnvironment,
		)
	}

	env.SelectedUser = slug

	return util.SaveConfig(conf)
}
