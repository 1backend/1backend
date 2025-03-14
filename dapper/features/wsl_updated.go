/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package features

import dt "github.com/openorch/openorch/dapper/types"

var WslUpdated = dt.Feature{
	ID:   "wsl-updated",
	Name: "WSL Updated",
	Arguments: []dt.Argument{
		{
			Name:    "wslVersion",
			Type:    dt.Int,
			Default: 2,
		},
	},
	PlatformScripts: map[dt.Platform]*dt.Scripts{
		dt.Windows: {
			Execute: &dt.Script{
				Source: `
Write-Host "Updating WSL kernel"
# https://github.com/microsoft/WSL/issues/9039#issuecomment-1286596985
wsl --update --web-download
Write-Host "Setting default WSL version to {{.wslVersion}}"
wsl --set-default-version {{.wslVersion}}
`,
				Runtime: "powershell",
			},
			Check: &dt.Script{
				Source: `
$wslStatusOutput = wsl --status
if ($null -ne $wslStatusOutput -and $wslStatusOutput.Contains("Default Version: {{.wslVersion}}")) {
    exit 0
} else {
	Write-Output "WSL is not updated: $wslStatusOutput"
    exit 1
}`,
				Runtime: "powershell",
			},
		},
	},
	PlatformFeatures: map[dt.Platform][]any{
		dt.Windows: {WslInstalled.ID},
	},
}
