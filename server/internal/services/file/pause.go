/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package fileservice

import (
	"fmt"

	downloadtypes "github.com/openorch/openorch/server/internal/services/file/types"
	"github.com/pkg/errors"
)

// pause a download
func (ds *FileService) pause(url string) error {
	d, exists := ds.getDownload(url)
	if !exists {
		return fmt.Errorf("url '%v' is not being downloaded", url)
	}

	d.Status = downloadtypes.DownloadStatusPaused

	err := ds.downloadStore.Upsert(d)
	if err != nil {
		return errors.Wrap(err, "failed to upsert download")
	}

	return nil
}
