package config

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

// Save [filePath | dirPath]
func Save(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	path := args[0]

	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.Wrap(err, fmt.Sprintf("path not found: '%v'", path))
	} else if err != nil {
		return errors.Wrap(err, "error checking path")
	}

	var configs []openapi.ConfigSvcSaveConfigRequest
	fileCount := 0
	if stat.IsDir() {
		// Handle directory: Iterate over files and collect configs
		err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error accessing file '%v'", filePath))
			}
			if !info.IsDir() {
				// Collect configs from each file in the directory
				fileCount++
				fileConfigs, err := extractConfigsFromFile(filePath)
				if err != nil {
					return err
				}
				configs = append(configs, fileConfigs...)
			}
			return nil
		})
		if err != nil {
			return err
		}
	} else {
		// Handle single file
		fileCount++
		fileConfigs, err := extractConfigsFromFile(path)
		if err != nil {
			return err
		}
		configs = append(configs, fileConfigs...)
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

// extract one or more configs from a file
func extractConfigsFromFile(filePath string) ([]openapi.ConfigSvcSaveConfigRequest, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to read file at '%v'", filePath))
	}

	// Determine whether the file contains single or multiple configs
	var configs []openapi.ConfigSvcSaveConfigRequest

	// Unmarshal as list first (multiple configs)
	if err := yaml.Unmarshal(data, &configs); err != nil {
		// If unmarshalling to list fails, attempt unmarshalling as single config
		var singleConfig openapi.ConfigSvcSaveConfigRequest
		if err := yaml.Unmarshal(data, &singleConfig); err != nil {
			return nil, errors.Wrap(
				err,
				fmt.Sprintf(
					"failed to parse configs file at '%v' as single or multiple configs",
					filePath,
				),
			)
		}
		configs = append(configs, singleConfig)
	}

	return configs, nil
}
