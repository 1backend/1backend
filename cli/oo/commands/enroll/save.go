package enroll

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

		path := args[0]

		stat, err := os.Stat(path)
		if os.IsNotExist(err) {
			return errors.Wrap(err, fmt.Sprintf("path not found: '%v'", path))
		} else if err != nil {
			return errors.Wrap(err, "error checking path")
		}

		fileCount := 0
		if stat.IsDir() {
			// Handle directory: Iterate over files and collect enrolls
			err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("error accessing file '%v'", filePath))
				}
				if !info.IsDir() {
					// Collect enrolls from each file in the directory
					fileCount++
					var fileEnrolls []openapi.UserSvcEnrollInput
					err = util.ExtractFromFile(filePath, &fileEnrolls)
					if err != nil {
						return err
					}
					enrolls = append(enrolls, fileEnrolls...)
				}
				return nil
			})
			if err != nil {
				return err
			}
		} else {
			// Handle single file
			fileCount++
			var fileEnrolls []openapi.UserSvcEnrollInput
			err = util.ExtractFromFile(path, &fileEnrolls)
			if err != nil {
				return err
			}
			enrolls = append(enrolls, fileEnrolls...)
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
