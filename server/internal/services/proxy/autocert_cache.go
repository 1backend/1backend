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
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/1backend/1backend/sdk/go/secrets"
	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
	"golang.org/x/crypto/acme/autocert"

	"github.com/pkg/errors"
)

type CertStore struct {
	EncryptionKey    string
	SyncCertsToFiles bool
	CertFolder       string
	Db               datastore.DataStore
}

//
// We log errors in these methods because we're not sure how autocert logs the errors.
//

func (cs *CertStore) Get(ctx context.Context, key string) ([]byte, error) {
	if cs.EncryptionKey == "" {
		return nil, errors.New("encryption key is not set")
	}

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
		return nil, autocert.ErrCacheMiss
	}

	cert, ok := certI.(*proxy.Cert)
	if !ok {
		return nil, errors.Errorf("expected cert type, got %T", certI)
	}

	decryptedData, err := secrets.Decrypt(cert.Cert, cs.EncryptionKey)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decrypt cert data for key '%s'", key)
	}

	data := []byte(decryptedData)

	if len(data) == 0 {
		return nil, errors.Errorf("cert data is empty for key '%s'", key)
	}

	return data, nil
}

func (cs *CertStore) Put(ctx context.Context, key string, data []byte) error {
	if cs.EncryptionKey == "" {
		return errors.New("encryption key is not set")
	}

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

	logger.Info(
		"Storing cert",
		slog.String("key", key),
		slog.Int("data_length", len(data)),
	)

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

	encryptedData, err := secrets.Encrypt(string(data), cs.EncryptionKey)
	if err != nil {
		return errors.Wrapf(err, "failed to encrypt cert data for key '%s'", key)
	}

	cert := &proxy.Cert{
		Id:   key,
		Cert: encryptedData,
	}

	if cs.SyncCertsToFiles {
		err := WriteCertKeyChainToFilesWithHost(
			cs.CertFolder,
			key,
			string(data),
		)
		if err != nil {
			logger.Error(
				"Failed to write cert to files",
				slog.String("key", key),
				slog.Any("error", err),
			)
		}
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

	rest := certPEM
	for {
		var block *pem.Block
		block, rest = pem.Decode(rest)
		if block == nil {
			break // No more PEM blocks
		}

		if block.Type != "CERTIFICATE" {
			continue // Skip non-certificate blocks (e.g., private keys)
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return info, errors.Wrap(err, "failed to parse x509 certificate")
		}

		// Fill info from the first certificate we find
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

		break // Stop after processing the first certificate
	}

	if info.CommonName == "" {
		return info, errors.New("no valid certificate found in PEM")
	}

	return info, nil
}

func WriteCertKeyChainToFilesWithHost(
	outputDir string,
	fallbackHost string,
	pemData string,
) error {
	var certBlock *pem.Block
	var keyBlock *pem.Block
	var chainPEM []byte
	var hostname string

	rest := []byte(pemData)

	for {
		var block *pem.Block
		block, rest = pem.Decode(rest)
		if block == nil {
			break
		}

		switch block.Type {
		case "CERTIFICATE":
			if certBlock == nil {
				certBlock = block
				cert, err := x509.ParseCertificate(block.Bytes)
				if err == nil {
					if cert.Subject.CommonName != "" {
						hostname = cert.Subject.CommonName
					} else if len(cert.DNSNames) > 0 {
						hostname = cert.DNSNames[0]
					}
				}
			} else {
				chainPEM = append(chainPEM, pem.EncodeToMemory(block)...)
			}
		case "PRIVATE KEY", "RSA PRIVATE KEY", "EC PRIVATE KEY", "ED25519 PRIVATE KEY":
			if keyBlock == nil {
				keyBlock = block
			}
		}
	}

	if certBlock == nil {
		return errors.New("no certificate found in PEM")
	}
	if hostname == "" {
		if fallbackHost != "" {
			hostname = fallbackHost
		} else {
			return errors.New("no hostname found in certificate and no fallback provided")
		}
	}

	liveDir := filepath.Join(outputDir, "live", hostname)
	if err := os.MkdirAll(liveDir, 0755); err != nil {
		return errors.Wrap(err, "failed to create live directory")
	}

	certPath := filepath.Join(liveDir, "cert.pem")
	keyPath := filepath.Join(liveDir, "privkey.pem")
	chainPath := filepath.Join(liveDir, "chain.pem")
	fullchainPath := filepath.Join(liveDir, "fullchain.pem")

	if err := os.WriteFile(certPath, pem.EncodeToMemory(certBlock), 0644); err != nil {
		return errors.Wrap(err, "failed to write cert.pem")
	}

	if keyBlock != nil {
		if err := os.WriteFile(keyPath, pem.EncodeToMemory(keyBlock), 0600); err != nil {
			return errors.Wrap(err, "failed to write privkey.pem")
		}
	}

	if len(chainPEM) > 0 {
		if err := os.WriteFile(chainPath, chainPEM, 0644); err != nil {
			return errors.Wrap(err, "failed to write chain.pem")
		}

		fullchain := append(pem.EncodeToMemory(certBlock), chainPEM...)
		if err := os.WriteFile(fullchainPath, fullchain, 0644); err != nil {
			return errors.Wrap(err, "failed to write fullchain.pem")
		}
	} else {
		// If no chain, fullchain == cert only
		if err := os.WriteFile(fullchainPath, pem.EncodeToMemory(certBlock), 0644); err != nil {
			return errors.Wrap(err, "failed to write fullchain.pem")
		}
	}

	return nil
}
