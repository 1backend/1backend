package definition

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/openorch/openorch/cli/oo/config"
	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save /home/user/definitionA.yaml
func Save(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	filePath := args[0]

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.Wrap(err, "cannot apply nonexistent file at '%v'")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("failed to open file: '%v'", filePath),
		)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"failed to stat service definition file: '%v'",
				filePath,
			),
		)
	}
	if fileInfo.Size() == 0 {
		return fmt.Errorf("service definition file is empty at '%v'", filePath)
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("failed to read definition file: '%v'", filePath),
		)
	}

	definition := openapi.RegistrySvcDefinition{}

	if err := yaml.Unmarshal(data, &definition); err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("failed to decode deployment file: '%v'", filePath),
		)
	}

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := sdk.NewApiClientFactory(url)

	_, _, err = cf.Client(sdk.WithToken(token)).
		RegistrySvcAPI.SaveDefinition(ctx).
		Body(openapi.RegistrySvcSaveDefinitionRequest{
			Definition: &definition,
		}).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to save service definition")
	}

	return nil
}
