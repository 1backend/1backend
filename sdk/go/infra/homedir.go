/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package infra

import (
	"os"
	"path"

	"github.com/pkg/errors"
)

const onebackendFolder = ".1backend"

type HomeDirOptions struct {
	Test bool

	// The config folder name. Not the full path.
	// This mostly exists to support running multiple local
	// servers for testing purposes.
	ConfigFolder string
}

// HomeDir of the 1backend server
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
		if options.ConfigFolder != "" {
			homeDir = options.ConfigFolder
		} else {
			homeDir, err = os.UserHomeDir()
			if err != nil {
				return "", errors.Wrap(err, "homedir creation failed")
			}
			fold := onebackendFolder

			homeDir = path.Join(homeDir, fold)
		}

	}

	return homeDir, nil
}
