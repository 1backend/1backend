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

const redirectSnapshotCacheKey = "\x00redirect-snapshot"

type redirectSnapshot struct {
	redirects map[string]proxy.Redirect
}

func cloneRedirect(redirect *proxy.Redirect) *proxy.Redirect {
	if redirect == nil {
		return nil
	}
	return &proxy.Redirect{
		Id:         redirect.Id,
		Target:     redirect.Target,
		StatusCode: redirect.StatusCode,
	}
}

func (rs *redirectSnapshot) redirect(id string) (*proxy.Redirect, bool) {
	if rs == nil {
		return nil, false
	}
	redirect, found := rs.redirects[id]
	if !found {
		return nil, false
	}
	return cloneRedirect(&redirect), true
}

func (cs *ProxyService) cachedRedirectSnapshot() (*redirectSnapshot, error) {
	if val, found := cs.redirectCache.Load(redirectSnapshotCacheKey); found {
		snapshot, ok := val.(*redirectSnapshot)
		if ok {
			return snapshot, nil
		}
		cs.redirectCache.Delete(redirectSnapshotCacheKey)
	}

	val, err, _ := cs.sf.Do(redirectSnapshotCacheKey, func() (interface{}, error) {
		if val, found := cs.redirectCache.Load(redirectSnapshotCacheKey); found {
			snapshot, ok := val.(*redirectSnapshot)
			if ok {
				return snapshot, nil
			}
			cs.redirectCache.Delete(redirectSnapshotCacheKey)
		}

		redirectIs, err := cs.redirectStore.Query().Find()
		if err != nil {
			return nil, errors.Wrap(err, "failed to query redirects")
		}

		redirects := make(map[string]proxy.Redirect, len(redirectIs))
		for _, redirectI := range redirectIs {
			redirect, ok := redirectI.(*proxy.Redirect)
			if !ok {
				return nil, errors.Errorf("expected redirect type, got %T", redirectI)
			}
			redirects[redirect.Id] = *redirect
		}

		snapshot := &redirectSnapshot{
			redirects: redirects,
		}
		cs.redirectCache.Store(redirectSnapshotCacheKey, snapshot)
		return snapshot, nil
	})
	if err != nil {
		return nil, err
	}

	return val.(*redirectSnapshot), nil
}

func (cs *ProxyService) cachedRedirects(ids []string) ([]proxy.Redirect, error) {
	snapshot, err := cs.cachedRedirectSnapshot()
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		redirects := make([]proxy.Redirect, 0, len(snapshot.redirects))
		for _, redirect := range snapshot.redirects {
			redirects = append(redirects, redirect)
		}
		return redirects, nil
	}

	redirects := make([]proxy.Redirect, 0, len(ids))
	seen := map[string]bool{}
	for _, id := range ids {
		if seen[id] {
			continue
		}
		seen[id] = true

		redirect, found := snapshot.redirect(id)
		if !found {
			continue
		}
		redirects = append(redirects, *redirect)
	}

	return redirects, nil
}

func (cs *ProxyService) invalidateRedirectCache() {
	// @todo this is not going to work in a distributed setting
	cs.redirectCache.Clear()
}
