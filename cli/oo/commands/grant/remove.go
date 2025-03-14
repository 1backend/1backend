package grant

// import (
// 	"fmt"
//
// 	"github.com/openorch/openorch/cli/oo/config"
// 	openapi "github.com/openorch/openorch/clients/go"
// 	sdk "github.com/openorch/openorch/sdk/go"
// 	"github.com/spf13/cobra"
// )
//
// // Remove [grantId1] [grantId2]
// func Remove(cmd *cobra.Command, args []string) error {
// 	ctx := cmd.Context()
//
// 	if len(args) == 0 {
// 		return fmt.Errorf("at least one ID must be specified")
// 	}
//
// 	url, token, err := config.GetSelectedUrlAndToken()
// 	if err != nil {
// 		return fmt.Errorf("cannot get environment URL: '%v'", err)
// 	}
//
// 	cf := sdk.NewApiClientFactory(url)
//
// 	_, _, err = cf.Client(sdk.WithToken(token)).
// 		UserSvcAPI.RemoveGrant(ctx).
// 		Body(openapi.SecretSvcRemoveSecretsRequest{
// 			Ids: args, // Use args directly as IDs
// 		}).
// 		Execute()
// 	if err != nil {
// 		return fmt.Errorf("error deleting secrets: '%v'", err)
// 	}
//
// 	fmt.Println("Secrets removed successfully")
// 	return nil
// }
