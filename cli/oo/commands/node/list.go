/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package node

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/1backend/1backend/cli/oo/util"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// List
func List(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	rsp, _, err := cf.Client(client.WithToken(token)).
		RegistrySvcAPI.ListNodes(ctx).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list nodes")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "NODE ID\tURL\tLAST HEARTBEAT")

	for _, node := range rsp.Nodes {
		heartbeat := ""
		if node.LastHeartbeat != nil {
			heartbeat = *node.LastHeartbeat
			parsedTime, err := time.Parse(time.RFC3339Nano, heartbeat)
			if err == nil {
				heartbeat = time.Since(parsedTime).
					Truncate(time.Second).
					String() +
					" ago"
			}
		}

		fmt.Fprintf(writer, "%s\t%s\t%s\n", node.Id, node.Url, heartbeat)
	}

	return nil
}

func roundDuration(d time.Duration) string {
	// You can adjust the rounding logic as needed
	if d < time.Minute {
		// Round to nearest second
		return fmt.Sprintf("%.0fs", d.Seconds())
	} else if d < time.Hour {
		// Round to nearest minute
		return fmt.Sprintf("%dm", int(d.Minutes()))
	} else {
		// Round to nearest hour
		return fmt.Sprintf("%dh", int(d.Hours()))
	}
}
