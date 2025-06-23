package instance

import (
	"github.com/1backend/1backend/cli/oo/util"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Remove
func Remove(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	instanceId := args[0]

	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	_, err = cf.Client(client.WithToken(token)).
		RegistrySvcAPI.RemoveInstance(ctx, instanceId).
		Execute()
	if err != nil {
		return errors.Wrap(err, "error removing service deployment")
	}

	return nil
}
