/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package boot

import "github.com/1backend/1backend/sdk/go/infra"

// DataStoreConfig returns the datastore settings derived from boot options and
// the usual OB_DB / OB_DB_CONNECTION_STRING / OB_DB_READ_CONNECTION_STRING
// environment variables.
func (o *Options) DataStoreConfig() infra.DataStoreConfig {
	return infra.DataStoreConfig{
		Test:                   o.Test,
		Db:                     o.Db,
		DbConnectionString:     o.DbConnectionString,
		ReadDbConnectionString: o.ReadDbConnectionString,
	}
}

// NewDataStoreFactory returns the SDK datastore factory for a service. The
// factory instruments stores with the SDK telemetry datastore wrapper.
func (o *Options) NewDataStoreFactory() (infra.DataStoreFactory, error) {
	if err := o.LoadEnvars(); err != nil {
		return nil, err
	}
	return infra.NewDataStoreFactory(o.DataStoreConfig())
}
