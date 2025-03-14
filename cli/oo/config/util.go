package config

import (
	"fmt"

	"github.com/openorch/openorch/cli/oo/types"
)

func GetSelectedEnv() (*types.Environment, error) {
	conf, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if conf.Environments == nil {
		return nil, fmt.Errorf("No environments")
	}

	env, ok := conf.Environments[conf.SelectedEnvironment]
	if !ok {
		return nil, fmt.Errorf(
			"failed to find selected env: %s",
			conf.SelectedEnvironment,
		)
	}

	return env, nil
}

func GetSelectedUrlAndToken() (string, string, error) {
	env, err := GetSelectedEnv()
	if err != nil {
		return "", "", err
	}

	return env.URL, env.Users[env.SelectedUser].Token, nil
}
