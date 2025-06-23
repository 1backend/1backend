package deployment

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/1backend/1backend/cli/oo/util"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// List
func List(cmd *cobra.Command, args []string, full bool) error {
	ctx := cmd.Context()

	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	rsp, _, err := cf.Client(client.WithToken(token)).
		DeploySvcAPI.ListDeployments(ctx).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to save service deployment")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "ID\tDEFINITION ID\tSTATUS\tDETAILS")

	for _, deployment := range rsp.Deployments {
		details := ""
		if deployment.Details != nil {
			details = *deployment.Details
		}
		charsToTrim := " .,"

		if len(details) > 75 && !full {
			details = strings.Trim(details[0:75], charsToTrim) + "…"
		}

		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%s\n",
			deployment.Id,
			deployment.DefinitionId,
			*deployment.Status,
			details,
		)
	}

	return nil
}
