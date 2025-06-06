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
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io"
	"log/slog"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"
	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"

	"github.com/pkg/errors"
)

type CertStore struct {
	Db datastore.DataStore
}

//
// We log errors in these methods because we're not sure how autocert logs the errors.
//

func (cs *CertStore) Get(ctx context.Context, key string) ([]byte, error) {
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

	certI, found, err := cs.Db.Query(
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

func (cs *CertStore) Put(ctx context.Context, key string, data []byte) error {
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

	var existingCert proxy.Cert
	certI, existingCertFound, err := cs.Db.Query(
		datastore.Id(key),
	).FindOne()
	if err != nil {
		return errors.Wrap(err, "failed to query cert store")
	}
	if existingCertFound {
		c, ok := certI.(*proxy.Cert)
		if !ok {
			return errors.Errorf("expected cert type, got %T", certI)
		}
		existingCert = *c
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	cert := &proxy.Cert{
		Id:   key,
		Cert: encoded,
	}

	now := time.Now()

	if existingCertFound {
		cert.CreatedAt = existingCert.CreatedAt
		cert.UpdatedAt = now
	} else {
		cert.CreatedAt = now
		cert.UpdatedAt = now
	}

	info, err := parseCertInfo(data)
	if err != nil {
		logger.Error(
			"Failed to parse cert info",
			slog.String("key", key),
			slog.Any("error", err),
		)
		// Do not return, it's not critical if we can't parse the cert info.
	} else {
		cert.NotBefore = info.NotBefore
		cert.NotAfter = info.NotAfter
		cert.CommonName = info.CommonName
		cert.DNSNames = info.DNSNames
		cert.Issuer = info.Issuer
		cert.SerialNumber = info.SerialNumber
		cert.IsCA = info.IsCA
		cert.SignatureAlgorithm = info.SignatureAlgorithm
		cert.PublicKeyAlgorithm = info.PublicKeyAlgorithm
		cert.PublicKeyBitLength = info.PublicKeyBitLength
	}

	err = cs.Db.Upsert(cert)
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

func (cs *CertStore) Delete(ctx context.Context, key string) error {
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

	err = cs.Db.Query(datastore.Id(key)).Delete()
	if err != nil {
		return errors.Wrap(err, "failed to delete cert from store")
	}

	return nil
}

type certInfo struct {
	NotBefore          time.Time
	NotAfter           time.Time
	CommonName         string
	DNSNames           []string
	Issuer             string
	SerialNumber       string
	IsCA               bool
	SignatureAlgorithm string
	PublicKeyAlgorithm string
	PublicKeyBitLength int
}

func parseCertInfo(certPEM []byte) (*certInfo, error) {
	info := &certInfo{}

	block, _ := pem.Decode(certPEM)
	if block == nil {
		return info, errors.New("failed to parse PEM block")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return info, errors.Wrap(err, "failed to parse x509 certificate")
	}

	info.NotBefore = cert.NotBefore
	info.NotAfter = cert.NotAfter
	info.CommonName = cert.Subject.CommonName
	info.DNSNames = cert.DNSNames
	info.Issuer = cert.Issuer.CommonName
	if cert.SerialNumber != nil {
		info.SerialNumber = cert.SerialNumber.String()
	}
	info.IsCA = cert.IsCA
	info.SignatureAlgorithm = cert.SignatureAlgorithm.String()
	info.PublicKeyAlgorithm = cert.PublicKeyAlgorithm.String()

	// PublicKey length (best-effort)
	switch pub := cert.PublicKey.(type) {
	case *rsa.PublicKey:
		info.PublicKeyBitLength = pub.N.BitLen()
	case *ecdsa.PublicKey:
		info.PublicKeyBitLength = pub.Curve.Params().BitSize
	case ed25519.PublicKey:
		info.PublicKeyBitLength = len(pub) * 8
	default:
		info.PublicKeyBitLength = 0
	}

	return info, nil
}
