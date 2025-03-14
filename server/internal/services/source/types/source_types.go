/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package source_svc

type ErrorResponse struct {
	Error string `json:"error"`
}

type Event struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}

type CheckoutRepoRequest struct {
	URL       string `json:"url"`         // Full repository URL (e.g., https://github.com/user/repo)
	Version   string `json:"version"`     // Branch, tag, or commit SHA
	Token     string `json:"token"`       // Token for HTTPS auth (optional for SSH)
	Username  string `json:"username"`    // Username for HTTPS or SSH user (optional for SSH)
	Password  string `json:"password"`    // Password or token for HTTPS auth
	SSHKey    string `json:"ssh_key"`     // SSH private key (optional for SSH connection)
	SSHKeyPwd string `json:"ssh_key_pwd"` // Password for SSH private key if encrypted (optional)
}

type CheckoutRepoResponse struct {
	Dir string `json:"dir"` // Directory where the repository was checked out
}
