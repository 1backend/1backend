package invite

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/1backend/1backend/cli/oo/config"
	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save [role] [--userId | --contactId]  | [filePath | dirPath]
func Save(cmd *cobra.Command, args []string, userId, contactId string) error {
	ctx := cmd.Context()
	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	var invites []openapi.UserSvcInviteInput

	if userId != "" || contactId != "" {
		invite := openapi.UserSvcInviteInput{
			Role: args[0],
		}
		if userId != "" {
			invite.UserId = openapi.PtrString(userId)
		}
		if contactId != "" {
			invite.ContactId = openapi.PtrString(contactId)
		}

		invites = append(invites, invite)
	} else {

		path := args[0]

		stat, err := os.Stat(path)
		if os.IsNotExist(err) {
			return errors.Wrap(err, fmt.Sprintf("path not found: '%v'", path))
		} else if err != nil {
			return errors.Wrap(err, "error checking path")
		}

		fileCount := 0
		if stat.IsDir() {
			// Handle directory: Iterate over files and collect invites
			err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("error accessing file '%v'", filePath))
				}
				if !info.IsDir() {
					// Collect invites from each file in the directory
					fileCount++
					var fileInvites []openapi.UserSvcInviteInput
					err = util.ExtractFromFile(filePath, &fileInvites)
					if err != nil {
						return err
					}
					invites = append(invites, fileInvites...)
				}
				return nil
			})
			if err != nil {
				return err
			}
		} else {
			// Handle single file
			fileCount++
			var fileInvites []openapi.UserSvcInviteInput
			err = util.ExtractFromFile(path, &fileInvites)
			if err != nil {
				return err
			}
			invites = append(invites, fileInvites...)
		}
	}

	// Make a single API call to save all invites
	_, hrsp, err := cf.Client(client.WithToken(token)).
		UserSvcAPI.SaveInvites(ctx).
		Body(openapi.UserSvcSaveInvitesRequest{
			Invites: invites,
		}).
		Execute()
	if err != nil {
		body, _ := ioutil.ReadAll(hrsp.Body)
		return errors.Wrap(err, "failed to save invites: "+string(body))
	}

	return nil
}
