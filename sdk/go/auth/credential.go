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
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId string   `json:"oui,omitempty"` // `oui`: 1backend user ids
	Slug   string   `json:"olu,omitempty"` // `olu`: 1backend slug
	Roles  []string `json:"ori,omitempty"` // `ori`: 1backend role ids

	// Organization IDs are already included within role IDs
	// (e.g., `user-svc:org:{org_dBZRCej3fo}:admin`).
	// Helper functions make it easy to extract them, so they aren't stored separately to save space.
	// However, role IDs don’t specify the active or default organization.
	// This field explicitly provides that information.
	ActiveOrganizationId string `json:"oao,omitempty"` // `ooi`: 1backend active organization id

	jwt.RegisteredClaims
}

type Credential struct {
	Slug     string `json:"slug,omitempty"`
	Contact  string `json:"contact,omitempty"`
	Password string `json:"password,omitempty"`
}

func (c *Credential) GetId() string {
	return c.Contact
}

func PublicKeyFromString(publicKeyPem string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPem))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	// Type assertion to convert from interface{} to *rsa.PublicKey
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return rsaPub, nil
}
