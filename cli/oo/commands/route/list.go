package route

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// List
func List(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)

	req := openapi.ProxySvcListRoutesRequest{}
	rsp, _, err := cf.Client(client.WithToken(token)).
		ProxySvcAPI.ListRoutes(ctx).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list routes")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"ROUTE ID\tTARGET",
	)

	for _, route := range rsp.Routes {

		fmt.Fprintf(
			writer,
			"%s\t%s\n",
			route.Id,
			route.Target,
		)
	}

	return nil
}
