package grant

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/openorch/openorch/cli/oo/config"
	"github.com/openorch/openorch/cli/oo/util"
	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save [filePath | dirPath]
func Save(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := sdk.NewApiClientFactory(url)

	path := args[0]

	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.Wrap(err, fmt.Sprintf("path not found: '%v'", path))
	} else if err != nil {
		return errors.Wrap(err, "error checking path")
	}

	var grants []openapi.UserSvcGrant

	fileCount := 0
	if stat.IsDir() {
		// Handle directory: Iterate over files and collect grants
		err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error accessing file '%v'", filePath))
			}
			if !info.IsDir() {
				// Collect grants from each file in the directory
				fileCount++
				var fileGrants []openapi.UserSvcGrant
				err = util.ExtractFromFile(filePath, &fileGrants)
				if err != nil {
					return err
				}
				grants = append(grants, fileGrants...)
			}
			return nil
		})
		if err != nil {
			return err
		}
	} else {
		// Handle single file
		fileCount++
		var fileGrants []openapi.UserSvcGrant
		err = util.ExtractFromFile(path, &fileGrants)
		if err != nil {
			return err
		}
		grants = append(grants, fileGrants...)
	}

	// Make a single API call to save all grants
	_, _, err = cf.Client(sdk.WithToken(token)).
		UserSvcAPI.SaveGrants(ctx).
		Body(openapi.UserSvcSaveGrantsRequest{
			Grants: grants,
		}).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to save grants")
	}

	return nil
}
