/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package features

import dt "github.com/openorch/openorch/dapper/types"

var GroupExists = dt.Feature{
	ID:   "group-exists",
	Name: "Group Exists",
	Arguments: []dt.Argument{
		{
			Name:    "groupname",
			Type:    dt.String,
			Default: "",
		},
	},
	PlatformScripts: map[dt.Platform]*dt.Scripts{
		dt.Linux: {
			Execute: &dt.Script{
				Source: `
# Check if the group exists, create if it doesn't
if ! getent group {{.groupname}} &>/dev/null; then
	sudo groupadd {{.groupname}}
	echo "Group {{.groupname}} created"
else
	echo "Group {{.groupname}} already exists"
fi
`,
				Runtime: dt.Bash,
				Sudo:    true,
			},
			Check: &dt.Script{
				Source: `
# Check if the group exists
if getent group {{.groupname}} &>/dev/null; then
	echo "Group {{.groupname}} exists"
	exit 0
else
	echo "Group {{.groupname}} does not exist"
	exit 1
fi
`,
				Runtime: dt.Bash,
				Sudo:    false,
			},
		},
	},
	PlatformFeatures: map[dt.Platform][]any{},
}
