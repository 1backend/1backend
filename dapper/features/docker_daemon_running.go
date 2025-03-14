/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package features

import dt "github.com/openorch/openorch/dapper/types"

const nameOfDistro = "dind"

var DockerDaemonRunning = dt.Feature{
	ID:   "docker-daemon-running",
	Name: "Docker Daemon Running",
	Arguments: []dt.Argument{
		{
			Name:    "username",
			Type:    dt.String,
			Default: "",
		},
		{
			Name:    "assetfolder",
			Type:    dt.String,
			Default: "",
		},
	},
	PlatformScripts: map[dt.Platform]*dt.Scripts{
		dt.Windows: {
			Check: &dt.Script{
				Source: `
# Function to check if the Docker daemon is running in WSL
function Check-DockerDaemon {
	$dockerStatus = wsl -d ` + nameOfDistro + ` -e /bin/ash -c "ps aux | grep dockerd | grep -v grep"
	if ($null -ne $dockerStatus -and $dockerStatus.Contains("dockerd")) {
		return $true
	} else {
		return $false
	}
}

try {
	if (Check-DockerDaemon) {
		Write-Host "Docker daemon is running"
		exit 0
	} else {
		Write-Host "Docker daemon is not running"
		exit 1
	}
} catch {
	Write-Error $_.Exception.Message
	exit 1
}
`,
				Runtime: dt.Powershell,
			},
			Execute: &dt.Script{
				// Since Detach: true is introduced it might be enough to do simple
				// wsl -d dind -e /bin/ash -c "dockerd --iptables=false --host=tcp://0.0.0.0:2375"
				// too but because of the redirects I'll leave it like this
				Source: `
Start-Process wsl -ArgumentList '-d', 'dind', '-e', '/bin/ash', '-c', '"dockerd --iptables=false --host=tcp://0.0.0.0:2375"' -NoNewWindow -RedirectStandardOutput "{{.assetfolder}}\dockerstd.log" -RedirectStandardError "{{.assetfolder}}\dockererr.log"
`,
				Runtime: dt.Powershell,
				// WSL distro kills itself if no windows process is attached to it
				// https://github.com/microsoft/WSL/issues/8854#issuecomment-1255501711
				Detach: true,
			},
		},
		dt.Linux: {
			Execute: &dt.Script{
				Source: `
if command -v systemctl &> /dev/null; then
	sudo systemctl start docker
elif command -v service &> /dev/null; then
	sudo service docker start
else
	echo "No known service management tool found to start Docker."
	exit 1
fi
`,
				Runtime: dt.Bash,
				Sudo:    true,
			},
			Check: &dt.Script{
				Source: `
if ! command -v docker &> /dev/null; then
	echo "Docker could not be found"
	exit 1
else
	if ! runuser -g "{{.username}}" -u "{{.username}}" -- bash -c 'docker ps' &> /dev/null; then
		echo "Docker is installed but the daemon is not running"
		exit 1
	else
		echo "Docker is installed and the daemon is running"
		exit 0
	fi
fi
`,
				Runtime: dt.Bash,
			},
		},
	},
	PlatformFeatures: map[dt.Platform][]any{
		dt.Windows: {
			// map[string]any{
			// 	"featureId": WslSetDefaultVersion.ID,
			// 	"args": []any{
			// 		"2",
			// 	},
			// },
			map[string]any{
				"featureId": FileDownloaded.ID,
				"args": []any{
					"https://singulatron.com/downloads/docker-latest.tar",
					"{{.assetfolder}}",
					"docker-latest.tar",
				},
			},
			map[string]any{
				"featureId": WSLTarLoaded.ID,
				"args": []any{
					"{{.assetfolder}}\\docker-latest.tar",
					nameOfDistro,
				},
			},
		},
		dt.Linux: {
			map[string]any{
				"featureId": DockerDaemonInstalled.ID,
			},
			map[string]any{
				"featureId": GroupExists.ID,
				"args":      []any{"docker"},
			},
			map[string]any{
				"featureId": UserInGroup.ID,
				"args":      []any{"{{.username}}", "docker"},
			},
		},
	},
}
