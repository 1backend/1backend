package user

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

// List Users
func ListUsers(
	cmd *cobra.Command,
	args []string,
	userId string,
	contactId string,
	limit int64,
) error {
	ctx := cmd.Context()

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)

	req := openapi.UserSvcListUsersRequest{}
	if userId != "" {
		req.Ids = []string{userId}
	}
	if contactId != "" {
		req.ContactId = openapi.PtrString(contactId)
	}

	if limit != 0 {
		req.Limit = openapi.PtrInt32(int32(limit))
	}

	rsp, _, err := cf.Client(client.WithToken(token)).
		UserSvcAPI.ListUsers(ctx).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list secrets")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"USER ID\tSLUG\tNAME\tROLES IDS\tCONTACT IDS",
	)

	for _, user := range rsp.Users {
		name := ""
		if user.Name != nil {
			name = *user.Name
		}
		roleIds := strings.Join(user.Roles, ", ")
		contactIds := strings.Join(user.ContactIds, ", ")

		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%s\t%s\n",
			user.Id,
			user.Slug,
			name,
			roleIds,
			contactIds,
		)
	}

	return nil
}
