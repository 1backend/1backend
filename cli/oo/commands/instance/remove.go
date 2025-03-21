package instance

import (
	"github.com/1backend/1backend/cli/oo/config"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Remove
func Remove(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	instanceId := args[0]

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := sdk.NewApiClientFactory(url)

	_, err = cf.Client(sdk.WithToken(token)).
		RegistrySvcAPI.RemoveInstance(ctx, instanceId).
		Execute()
	if err != nil {
		return errors.Wrap(err, "error removing service deployment")
	}

	return nil
}
