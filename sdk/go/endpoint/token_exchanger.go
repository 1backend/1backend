package endpoint

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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
	// ExchangeToken exchanges the original token for a new token for a different app.
	// It caches the exchanged tokens to avoid unnecessary calls to the user service.
	ExchangeToken(
		ctx context.Context,
		token string,
		newApp string,
	) (string, error)

	AppIdFromHost(ctx context.Context, appHost string) (string, error)
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
	ctx context.Context,
	token string,
	newAppHost string,
) (string, error) {

	cacheKey := generateTokenExchangeKey(token, newAppHost)
	if value, found := te.cache.Get(cacheKey); found {
		if cached, ok := value.(*TokenExchangeResponse); ok {
			return cached.Token, nil
		}
	}

	req := openapi.UserSvcExchangeTokenRequest{
		NewAppHost: newAppHost,
	}

	rsp, _, err := te.clientFactory.Client(
		client.WithToken(token),
	).
		UserSvcAPI.
		ExchangeToken(ctx).
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

func (te *TokenExchangerImpl) AppIdFromHost(
	ctx context.Context,
	appHost string,
) (
	string, error,
) {
	if appHost == "" {
		return "", errors.New("app host is empty")
	}

	if strings.HasPrefix(appHost, "app_") {
		return "", fmt.Errorf("appHost is already an app id: '%s'", appHost)
	}

	rsp, _, err := te.clientFactory.Client().
		UserSvcAPI.
		ReadApp(ctx).
		Body(openapi.UserSvcReadAppRequest{
			Host: &appHost,
		}).
		Execute()
	if err != nil {
		return "", errors.Wrap(err, "failed to list app")
	}

	if rsp.App.Id == "" {
		return "", fmt.Errorf("app not found for host: '%s'", appHost)
	}

	return rsp.App.Id, nil
}

func generateTokenExchangeKey(token, newApp string) string {
	hash := sha256.Sum256([]byte(token))
	return fmt.Sprintf("%s:%s", hex.EncodeToString(hash[:]), newApp)
}
