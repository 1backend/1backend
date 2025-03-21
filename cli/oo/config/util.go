package config

import (
	"fmt"

	"github.com/1backend/1backend/cli/oo/types"
	"github.com/pkg/errors"
)

func GetSelectedEnv() (*types.Environment, error) {
	conf, err := LoadConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}

	if conf.Environments == nil {
		return nil, fmt.Errorf("no environments")
	}

	env, ok := conf.Environments[conf.SelectedEnvironment]
	if !ok {
		return nil, fmt.Errorf(
			"failed to find selected env '%s'",
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

	selectedUser, ok := env.Users[env.SelectedUser]
	if !ok {
		return "", "", fmt.Errorf("no user selected. maybe try logging in first?")
	}

	return env.URL, selectedUser.Token, nil
}
