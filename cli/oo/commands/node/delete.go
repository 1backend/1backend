package node

import (
	"fmt"

	"github.com/openorch/openorch/cli/oo/config"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Delete [nodeUrl]
func Delete(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	nodeUrl := args[0]

	ur, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := sdk.NewApiClientFactory(ur)

	_, err = cf.Client(sdk.WithToken(token)).
		RegistrySvcAPI.DeleteNode(ctx, nodeUrl).
		Execute()
	if err != nil {
		return fmt.Errorf("Error deleting node: '%v'", err)
	}

	return nil
}
