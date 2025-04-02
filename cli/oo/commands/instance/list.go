package instance

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/1backend/1backend/cli/oo/config"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// List
func List(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	rsp, _, err := cf.Client(client.WithToken(token)).
		RegistrySvcAPI.ListInstances(ctx).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list instances")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "INSTANCE ID\tURL\tSTATUS\tOWNER SLUG\tLAST HEARTBEAT")

	for _, instance := range rsp.Instances {
		heartbeat := ""
		if instance.LastHeartbeat != nil {
			heartbeat = *instance.LastHeartbeat
			parsedTime, err := time.Parse(time.RFC3339Nano, heartbeat)
			if err == nil {
				heartbeat = time.Since(parsedTime).
					Truncate(time.Second).
					String() +
					" ago"
			}
		}

		ownerSlug := ""
		if instance.Slug != nil {
			ownerSlug = *instance.Slug
		}

		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%s\t%s\n",
			instance.Id,
			instance.Url,
			instance.Status,
			ownerSlug,
			heartbeat,
		)
	}

	return nil
}
