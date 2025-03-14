/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package features

import dt "github.com/openorch/openorch/dapper/types"

var DockerDaemonInstalled = dt.Feature{
	ID:   "docker-daemon-installed",
	Name: "Docker Daemon Installed",
	PlatformScripts: map[dt.Platform]*dt.Scripts{
		dt.Linux: {
			Execute: &dt.Script{
				Source: `
# Update the apt package index and install packages to allow apt to use a repository over HTTPS
sudo apt-get update
sudo apt-get install \
	apt-transport-https \
	ca-certificates \
	curl \
	gnupg-agent \
	software-properties-common -y

# Add Docker’s official GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# Verify that you now have the key with the fingerprint 
sudo apt-key fingerprint 0EBFCD88

# Set up the stable repository
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

# Update the apt package index, and install the latest version of Docker Engine and containerd
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io -y

# Add current user to the docker group so you can execute docker commands without sudo
sudo usermod -aG docker $USER

# Instructions to log out and back in for group changes to take effect
echo "Please log out and log back in so that your group membership is re-evaluated."
`,
				Runtime: dt.Bash,
				Sudo:    true,
			},

			Check: &dt.Script{
				Source: `
if ! command -v docker &> /dev/null
then
	echo "Docker could not be found"
	exit 1
else
	echo "Docker is installed"
	exit 0
fi
`,
				Runtime: dt.Bash,
			},
		},
	},
}
