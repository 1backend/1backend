package definition

import (
	"fmt"

	"github.com/openorch/openorch/cli/oo/config"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/spf13/cobra"
)

func Delete(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	serviceDefinitionId := args[0]

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return fmt.Errorf("Cannot get env url: '%v'", err)
	}

	cf := sdk.NewApiClientFactory(url)

	_, err = cf.Client(sdk.WithToken(token)).
		RegistrySvcAPI.DeleteDefinition(ctx, serviceDefinitionId).
		Execute()
	if err != nil {
		return fmt.Errorf("Error deleting service definition: '%v'", err)
	}

	return nil
}
