/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package features

import dt "github.com/openorch/openorch/dapper/types"

var WSLTarLoaded = dt.Feature{
	ID:   "wsl-tar-loaded",
	Name: "WSL Tar Loaded",
	Arguments: []dt.Argument{
		{
			Name:    "tarpath",
			Type:    dt.String,
			Default: "",
		},
		{
			Name:    "distroname",
			Type:    dt.String,
			Default: "",
		},
	},
	PlatformScripts: map[dt.Platform]*dt.Scripts{
		dt.Windows: {
			Execute: &dt.Script{
				Source: `
$distroname = "{{.distroname}}"
$tarpath = "{{.tarpath}}"
$wslPath = "C:\WSL\$distroname"

if (-Not (Test-Path -Path $wslPath)) {
    Write-Output "Path $wslPath does not exist. Creating it..."
    New-Item -ItemType Directory -Path $wslPath -Force
}

Write-Host "Importing $distroname."
wsl --import $distroname $wslPath $tarpath --version 2

# Set the distribution as default
# See https://github.com/microsoft/WSL/issues/5923
wsl -s $distroname
`,
				Runtime: dt.Powershell,
				Sudo:    false,
			},
			Check: &dt.Script{
				Source: `
$distroname = "{{.distroname}}"
$loadedDistros = wsl -l -q
if ($null -ne $loadedDistros -and $loadedDistros.Contains($distroname)) {
    Write-Host "$distroname is already loaded."
    exit 0
} else {
    Write-Host "$distroname is not loaded."
    exit 1
}
`,
				Runtime: dt.Powershell,
				Sudo:    false,
			},
		},
	},
	PlatformFeatures: map[dt.Platform][]any{
		dt.Windows: {
			map[string]any{
				"featureId": WslUpdated.ID,
				"args": []any{
					"2",
				},
			},
		},
	},
}
