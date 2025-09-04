/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package permit

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save [filePath | dirPath]
func Save(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	path := args[0]

	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.Wrap(err, fmt.Sprintf("path not found: '%v'", path))
	} else if err != nil {
		return errors.Wrap(err, "error checking path")
	}

	var permits []openapi.UserSvcPermitInput

	fileCount := 0
	if stat.IsDir() {
		// Handle directory: Iterate over files and collect permits
		err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error accessing file '%v'", filePath))
			}
			if !info.IsDir() {
				// Collect permits from each file in the directory
				fileCount++
				var filePermits []openapi.UserSvcPermitInput
				err = util.ExtractFromFile(filePath, &filePermits)
				if err != nil {
					return err
				}
				permits = append(permits, filePermits...)
			}
			return nil
		})
		if err != nil {
			return err
		}
	} else {
		// Handle single file
		fileCount++
		var filePermits []openapi.UserSvcPermitInput
		err = util.ExtractFromFile(path, &filePermits)
		if err != nil {
			return err
		}
		permits = append(permits, filePermits...)
	}

	// Make a single API call to save all permits
	_, _, err = cf.Client(client.WithToken(token)).
		UserSvcAPI.SavePermits(ctx).
		Body(openapi.UserSvcSavePermitsRequest{
			Permits: permits,
		}).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to save permits")
	}

	return nil
}
