package instance

import (
	"fmt"

	"github.com/openorch/openorch/cli/oo/config"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/spf13/cobra"
)

// Remove
func Remove(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	instanceId := args[0]

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return fmt.Errorf("Cannot get env url: '%v'", err)
	}

	cf := sdk.NewApiClientFactory(url)

	_, err = cf.Client(sdk.WithToken(token)).
		RegistrySvcAPI.RemoveInstance(ctx, instanceId).
		Execute()
	if err != nil {
		return fmt.Errorf("Error removing service deployment: '%v'", err)
	}

	return nil
}
