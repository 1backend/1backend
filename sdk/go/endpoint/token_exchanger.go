package endpoint

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
)

type TokenExchangeResponse struct {
	Token     string
	ExpiresAt time.Time
}

// TokenExchanger is an interface for exchanging tokens between different apps or devices.
// It caches the exchanged tokens to avoid unnecessary calls to the user service.
type TokenExchanger interface {
	// ExchangeToken exchanges the original token for a new token for a different app or device.
	// It caches the response using a key derived from the original token, new app, and new device.
	ExchangeToken(
		request *http.Request,
		newApp string,
		newDevice *string,
	) (string, error)
}

type TokenExchangerImpl struct {
	Testing       bool
	clientFactory client.ClientFactory
	cache         *ristretto.Cache
}

func NewTokenExchanger(
	clientFactory client.ClientFactory,
) TokenExchanger {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e5,
		MaxCost:     1 << 20,
		BufferItems: 64,
	})
	if err != nil {
		panic(errors.Wrap(err, "failed to create ristretto cache").Error())
	}

	return &TokenExchangerImpl{
		clientFactory: clientFactory,
		cache:         cache,
	}
}

func (te *TokenExchangerImpl) ExchangeToken(
	request *http.Request,
	newApp string,
	newDevice *string,
) (string, error) {
	originalToken := strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer ")
	if originalToken == "" {
		return "", errors.New("missing original Authorization token")
	}

	cacheKey := generateTokenExchangeKey(originalToken, newApp, newDevice)
	if value, found := te.cache.Get(cacheKey); found {
		if cached, ok := value.(*TokenExchangeResponse); ok {
			return cached.Token, nil
		}
	}

	req := openapi.UserSvcExchangeTokenRequest{
		NewApp:    newApp,
		NewDevice: newDevice,
	}

	rsp, _, err := te.clientFactory.Client(
		client.WithTokenFromRequest(request),
	).
		UserSvcAPI.
		ExchangeToken(request.Context()).
		Body(req).
		Execute()
	if err != nil {
		return "", errors.Wrap(err, "failed to exchange token")
	}

	expiresAt, err := time.Parse(time.RFC3339, rsp.Token.ExpiresAt)
	if err != nil {
		return "", errors.Wrap(err, "invalid expiration time in response")
	}

	ttl := time.Until(expiresAt)
	te.cache.SetWithTTL(cacheKey, &TokenExchangeResponse{
		Token:     rsp.Token.Token,
		ExpiresAt: expiresAt,
	}, 1, ttl)

	if te.Testing {
		te.cache.Wait()
	}

	return rsp.Token.Token, nil
}

func generateTokenExchangeKey(token, newApp string, newDevice *string) string {
	device := "unknown"
	if newDevice != nil && *newDevice != "" {
		device = *newDevice
	}

	hash := sha256.Sum256([]byte(token))
	return fmt.Sprintf("%s:%s:%s", hex.EncodeToString(hash[:]), newApp, device)
}
