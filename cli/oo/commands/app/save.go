package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save handles the `oo app save` command.
func Save(cmd *cobra.Command, args []string, id string, host string) error {
	ctx := cmd.Context()

	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)

	inputs := []openapi.UserSvcApp{}
	if host != "" {
		inputs = append(inputs, openapi.UserSvcApp{Id: id, Host: host})
	}

	if len(args) > 0 {
		path := args[0]
		fileInputs, err := collectAppsFromPath(path)
		if err != nil {
			return err
		}
		inputs = append(inputs, fileInputs...)
	}

	if len(inputs) == 0 {
		return errors.New("no app data provided")
	}

	for _, input := range inputs {
		if input.Host == "" {
			return errors.New("host is required for each app")
		}

		req := openapi.UserSvcReadAppRequest{}
		req.SetHost(input.Host)

		rsp, _, err := cf.Client(client.WithToken(token)).
			UserSvcAPI.ReadApp(ctx).
			Body(req).
			Execute()
		if err != nil {
			return errors.Wrapf(err, "failed to save app '%s'", input.Host)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Saved app %s (ID: %s)\n", rsp.App.Host, rsp.App.Id)
	}

	return nil
}

func collectAppsFromPath(path string) ([]openapi.UserSvcApp, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, errors.Wrap(err, fmt.Sprintf("path not found: '%v'", path))
	} else if err != nil {
		return nil, errors.Wrap(err, "error checking path")
	}

	inputs := []openapi.UserSvcApp{}

	if stat.IsDir() {
		err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error accessing file '%v'", filePath))
			}
			if info.IsDir() {
				return nil
			}

			fileInputs, err := extractAppsFromFile(filePath)
			if err != nil {
				return err
			}
			inputs = append(inputs, fileInputs...)
			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		fileInputs, err := extractAppsFromFile(path)
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, fileInputs...)
	}

	return inputs, nil
}

func extractAppsFromFile(filePath string) ([]openapi.UserSvcApp, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to read file at '%v'", filePath))
	}

	var inputs []openapi.UserSvcApp
	if err := yaml.Unmarshal(data, &inputs); err != nil {
		var single openapi.UserSvcApp
		if err := yaml.Unmarshal(data, &single); err != nil {
			return nil, errors.Wrap(
				err,
				fmt.Sprintf("failed to parse apps file at '%v'", filePath),
			)
		}
		inputs = append(inputs, single)
	}

	return inputs, nil
}
