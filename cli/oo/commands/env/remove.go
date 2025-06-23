package env

import (
	"fmt"

	"github.com/1backend/1backend/cli/oo/util"
	"github.com/spf13/cobra"
)

// Remove prod
func Remove(cmd *cobra.Command, args []string) error {
	conf, err := util.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	shortName := args[0]
	if _, ok := conf.Environments[shortName]; !ok {
		return fmt.Errorf("environment %s not found", shortName)
	}

	delete(conf.Environments, shortName)

	return util.SaveConfig(conf)
}
