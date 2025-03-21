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

	// IsAdmin returns true if the user has
	// the `user-svc:admin` role.
	IsAdmin(userSvcPublicKey string, token string) (bool, error)

	// IsAdminFromRequest returns true if the user has
	// the `user-svc:admin` role.
	IsAdminFromRequest(userSvcPublicKey string, r *http.Request) (bool, error)

	// RolesInOrganization extracts organization-specific roles
	// from a request. Given a role string like
	// `user-svc:org:{org_dBZRCej3fo}:admin`, it returns `[admin]`.
	RolesInOrganization(userSvcPublicKey string, token string) ([]string, error)

	// RolesInOrganizationFromRequest extracts organization-specific roles
	// from a request. Given a role string like
	// `user-svc:org:{org_dBZRCej3fo}:admin`, it returns `[admin]`.
	RolesInOrganizationFromRequest(userSvcPublicKey string, r *http.Request) ([]string, error)
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

func (a AuthorizerImpl) RolesInOrgFromRequest(userSvcPublicKey string, r *http.Request) ([]string, error) {
	tokenString, hasToken := a.TokenFromRequest(r)
	if !hasToken {
		return nil, fmt.Errorf("no token found in request")
	}

	return a.RolesInOrg(userSvcPublicKey, tokenString)
}

func (a AuthorizerImpl) RolesInOrg(userSvcPublicKey, token string) ([]string, error) {
	claims, err := a.ParseJWT(userSvcPublicKey, token)
	if err != nil {
		return nil, err
	}

	return ExtractOrganizationRoles(claims.RoleIds), nil
}

func ExtractOrganizationRoles(roleIds []string) []string {
	ret := []string{}

	for _, roleId := range roleIds {
		// @todo constant
		if strings.HasPrefix(roleId, "user-svc:org:{") {
			v := strings.Split(roleId, "{")[1]
			v = strings.Split(v, "}:")[1]
			ret = append(ret, v)
		}
	}

	return ret
}
