package invite

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/1backend/1backend/cli/oo/config"
	openapi "github.com/1backend/1backend/clients/go"
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

	req := openapi.UserSvcListInvitesRequest{}

	rsp, _, err := cf.Client(client.WithToken(token)).
		UserSvcAPI.ListInvites(ctx).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list invites")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"INVITE ID\tROLE\tUSER ID\tCONTACT ID",
	)

	for _, invite := range rsp.Invites {
		userId := ""
		if invite.UserId != nil {
			userId = *invite.UserId
		}

		contactId := ""
		if invite.ContactId != nil {
			contactId = *invite.ContactId
		}

		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%s\n",
			invite.Id,
			invite.Role,
			userId,
			contactId,
		)
	}

	return nil
}
