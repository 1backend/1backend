package app

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

// List handles the `oo app list` command.
func List(cmd *cobra.Command, args []string, ids []string, hosts []string) error {
	ctx := cmd.Context()

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)

	req := openapi.UserSvcListAppsRequest{}
	if len(ids) > 0 {
		req.Ids = ids
	}
	if len(hosts) > 0 {
		req.Host = hosts
	}

	rsp, _, err := cf.Client(client.WithToken(token)).
		UserSvcAPI.ListApps(ctx).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list apps")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "APP ID\tHOST")

	for _, app := range rsp.Apps {
		fmt.Fprintf(writer, "%s\t%s\n", app.Id, app.Host)
	}

	return nil
}
