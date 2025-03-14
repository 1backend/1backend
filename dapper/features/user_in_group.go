/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package features

import dt "github.com/openorch/openorch/dapper/types"

var UserInGroup = dt.Feature{
	ID:   "user-in-group",
	Name: "User in Group",
	Arguments: []dt.Argument{
		{
			Name:    "username",
			Type:    dt.String,
			Default: "",
		},
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
# Add the user to the group if not already a member
if ! groups {{.username}} | grep &>/dev/null "\b{{.groupname}}\b"; then
	sudo usermod -aG {{.groupname}} {{.username}}
	echo "User {{.username}} added to the {{.groupname}} group"
	# Use newgrp to apply the new group membership immediately
	sudo -u "{{.username}}" newgrp {{.groupname}}
else
	echo "Cannot add user {{.username}}: already in the {{.groupname}} group"
fi
`,
				Runtime: dt.Bash,
				Sudo:    true,
				Reboot:  true,
			},
			Check: &dt.Script{
				Source: `
# Check if the user is in the {{.groupname}} group
if groups {{.username}} | grep &>/dev/null "\b{{.groupname}}\b"; then
	echo "User {{.username}} is in the {{.groupname}} group"
	exit 0
else
	echo "User {{.username}} is not in the {{.groupname}} group"
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
