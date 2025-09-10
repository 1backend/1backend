/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package permit

// import (
// 	"fmt"
//
// 	"github.com/1backend/1backend/cli/oo/config"
// 	openapi "github.com/1backend/1backend/clients/go"
// 	sdk "github.com/1backend/1backend/sdk/go"
// 	"github.com/spf13/cobra"
// )
//
// // Remove [permitId1] [permitId2]
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
// 	cf := client.NewApiClientFactory(url)
//
// 	_, _, err = cf.Client(client.WithToken(token)).
// 		UserSvcAPI.RemovePermit(ctx).
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
