package enroll

import (
	"fmt"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Remove deletes enrolls by ID.
func Remove(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if len(args) == 0 {
		return fmt.Errorf("at least one enroll ID must be specified")
	}

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)

	_, hrsp, err := cf.Client(client.WithToken(token)).
		UserSvcAPI.DeleteEnrolls(ctx).
		Body(openapi.UserSvcDeleteEnrollsRequest{
			Ids: args,
		}).
		Execute()
	if err != nil {
		return util.ErrorWithBody(err, hrsp, "failed to delete enrolls")
	}

	return nil
}
