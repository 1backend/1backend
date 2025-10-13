package deployment

import (
	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Delete(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	serviceDefinitionId := args[0]

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	_, _, err = cf.Client(client.WithToken(token)).
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
