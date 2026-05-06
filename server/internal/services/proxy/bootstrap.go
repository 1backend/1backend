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

	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
)

func (cs *ProxyService) BootstrapSaveRoutes(ctx context.Context, routes []proxy.RouteInput) error {
	if len(routes) == 0 {
		return nil
	}
	_, err := cs.saveRoutes(&proxy.SaveRoutesRequest{
		Routes: routes,
	})
	return err
}

func (cs *ProxyService) BootstrapSaveRedirects(ctx context.Context, redirects []proxy.RedirectInput) error {
	if len(redirects) == 0 {
		return nil
	}
	_, err := cs.saveRedirects(&proxy.SaveRedirectsRequest{
		Redirects: redirects,
	})
	return err
}
