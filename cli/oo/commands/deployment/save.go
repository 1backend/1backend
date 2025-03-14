package deployment

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

// Save /home/user/deploymentA.yaml
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
			fmt.Sprintf("failed to stat deployment file: '%v'", filePath),
		)
	}
	if fileInfo.Size() == 0 {
		return fmt.Errorf("service deployment file is empty at '%v'", filePath)
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("failed to read deployment file: '%v'", filePath),
		)
	}

	deployment := openapi.DeploySvcDeployment{}

	if err := yaml.Unmarshal(data, &deployment); err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("failed to decode deployment file: '%v'", filePath),
		)
	}

	stat := openapi.DeploymentStatusDeploying
	deployment.Status = &stat

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := sdk.NewApiClientFactory(url)

	_, _, err = cf.Client(sdk.WithToken(token)).
		DeploySvcAPI.SaveDeployment(ctx).
		Body(openapi.DeploySvcSaveDeploymentRequest{
			Deployment: &deployment,
		}).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to save deployment")
	}

	return nil
}
