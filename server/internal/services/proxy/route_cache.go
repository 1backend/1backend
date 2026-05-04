/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxyservice

import (
	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
	"github.com/pkg/errors"
)

const routeSnapshotCacheKey = "\x00route-snapshot"

type routeSnapshot struct {
	routes map[string]proxy.Route
}

func cloneRoute(route *proxy.Route) *proxy.Route {
	if route == nil {
		return nil
	}
	return &proxy.Route{
		Id:     route.Id,
		Target: route.Target,
	}
}

func (rs *routeSnapshot) route(id string) (*proxy.Route, bool) {
	if rs == nil {
		return nil, false
	}
	route, found := rs.routes[id]
	if !found {
		return nil, false
	}
	return cloneRoute(&route), true
}

func (cs *ProxyService) cachedRouteSnapshot() (*routeSnapshot, error) {
	if val, found := cs.routeCache.Load(routeSnapshotCacheKey); found {
		snapshot, ok := val.(*routeSnapshot)
		if ok {
			return snapshot, nil
		}
		cs.routeCache.Delete(routeSnapshotCacheKey)
	}

	val, err, _ := cs.sf.Do(routeSnapshotCacheKey, func() (interface{}, error) {
		if val, found := cs.routeCache.Load(routeSnapshotCacheKey); found {
			snapshot, ok := val.(*routeSnapshot)
			if ok {
				return snapshot, nil
			}
			cs.routeCache.Delete(routeSnapshotCacheKey)
		}

		routeIs, err := cs.routeStore.Query().Find()
		if err != nil {
			return nil, errors.Wrap(err, "failed to query routes")
		}

		routes := make(map[string]proxy.Route, len(routeIs))
		for _, routeI := range routeIs {
			route, ok := routeI.(*proxy.Route)
			if !ok {
				return nil, errors.Errorf("expected route type, got %T", routeI)
			}
			routes[route.Id] = *route
		}

		snapshot := &routeSnapshot{
			routes: routes,
		}
		cs.routeCache.Store(routeSnapshotCacheKey, snapshot)
		return snapshot, nil
	})
	if err != nil {
		return nil, err
	}

	return val.(*routeSnapshot), nil
}

func (cs *ProxyService) cachedRouteByID(id string) (*proxy.Route, bool, error) {
	snapshot, err := cs.cachedRouteSnapshot()
	if err != nil {
		return nil, false, err
	}
	route, found := snapshot.route(id)
	return route, found, nil
}

func (cs *ProxyService) cachedRoutes(ids []string) ([]proxy.Route, error) {
	snapshot, err := cs.cachedRouteSnapshot()
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		routes := make([]proxy.Route, 0, len(snapshot.routes))
		for _, route := range snapshot.routes {
			routes = append(routes, route)
		}
		return routes, nil
	}

	routes := make([]proxy.Route, 0, len(ids))
	seen := map[string]bool{}
	for _, id := range ids {
		if seen[id] {
			continue
		}
		seen[id] = true

		route, found := snapshot.route(id)
		if !found {
			continue
		}
		routes = append(routes, *route)
	}

	return routes, nil
}

func (cs *ProxyService) invalidateRouteCache() {
	// @todo this is not going to work in a distributed setting
	cs.routeCache.Clear()
}
