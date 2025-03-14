package env

import (
	"fmt"

	"github.com/openorch/openorch/cli/oo/config"
	"github.com/openorch/openorch/cli/oo/types"
	"github.com/spf13/cobra"
)

// Add prod http://someaddress.com:8090 "A description"
func Add(cmd *cobra.Command, args []string) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	shortName := args[0]
	url := args[1]
	var longDesc string
	if len(args) > 2 {
		longDesc = args[2]
	}

	env, ok := conf.Environments[shortName]
	if !ok {
		if conf.Environments == nil {
			conf.Environments = map[string]*types.Environment{}
		}
		conf.Environments[shortName] = &types.Environment{
			ShortName:   shortName,
			URL:         url,
			Description: longDesc,
		}
		conf.SelectedEnvironment = shortName
		return config.SaveConfig(conf)
	}

	env.URL = url
	env.Description = longDesc

	return config.SaveConfig(conf)
}
