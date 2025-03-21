package deployment

import (
	"github.com/1backend/1backend/cli/oo/config"
	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Delete(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	serviceDefinitionId := args[0]

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := sdk.NewApiClientFactory(url)

	_, _, err = cf.Client(sdk.WithToken(token)).
		DeploySvcAPI.DeleteDeployment(ctx).
		Body(openapi.DeploySvcDeleteDeploymentRequest{
			DeploymentId: serviceDefinitionId,
		}).
		Execute()
	if err != nil {
		return errors.Wrap(err, "error deleting service deployment")
	}

	return nil
}
