/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package features

import (
	dt "github.com/openorch/openorch/dapper/types"
)

var WslEnabled = dt.Feature{
	ID:   "wsl-enabled",
	Name: "WSL Enabled",
	PlatformScripts: map[dt.Platform]*dt.Scripts{
		dt.Windows: {
			Execute: &dt.Script{
				Source:  `dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart`,
				Runtime: "powershell",
			},
			Check: &dt.Script{
				Source: `
$wslFeature = dism.exe /online /get-featureinfo /featurename:Microsoft-Windows-Subsystem-Linux
if ($wslFeature.Contains("State : Enabled")) {
    return $true
} else {
    return $false
}`,
				Runtime: "powershell",
			},
		},
	},
	PlatformFeatures: map[dt.Platform][]any{
		dt.Windows: {VirtualMachinePlatformFeature.ID},
	},
}
