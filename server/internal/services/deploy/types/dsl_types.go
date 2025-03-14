/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package deploy_svc

type CommandType string

const (
	CommandTypeStart CommandType = "START"
	CommandTypeKill  CommandType = "KILL"
	CommandTypeScale CommandType = "SCALE"
)

type Command struct {
	Action       CommandType
	NodeId       string
	NodeUrl      string
	DeploymentId *string
	InstanceId   *string
	InstanceUrl  *string
}
