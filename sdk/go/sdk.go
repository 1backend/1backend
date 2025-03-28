/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package sdk

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"

	openapi "github.com/1backend/1backend/clients/go"
)

type Claims struct {
	UserId  string   `json:"oui"` // `oui`: 1backend user ids
	Slug    string   `json:"olu"` // `olu`: 1backend slug
	RoleIds []string `json:"ori"` // `ori`: 1backend role ids
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

var sonyFlake *sonyflake.Sonyflake

func init() {
	sonyFlake = sonyflake.NewSonyflake(sonyflake.Settings{})
	if sonyFlake == nil {
		panic("Sonyflake not created")
	}
}

const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Id(prefix string) string {
	number, err := sonyFlake.NextID()
	if err != nil {
		panic(err)
	}

	if number == 0 {
		return string(base62[0])
	}

	b := make([]byte, 0)
	for number > 0 {
		remainder := number % 62
		number = number / 62
		b = append([]byte{base62[remainder]}, b...)
	}

	return prefix + "_" + string(b)
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

func Marshal(value any) *string {
	jsonBytes, _ := json.Marshal(value)

	v := string(jsonBytes)
	return &v
}

// OpenAPIError checks if an error is a GenericOpenAPIError and returns a meaningful error.
func OpenAPIError(err error) error {
	if err == nil {
		return nil
	}

	// Check if it's a GenericOpenAPIError
	if apiErr, ok := err.(*openapi.GenericOpenAPIError); ok {
		var errorResponse map[string]interface{}
		if unmarshalErr := json.Unmarshal(apiErr.Body(), &errorResponse); unmarshalErr == nil {
			if message, exists := errorResponse["error"]; exists {
				return errors.New(message.(string))
			}
			return fmt.Errorf("unknown error format: %v", string(apiErr.Body()))
		}
		return fmt.Errorf("failed to unmarshal API error response: %v", string(apiErr.Body()))
	}

	// Return the original error if it's not a GenericOpenAPIError
	return err
}
