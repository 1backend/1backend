/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package node_types

import sdk "github.com/1backend/1backend/sdk/go"

type Options struct {
	Port        int
	GpuPlatform string

	// 1Backend Server Address
	// Crucial for distributed features.
	// Please see the documentation for the envar OB_URL
	Address string

	Az         string
	Region     string
	LLMHost    string
	VolumeName string
	ConfigPath string

	// eg. mysql, postgres
	Db string

	// Connection string eg.
	// "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
	DbConnectionString string

	// Crucial for distributed features.
	// Please see the documentation for the envar OB_NODE_ID
	NodeId string

	// DbPrefix allows us to have isolated envs for different test cases
	// but still make multiple nodes in those test cases use the same
	// shard of the db.
	DbPrefix string

	SourceControlToken  string
	SecretEncryptionKey string

	// so ugly, only temporary
	ClientFactory sdk.ClientFactory

	Test bool
}
