/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxyservice

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"log/slog"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"
	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"

	"github.com/pkg/errors"
)

//
// We log errors in these methods because we're not sure how autocert logs the errors.
//

func (cs *ProxyService) Get(ctx context.Context, key string) ([]byte, error) {
	var err error

	defer func() {
		if err != nil {
			logger.Error(
				"Failed to get cert from store",
				slog.String("key", key),
				slog.Any("error", err),
			)
		}
	}()

	certI, found, err := cs.certStore.Query(
		datastore.Id(key),
	).FindOne()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query cert store")
	}

	if !found {
		return nil, errors.Errorf("cert not found for key '%s'", key)
	}

	cert, ok := certI.(*proxy.Cert)
	if !ok {
		return nil, errors.Errorf("expected cert type, got %T", certI)
	}

	decoder := base64.NewDecoder(base64.StdEncoding, bytes.NewReader([]byte(cert.Cert)))
	data, err := io.ReadAll(decoder)
	if err != nil && err != io.EOF {

		return nil, errors.Wrap(err, "failed to decode cert data")
	}

	if len(data) == 0 {
		return nil, errors.Errorf("cert data is empty for key '%s'", key)
	}

	return data, nil
}

func (cs *ProxyService) Put(ctx context.Context, key string, data []byte) error {
	var err error

	defer func() {
		if err != nil {
			logger.Error(
				"Failed to put cert to store",
				slog.String("key", key),
				slog.Any("error", err),
			)
		}
	}()

	encoded := base64.StdEncoding.EncodeToString(data)

	cert := &proxy.Cert{
		Id:   key,
		Cert: encoded,
	}

	err = cs.certStore.Upsert(cert)
	if err != nil {
		logger.Error(
			"Failed to upsert cert in store",
			slog.String("key", key),
			slog.Any("error", err),
		)
		return errors.Wrap(err, "failed to upsert cert in store")
	}

	return nil
}

func (cs *ProxyService) Delete(ctx context.Context, key string) error {
	var err error

	defer func() {
		if err != nil {
			logger.Error(
				"Failed to delete cert from store",
				slog.String("key", key),
				slog.Any("error", err),
			)
		}
	}()

	err = cs.certStore.Query(datastore.Id(key)).Delete()
	if err != nil {
		return errors.Wrap(err, "failed to delete cert from store")
	}

	return nil
}
