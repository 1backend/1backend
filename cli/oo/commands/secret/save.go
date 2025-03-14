package secret

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/openorch/openorch/cli/oo/config"
	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save [key] [value] | [filePath | dirPath]
func Save(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := sdk.NewApiClientFactory(url)

	// Determine the save type: single key-value pair, file, or directory
	if len(args) == 2 {
		// Case 1: Single key-value pair
		key := args[0]
		value := args[1]

		secret := openapi.SecretSvcSecret{
			Key:   openapi.PtrString(key),
			Value: openapi.PtrString(value),
		}

		_, _, err = cf.Client(sdk.WithToken(token)).
			SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecret{secret},
			}).
			Execute()
		if err != nil {
			return errors.Wrap(err, "failed to save single secret")
		}

		return nil
	} else if len(args) == 1 {
		// Case 2: File or directory-based secrets
		path := args[0]

		stat, err := os.Stat(path)
		if os.IsNotExist(err) {
			return errors.Wrap(err, fmt.Sprintf("path not found: '%v'", path))
		} else if err != nil {
			return errors.Wrap(err, "error checking path")
		}

		var secrets []openapi.SecretSvcSecret
		fileCount := 0
		if stat.IsDir() {
			// Handle directory: Iterate over files and collect secrets
			err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("error accessing file '%v'", filePath))
				}
				if !info.IsDir() {
					// Collect secrets from each file in the directory
					fileCount++
					fileSecrets, err := extractSecretsFromFile(filePath)
					if err != nil {
						return err
					}
					secrets = append(secrets, fileSecrets...)
				}
				return nil
			})
			if err != nil {
				return err
			}
		} else {
			// Handle single file
			fileCount++
			fileSecrets, err := extractSecretsFromFile(path)
			if err != nil {
				return err
			}
			secrets = append(secrets, fileSecrets...)
		}

		// Make a single API call to save all secrets
		_, _, err = cf.Client(sdk.WithToken(token)).
			SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: secrets,
			}).
			Execute()
		if err != nil {
			return errors.Wrap(err, "failed to save secrets")
		}

		return nil
	}

	return errors.New("invalid arguments: use 'save [key] [value]' or 'save [filePath | dirPath]'")
}

// extract one or more secrets from a file
func extractSecretsFromFile(filePath string) ([]openapi.SecretSvcSecret, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to read file at '%v'", filePath))
	}

	// Determine whether the file contains single or multiple secrets
	var secrets []openapi.SecretSvcSecret

	// Unmarshal as list first (multiple secrets)
	if err := yaml.Unmarshal(data, &secrets); err != nil {
		// If unmarshalling to list fails, attempt unmarshalling as single secret
		var singleSecret openapi.SecretSvcSecret
		if err := yaml.Unmarshal(data, &singleSecret); err != nil {
			return nil, errors.Wrap(
				err,
				fmt.Sprintf(
					"failed to parse secrets file at '%v' as single or multiple secrets",
					filePath,
				),
			)
		}
		secrets = append(secrets, singleSecret)
	}

	return secrets, nil
}
