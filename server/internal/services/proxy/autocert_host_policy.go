/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxyservice

import (
	"context"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/pkg/errors"
)

func (cs *ProxyService) HostPolicy(ctx context.Context, host string) error {
	_, found, err := cs.routeStore.Query(
		datastore.Id(host),
	).FindOne()
	if err != nil {
		return errors.Wrap(err, "failed to query route by host")
	}

	if !found {
		return errors.New("host not found in routes")
	}

	return nil
}
