package sdk

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

// Authorizer can extract roles from tokens.
// This interface is not the only thing that does authorization, however.
// Authorization also happens by calling the User Svc to check if a user has a specific permission to call an endpoint.
type Authorizer interface {
	TokenFromRequest(r *http.Request) (string, bool)
	ParseJWT(userSvcPublicKey, token string) (*Claims, error)
	ParseJWTFromRequest(userSvcPublicKey string, r *http.Request) (*Claims, error)
	IsAdmin(userSvcPublicKey string, token string) (bool, error)
	IsAdminFromRequest(userSvcPublicKey string, r *http.Request) (bool, error)
}

type AuthorizerImpl struct{}

func (a AuthorizerImpl) ParseJWTFromRequest(userSvcPublicKey string, r *http.Request) (*Claims, error) {
	tokenString, hasToken := a.TokenFromRequest(r)
	if !hasToken {
		return nil, fmt.Errorf("no token found in request")
	}

	return a.ParseJWT(userSvcPublicKey, tokenString)
}

func (a AuthorizerImpl) ParseJWT(userSvcPublicKey, token string) (*Claims, error) {
	publicKey, err := PublicKeyFromString(userSvcPublicKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get public key from string")
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %v", err)
	}

	if claims, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid JWT token")
}

func (a AuthorizerImpl) TokenFromRequest(r *http.Request) (string, bool) {
	authHeader := r.Header.Get("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	if authHeader == "" || authHeader == "Bearer" {
		return "", false
	}

	return authHeader, true
}

func (a AuthorizerImpl) IsAdminFromRequest(userSvcPublicKey string, r *http.Request) (bool, error) {
	tokenString, hasToken := a.TokenFromRequest(r)
	if !hasToken {
		return false, fmt.Errorf("no token found in request")
	}

	return a.IsAdmin(userSvcPublicKey, tokenString)
}

func (a AuthorizerImpl) IsAdmin(userSvcPublicKey, token string) (bool, error) {
	claims, err := a.ParseJWT(userSvcPublicKey, token)
	if err != nil {
		return false, err
	}

	for _, roleId := range claims.RoleIds {
		// @todo remove constant
		if roleId == "user-svc:admin" {
			return true, nil
		}
	}

	return false, nil
}
