package config

import (
	"fmt"
	"os"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save [filePath | dirPath]
func Save(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	path := args[0]

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return errors.Wrap(err, fmt.Sprintf("path not found: '%v'", path))
	} else if err != nil {
		return errors.Wrap(err, "error checking path")
	}

	configs, err := util.CollectFromPath[openapi.ConfigSvcSaveConfigRequest](path, "configs")
	if err != nil {
		return err
	}

	for _, configRequest := range configs {
		_, _, err = cf.Client(client.WithToken(token)).
			ConfigSvcAPI.SaveConfig(ctx).
			Body(configRequest).
			Execute()
		if err != nil {
			return errors.Wrap(err, "failed to save configs")
		}
	}

	return nil

}
