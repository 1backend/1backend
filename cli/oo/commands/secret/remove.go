package secret

import (
	"fmt"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Remove --key=key1 --id=id1 --id=id2
func Remove(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// Retrieve flags for keys and ids
	keys, _ := cmd.Flags().GetStringArray("key")
	ids, _ := cmd.Flags().GetStringArray("id")

	// Ensure that at least one key or id is specified
	if len(keys) == 0 && len(ids) == 0 {
		return fmt.Errorf("at least one --key or --id must be specified")
	}

	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env URL")
	}

	cf := client.NewApiClientFactory(url)

	_, _, err = cf.Client(client.WithToken(token)).
		SecretSvcAPI.RemoveSecrets(ctx).
		Body(openapi.SecretSvcRemoveSecretsRequest{
			Keys: keys,
			Ids:  ids,
		}).
		Execute()
	if err != nil {
		return errors.Wrap(err, "error deleting secrets")
	}

	return nil
}
