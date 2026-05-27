package enroll

import (
	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// PatchRole updates an existing enroll's role by ID.
func PatchRole(cmd *cobra.Command, args []string, appHost string) error {
	ctx := cmd.Context()
	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	enrollId := args[0]
	role := args[1]

	enroll := openapi.UserSvcEnrollInput{
		Id:   openapi.PtrString(enrollId),
		Role: role,
	}
	if appHost != "" {
		enroll.AppHost = openapi.PtrString(appHost)
	}

	cf := client.NewApiClientFactory(url)
	_, hrsp, err := cf.Client(client.WithToken(token)).
		UserSvcAPI.SaveEnrolls(ctx).
		Body(openapi.UserSvcSaveEnrollsRequest{
			Enrolls: []openapi.UserSvcEnrollInput{enroll},
		}).
		Execute()
	if err != nil {
		return util.ErrorWithBody(err, hrsp, "failed to patch enroll role")
	}

	return nil
}
