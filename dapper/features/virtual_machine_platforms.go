/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package features

import dt "github.com/openorch/openorch/dapper/types"

var VirtualMachinePlatformFeature = dt.Feature{
	ID:   "virtual-machine-platform-enabled",
	Name: "Virtual Machine Platform Enabled",
	PlatformScripts: map[dt.Platform]*dt.Scripts{
		dt.Windows: {
			Execute: &dt.Script{
				Source:  `dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart`,
				Runtime: "powershell",
			},
			Check: &dt.Script{
				Source: `
try {
    $vmPlatformFeature = dism.exe /online /get-featureinfo /featurename:VirtualMachinePlatform
    if ($vmPlatformFeature.Contains("State : Enabled")) {
		exit 0
    } else {
        exit 1
    }
} catch {
    Write-Host "Failed to check or enable VirtualMachinePlatform: $_"
    exit 1
}`,
				Runtime: "powershell",
			},
			Uninstall: &dt.Script{
				Source: `
Write-Host "Disabling VirtualMachinePlatform feature"
dism.exe /online /disable-feature /featurename:VirtualMachinePlatform /norestart
`,
				Runtime: "powershell",
			},
		},
	},
	Features: nil,
}
