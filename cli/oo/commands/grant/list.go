package grant

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/openorch/openorch/cli/oo/config"
	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
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

	cf := sdk.NewApiClientFactory(url)

	req := openapi.UserSvcListGrantsRequest{}

	rsp, _, err := cf.Client(sdk.WithToken(token)).
		UserSvcAPI.ListGrants(ctx).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list grants")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"GRANT ID\tPERMISSION ID\tSLUGS",
	)

	for _, grant := range rsp.Grants {
		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\n",
			*grant.Id,
			*grant.PermissionId,
			strings.Join(grant.Slugs, ", "),
		)
	}

	return nil
}
