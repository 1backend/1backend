/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package auth

import (
	"fmt"
	"net/http"
	"strings"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/samber/lo"
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

	// Organizations extracts organizations and organization-internal roles
	// from a request. Given a role string like
	// `user-svc:org:{org_dBZRCej3fo}:admin`, it returns `{"org_dBZRCej3fo": ["admin"]}`.
	Organizations(userSvcPublicKey string, token string) (map[string][]string, error)

	// OrganizationsFromRequest extracts organizations and organization-internal roles
	// from a request. Given a role string like
	// `user-svc:org:{org_dBZRCej3fo}:admin`, it returns `{"org_dBZRCej3fo": ["admin"]}`.
	OrganizationsFromRequest(userSvcPublicKey string, r *http.Request) (map[string][]string, error)
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

func (a AuthorizerImpl) OrganizationsFromRequest(userSvcPublicKey string, r *http.Request) (map[string][]string, error) {
	tokenString, hasToken := a.TokenFromRequest(r)
	if !hasToken {
		return nil, fmt.Errorf("no token found in request")
	}

	return a.Organizations(userSvcPublicKey, tokenString)
}

func (a AuthorizerImpl) Organizations(userSvcPublicKey, token string) (map[string][]string, error) {
	claims, err := a.ParseJWT(userSvcPublicKey, token)
	if err != nil {
		return nil, err
	}

	return ExtractOrganizationRoles(claims.RoleIds), nil
}

func ExtractOrganizationRoles(roleIds []string) map[string][]string {
	ret := map[string][]string{}

	for _, roleId := range roleIds {
		// @todo constant
		if strings.HasPrefix(roleId, "user-svc:org:{") {
			v := strings.Split(roleId, "{")[1]
			parts := strings.Split(v, "}:")
			orgId := parts[0]
			role := parts[1]

			if ret[orgId] == nil {
				ret[orgId] = []string{}
			}

			ret[orgId] = append(ret[orgId], role)
		}
	}

	return ret
}

// OwnsRole determines if the user owns the specified role based on the role ID.
// This simple function is a critical part of the authorization logic.
//
// A user "owns" a role in the following cases:
// - A static role where the role ID is prefixed with the caller's slug.
// - Any dynamic or static role where the caller is an admin.
//
// Examples:
// - A user with the slug "joe-doe" owns roles like "joe-doe:any-custom-role".
// - A user with any slug who has the role "my-service:admin" owns "my-service:user".
// - A user with any slug who has the role "user-svc:org:{%orgId}:admin" owns "user-svc:org:{%orgId}:user".
func OwnsRole(claim *Claims, roleId string) bool {
	// @todo Probably not great in terms of zero trust design ; )
	if lo.Contains(claim.RoleIds, "user-svc:admin") {
		return true
	}

	if strings.HasPrefix(roleId, claim.Slug) {
		return true
	}

	idx := strings.LastIndex(roleId, ":")
	if idx == -1 {
		return false
	}

	rolePrefix := roleId[:idx]

	for _, userRole := range claim.RoleIds {
		idx := strings.LastIndex(userRole, ":")
		if idx == -1 {
			continue
		}

		userRolePrefix := userRole[:idx]
		userRoleSuffix := userRole[idx+1:]

		if userRolePrefix == rolePrefix && userRoleSuffix == "admin" {
			return true
		}
	}

	return false
}

func TokenFromClient(client *openapi.APIClient) string {
	userToken := client.GetConfig().DefaultHeader["Authorization"]
	userToken = strings.Replace(userToken, "Bearer ", "", -1)

	return userToken
}
