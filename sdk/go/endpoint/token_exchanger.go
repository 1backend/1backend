package endpoint

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
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

	// AppFromRequestHost extracts the app name from the request host.
	// This is useful for example when you want to exchange a service token from the default "unnamed" app
	// to a specific app based on the request host.
	//
	// Especially useful when the request doesn't contain a token (eg. the endpoint is public).
	// When the endpoint is authenticated, the token is usuallly used to determine the app.
	AppFromRequestHost(request *http.Request) (string, error)
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
	newApp string,
) (string, error) {

	cacheKey := generateTokenExchangeKey(token, newApp)
	if value, found := te.cache.Get(cacheKey); found {
		if cached, ok := value.(*TokenExchangeResponse); ok {
			return cached.Token, nil
		}
	}

	req := openapi.UserSvcExchangeTokenRequest{
		NewApp: newApp,
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

func (te *TokenExchangerImpl) AppFromRequestHost(request *http.Request) (
	string, error,
) {
	if request == nil || request.Host == "" {
		return "", errors.New("request or request host is empty")
	}

	host, _, err := net.SplitHostPort(request.Host)
	if err != nil {
		return "", errors.Wrap(err, "failed to split host and port from request host")
	}
	if host == "" {
		return "", errors.New("host is empty")
	}

	return host, nil
}

func generateTokenExchangeKey(token, newApp string) string {
	hash := sha256.Sum256([]byte(token))
	return fmt.Sprintf("%s:%s:%s", hex.EncodeToString(hash[:]), newApp)
}
