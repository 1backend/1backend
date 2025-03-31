package user

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/1backend/1backend/cli/oo/config"
	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
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

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := sdk.NewApiClientFactory(url)

	req := openapi.UserSvcListUsersRequest{}
	if userId != "" {
		req.UserId = openapi.PtrString(userId)
	}
	if contactId != "" {
		req.ContactId = openapi.PtrString(contactId)
	}

	if req.Query == nil {
		req.Query = &openapi.DatastoreQuery{}
	}
	if limit != 0 {
		req.Query.Limit = openapi.PtrInt32(int32(limit))
	}

	rsp, _, err := cf.Client(sdk.WithToken(token)).
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
		roleIds := strings.Join(user.RoleIds, ", ")
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
