package definition

import (
	"github.com/1backend/1backend/cli/oo/config"
	"github.com/1backend/1backend/sdk/go/client"
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

	cf := client.NewApiClientFactory(url)

	_, err = cf.Client(client.WithToken(token)).
		RegistrySvcAPI.DeleteDefinition(ctx, serviceDefinitionId).
		Execute()
	if err != nil {
		return errors.Wrap(err, "Error deleting service definition")
	}

	return nil
}
