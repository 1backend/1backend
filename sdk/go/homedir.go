/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package sdk

import (
	"os"
	"path"

	"github.com/pkg/errors"
)

const onebackendFolder = ".1backend"

type HomeDirOptions struct {
	Test bool
}

func HomeDir(options HomeDirOptions) (string, error) {
	var (
		homeDir string
		err     error
	)
	if options.Test {
		homeDir, err = os.MkdirTemp("", "1backend-")
		if err != nil {
			return "", errors.Wrap(err,
				"homedir creation failed",
			)
		}
	} else {
		homeDir, err = os.UserHomeDir()
		if err != nil {
			return "", errors.Wrap(err, "homedir creation failed")
		}
		homeDir = path.Join(homeDir, onebackendFolder)
	}

	return homeDir, nil
}
