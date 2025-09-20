package userservice

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func generateRSAKeys(bits int) (privateKeyPem, publicKeyPem string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", fmt.Errorf("error generating RSA key: %v", err)
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privateKeyPem = string(pem.EncodeToMemory(privateKeyBlock))

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", fmt.Errorf("error marshaling public key: %v", err)
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicKeyPem = string(pem.EncodeToMemory(publicKeyBlock))

	return privateKeyPem, publicKeyPem, nil
}

func privateKeyFromString(privateKeyPem string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyPem))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf(
			"failed to decode PEM block containing private key",
		)
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA private key: %v", err)
	}

	return privateKey, nil
}

func (s *UserService) generateJWT(
	appId string,
	user *usertypes.User,
	roles []string,
	activeOrganizationId string,
	privateKey *rsa.PrivateKey,
	device string,
) (*auth.Claims, string, error) {
	now := time.Now()

	claims := &auth.Claims{
		AppId:                appId,
		UserId:               user.Id,
		Slug:                 user.Slug,
		Roles:                roles,
		Device:               device,
		ActiveOrganizationId: activeOrganizationId,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        sdk.Id(""),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.options.TokenExpiration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		// Log the error if signing fails
		log.Printf("Error signing token: %v\n", err)
		return nil, "", fmt.Errorf("failed to sign token: %v", err)
	}

	return claims, tokenString, nil
}
