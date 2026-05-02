/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package telemetry

import (
	"context"
	"sync"

	"github.com/1backend/1backend/sdk/go/datastore"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	autoIndexMu     sync.RWMutex
	autoIndexStores = map[string]autoIndexStore{}
)

type autoIndexStore struct {
	backend  string
	table    string
	provider datastore.AutoIndexStatsProvider
}

// RegisterAutoIndexStatsProvider exposes automatic-index state for Prometheus
// collection. The provider is normally registered by InstrumentDataStore.
func RegisterAutoIndexStatsProvider(backend, table string, provider any) {
	statsProvider, ok := provider.(datastore.AutoIndexStatsProvider)
	if !ok || statsProvider == nil {
		return
	}

	ensureInstruments()

	backend = normalizeBackend(backend)
	key := backend + "/" + table

	autoIndexMu.Lock()
	defer autoIndexMu.Unlock()
	autoIndexStores[key] = autoIndexStore{
		backend:  backend,
		table:    table,
		provider: statsProvider,
	}
}

func observeAutoIndexes(_ context.Context, observer metric.Observer) error {
	autoIndexMu.RLock()
	stores := make([]autoIndexStore, 0, len(autoIndexStores))
	for _, store := range autoIndexStores {
		stores = append(stores, store)
	}
	autoIndexMu.RUnlock()

	for _, store := range stores {
		stats := store.provider.AutoIndexStats()
		baseAttrs := []attribute.KeyValue{
			attribute.String("datastore.backend", normalizeBackend(store.backend)),
			attribute.String("db.collection.name", store.table),
			attribute.String("datastore.autoindex.backend", stats.Backend),
		}

		supported := int64(0)
		if stats.Supported {
			supported = 1
		}
		observer.ObserveInt64(autoIndexSupported, supported, metric.WithAttributes(baseAttrs...))

		for _, shape := range stats.Shapes {
			attrs := append([]attribute.KeyValue{}, baseAttrs...)
			attrs = append(attrs,
				attribute.String("datastore.query.fingerprint", shape.Fingerprint),
				attribute.Bool("datastore.autoindex.eligible", shape.Eligible),
				attribute.String("datastore.autoindex.reason", shape.Reason),
			)
			observer.ObserveInt64(autoIndexShapes, int64(shape.Hits), metric.WithAttributes(attrs...))
		}

		for _, index := range stats.Indexes {
			attrs := append([]attribute.KeyValue{}, baseAttrs...)
			attrs = append(attrs,
				attribute.String("datastore.index.fingerprint", index.Fingerprint),
				attribute.String("datastore.index.kind", string(index.Kind)),
				attribute.String("datastore.index.status", string(index.Status)),
				attribute.String("datastore.index.method", index.Method),
				attribute.String("datastore.index.name", index.Name),
			)
			observer.ObserveInt64(autoIndexInfo, 1, metric.WithAttributes(attrs...))
		}
	}

	return nil
}
