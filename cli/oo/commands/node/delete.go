package node

import (
	"github.com/1backend/1backend/cli/oo/util"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Delete [nodeUrl]
func Delete(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	nodeUrl := args[0]

	ur, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(ur)

	_, err = cf.Client(client.WithToken(token)).
		RegistrySvcAPI.DeleteNode(ctx, nodeUrl).
		Execute()
	if err != nil {
		return errors.Wrap(err, "error deleting node")
	}

	return nil
}
