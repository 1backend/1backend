package proxyservice_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"testing"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/datastore/localstore"
	proxyservice "github.com/1backend/1backend/server/internal/services/proxy"
	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// This is one of those rare tests where we don't do black box testing,
// eg. instead of booting up a whole server we just test the CertStore directly.
// This is because the CertStore PUT operation is not done trough the API but
// by autocert.

func generateTestCert() []byte {
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "test.example.com",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(24 * time.Hour),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
		DNSNames:              []string{"test.example.com", "www.test.example.com"},
	}

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	return certPEM
}

func TestAutocertCache(t *testing.T) {
	ctx := context.Background()
	randomFilePath := t.TempDir() + "/" + uuid.New().String() + ".json"

	localstore, err := localstore.NewLocalStore(
		&proxy.Cert{},
		randomFilePath,
	)
	require.NoError(t, err)

	cache := &proxyservice.CertStore{
		Db: localstore,
	}

	key := "test-cert"
	certPEM := generateTestCert()

	t.Run("put works", func(t *testing.T) {
		err = cache.Put(ctx, key, certPEM)
		require.NoError(t, err)
	})

	t.Run("get works", func(t *testing.T) {
		retrieved, err := cache.Get(ctx, key)
		require.NoError(t, err)
		require.Equal(t, certPEM, retrieved)
	})

	var cert *proxy.Cert

	t.Run("cert metadata is correct", func(t *testing.T) {
		certI, found, err := localstore.Query(
			datastore.Id(key),
		).FindOne()
		require.NoError(t, err)
		require.True(t, found)

		cert = certI.(*proxy.Cert)

		require.Equal(t, "test-cert", cert.Id)
		require.Equal(t, "test.example.com", cert.CommonName)
		require.True(t, cert.IsCA)
		require.Contains(t, cert.DNSNames, "test.example.com")
		require.Contains(t, cert.DNSNames, "www.test.example.com")
		require.NotEmpty(t, cert.PublicKeyBitLength)
	})

	t.Run("put again to ensure update", func(t *testing.T) {
		err = cache.Put(ctx, key, certPEM)
		require.NoError(t, err)
		certI2, _, _ := localstore.Query(
			datastore.Id(key),
		).FindOne()
		cert2 := certI2.(*proxy.Cert)
		require.Equal(t, cert.CreatedAt, cert2.CreatedAt)      // CreatedAt should remain
		require.True(t, cert2.UpdatedAt.After(cert.UpdatedAt)) // UpdatedAt should be updated
	})

	t.Run("delete works", func(t *testing.T) {
		err = cache.Delete(ctx, key)
		require.NoError(t, err)
	})

	t.Run("ensure deletion", func(t *testing.T) {
		_, found, _ := localstore.Query(
			datastore.Id(key),
		).FindOne()
		require.False(t, found)
	})

	t.Run("get after delete fails", func(t *testing.T) {
		_, err = cache.Get(ctx, key)
		require.Error(t, err)
	})
}
