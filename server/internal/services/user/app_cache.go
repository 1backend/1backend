/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

type appLookupResult struct {
	app   *user.App
	found bool
}

func cloneApp(app *user.App) *user.App {
	if app == nil {
		return nil
	}
	return &user.App{
		Id:   app.Id,
		Host: app.Host,
	}
}

func (s *UserService) cacheApp(app *user.App) {
	app = cloneApp(app)
	if app == nil || app.Id == "" {
		return
	}

	s.appCacheByID.Store(app.Id, app)
	if app.Host != "" {
		s.appCacheByHost.Store(app.Host, app)
	}
}

func (s *UserService) invalidateAppCache(app *user.App) {
	if app == nil {
		return
	}
	if app.Id != "" {
		s.appCacheByID.Delete(app.Id)
	}
	if app.Host != "" {
		s.appCacheByHost.Delete(app.Host)
	}
}

func (s *UserService) cachedAppByID(id string) (*user.App, bool) {
	if id == "" {
		return nil, false
	}

	val, found := s.appCacheByID.Load(id)
	if !found {
		return nil, false
	}

	app, ok := val.(*user.App)
	if !ok {
		s.appCacheByID.Delete(id)
		return nil, false
	}

	return cloneApp(app), true
}

func (s *UserService) cachedAppByHost(host string) (*user.App, bool) {
	if host == "" {
		return nil, false
	}

	val, found := s.appCacheByHost.Load(host)
	if !found {
		return nil, false
	}

	app, ok := val.(*user.App)
	if !ok {
		s.appCacheByHost.Delete(host)
		return nil, false
	}

	return cloneApp(app), true
}

func (s *UserService) appByID(id string) (*user.App, bool, error) {
	if app, found := s.cachedAppByID(id); found {
		return app, true, nil
	}

	val, err, _ := s.appLookupGroup.Do("id:"+id, func() (interface{}, error) {
		if app, found := s.cachedAppByID(id); found {
			return appLookupResult{app: app, found: true}, nil
		}

		appI, found, err := s.appStore.Query(
			datastore.Equals(datastore.Field("id"), id),
		).FindOne()
		if err != nil {
			return nil, err
		}
		if !found {
			return appLookupResult{}, nil
		}

		app, ok := appI.(*user.App)
		if !ok {
			return nil, errors.Errorf("expected app type, got %T", appI)
		}

		s.cacheApp(app)
		return appLookupResult{app: cloneApp(app), found: true}, nil
	})
	if err != nil {
		return nil, false, err
	}

	res := val.(appLookupResult)
	return cloneApp(res.app), res.found, nil
}

func (s *UserService) appByHost(host string) (*user.App, bool, error) {
	if app, found := s.cachedAppByHost(host); found {
		return app, true, nil
	}

	val, err, _ := s.appLookupGroup.Do("host:"+host, func() (interface{}, error) {
		if app, found := s.cachedAppByHost(host); found {
			return appLookupResult{app: app, found: true}, nil
		}

		appI, found, err := s.appStore.Query(
			datastore.Equals(datastore.Field("host"), host),
		).FindOne()
		if err != nil {
			return nil, err
		}
		if !found {
			return appLookupResult{}, nil
		}

		app, ok := appI.(*user.App)
		if !ok {
			return nil, errors.Errorf("expected app type, got %T", appI)
		}

		s.cacheApp(app)
		return appLookupResult{app: cloneApp(app), found: true}, nil
	})
	if err != nil {
		return nil, false, err
	}

	res := val.(appLookupResult)
	return cloneApp(res.app), res.found, nil
}
