package enroll

import (
	"io/ioutil"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save [role] [--userId | --contactId]  | [filePath | dirPath]
func Save(cmd *cobra.Command, args []string, userId, contactId string) error {
	ctx := cmd.Context()
	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	var enrolls []openapi.UserSvcEnrollInput

	if userId != "" || contactId != "" {
		enroll := openapi.UserSvcEnrollInput{
			Role: args[0],
		}
		if userId != "" {
			enroll.UserId = openapi.PtrString(userId)
		}
		if contactId != "" {
			enroll.ContactId = openapi.PtrString(contactId)
		}

		enrolls = append(enrolls, enroll)
	} else {
		enrolls, err = util.CollectFromPath[openapi.UserSvcEnrollInput](args[0], "enrolls")
		if err != nil {
			return err
		}
	}

	_, hrsp, err := cf.Client(client.WithToken(token)).
		UserSvcAPI.SaveEnrolls(ctx).
		Body(openapi.UserSvcSaveEnrollsRequest{
			Enrolls: enrolls,
		}).
		Execute()
	if err != nil {
		body, _ := ioutil.ReadAll(hrsp.Body)
		return errors.Wrap(err, "failed to save enrolls: "+string(body))
	}

	return nil
}
