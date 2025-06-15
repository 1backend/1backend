package proxyservice_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
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
		EncryptionKey: "testEncryptionKeytestEncryptionK",
		Db:            localstore,
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

type certTestCase struct {
	Cert         string
	ExpectedHost string
	ExpectChain  bool // Whether we expect chain.pem to be created
}

var certCases = []certTestCase{
	{
		Cert:
		// openssl ecparam -genkey -name prime256v1 -out ec.key
		// openssl req -new -x509 -key ec.key -out ec.crt -days 365 -subj "/CN=test-ec.local"
		`-----BEGIN EC PARAMETERS-----
BggqhkjOPQMBBw==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIDC3+7pySTQl6WRBuefZdxg/t5Zec6BSlGww+w1z/6N5oAoGCCqGSM49
AwEHoUQDQgAEowxQJ66sxQUzPmWsUWtHHVsNaysH3ytAqlwWk1OmqeaiTqfiomQd
dkmLJt8ur4uWzXEz0rszpx+kkKC0O5wK1Q==
-----END EC PRIVATE KEY-----
-----BEGIN CERTIFICATE-----
MIIBhTCCASugAwIBAgIUQYwEd3SCWKXIQpi1LmyH5NmfLUMwCgYIKoZIzj0EAwIw
GDEWMBQGA1UEAwwNdGVzdC1lYy5sb2NhbDAeFw0yNTA2MTUxMTM5NDhaFw0yNjA2
MTUxMTM5NDhaMBgxFjAUBgNVBAMMDXRlc3QtZWMubG9jYWwwWTATBgcqhkjOPQIB
BggqhkjOPQMBBwNCAASjDFAnrqzFBTM+ZaxRa0cdWw1rKwffK0CqXBaTU6ap5qJO
p+KiZB12SYsm3y6vi5bNcTPSuzOnH6SQoLQ7nArVo1MwUTAdBgNVHQ4EFgQUgcgn
5D79ptSUBKu3TFrfTljdUTwwHwYDVR0jBBgwFoAUgcgn5D79ptSUBKu3TFrfTljd
UTwwDwYDVR0TAQH/BAUwAwEB/zAKBggqhkjOPQQDAgNIADBFAiA6llltovI3Gt8w
z8CriygJzZdjwRQxeTTUyYTxpZb43gIhAJqAMTOb00C5QiNZZji04AyACQX1z6Dq
ODVqQ2xusrR7
-----END CERTIFICATE-----`,
		ExpectedHost: "test-ec.local",
		ExpectChain:  false,
	},
}

func TestWriteCertKeyChainToFilesWithHost(t *testing.T) {
	for i, cas := range certCases {
		fallbackHost := fmt.Sprintf("test-cert-%v", i)

		t.Run(fallbackHost, func(t *testing.T) {
			certFolder := t.TempDir()

			err := proxyservice.WriteCertKeyChainToFilesWithHost(
				certFolder,
				fallbackHost,
				cas.Cert,
			)
			require.NoError(t, err)

			expectedHost := cas.ExpectedHost
			if expectedHost == "" {
				expectedHost = fallbackHost
			}

			basePath := filepath.Join(certFolder, "live", expectedHost)

			certFilePath := filepath.Join(basePath, "cert.pem")
			keyFilePath := filepath.Join(basePath, "privkey.pem")
			chainFilePath := filepath.Join(basePath, "chain.pem")
			fullchainFilePath := filepath.Join(basePath, "fullchain.pem")

			_, err = os.Stat(certFilePath)
			require.NoError(t, err)

			_, err = os.Stat(keyFilePath)
			require.NoError(t, err)

			_, err = os.Stat(fullchainFilePath)
			require.NoError(t, err)

			if cas.ExpectChain {
				_, err = os.Stat(chainFilePath)
				require.NoError(t, err)
			}
		})
	}
}
