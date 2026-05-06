package redirect

import (
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/1backend/1backend/cli/oo/util"
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

	ids, err := cmd.Flags().GetStringArray("ids")
	if err != nil {
		return errors.Wrap(err, "cannot read ids flag")
	}

	rsp := &ListRedirectsResponse{}
	err = doJSON(ctx, url, token, http.MethodPost, "/proxy-svc/redirects", ListRedirectsRequest{
		Ids: ids,
	}, rsp)
	if err != nil {
		return errors.Wrap(err, "failed to list redirects")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"REDIRECT ID\tTARGET\tSTATUS CODE",
	)

	for _, redirect := range rsp.Redirects {
		fmt.Fprintf(
			writer,
			"%s\t%s\t%d\n",
			redirect.Id,
			redirect.Target,
			redirect.StatusCode,
		)
	}

	return nil
}
