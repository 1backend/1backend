/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package permit

import (
	"fmt"
	"os"
	"strings"
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

	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	req := openapi.UserSvcListPermitsRequest{}

	rsp, _, err := cf.Client(client.WithToken(token)).
		UserSvcAPI.ListPermits(ctx).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list permits")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"PERMIT ID\tPERMISSION\tSLUGS\tROLES",
	)

	for _, permit := range rsp.Permits {
		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%s\n",
			permit.Id,
			permit.Permission,
			strings.Join(permit.Slugs, ", "),
			strings.Join(permit.Roles, ", "),
		)
	}

	return nil
}
