package secret

import (
	"fmt"

	"github.com/openorch/openorch/cli/oo/config"
	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
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

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return fmt.Errorf("cannot get env URL: '%v'", err)
	}

	cf := sdk.NewApiClientFactory(url)

	_, _, err = cf.Client(sdk.WithToken(token)).
		SecretSvcAPI.RemoveSecrets(ctx).
		Body(openapi.SecretSvcRemoveSecretsRequest{
			Keys: keys,
			Ids:  ids,
		}).
		Execute()
	if err != nil {
		return fmt.Errorf("error deleting secrets: '%v'", err)
	}

	return nil
}
