package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"time"
)

const tokenFingerprintBytes = 8

// TokenDebugAttrs returns stable, non-secret identifiers that make JWT-related
// failures diagnosable in logs without emitting the raw token.
func TokenDebugAttrs(token string) []slog.Attr {
	if token == "" {
		return nil
	}

	attrs := []slog.Attr{
		slog.String("tokenFingerprint", TokenFingerprint(token)),
	}

	claims, err := (AuthorizerImpl{}).ParseJWTUnverified(token)
	if err != nil {
		return append(attrs, slog.String("tokenClaimsError", err.Error()))
	}

	if claims.UserId != "" {
		attrs = append(attrs, slog.String("tokenUserId", claims.UserId))
	}
	if claims.AppId != "" {
		attrs = append(attrs, slog.String("tokenAppId", claims.AppId))
	}
	if claims.Slug != "" {
		attrs = append(attrs, slog.String("tokenSlug", claims.Slug))
	}
	if claims.Device != "" {
		attrs = append(attrs, slog.String("tokenDevice", claims.Device))
	}
	if claims.ActiveOrganizationId != "" {
		attrs = append(attrs, slog.String("tokenActiveOrganizationId", claims.ActiveOrganizationId))
	}
	if claims.ExpiresAt != nil {
		attrs = append(attrs, slog.String("tokenExpiresAt", claims.ExpiresAt.Time.UTC().Format(time.RFC3339)))
	}

	return attrs
}

func TokenFingerprint(token string) string {
	if token == "" {
		return ""
	}

	hash := TokenHash(token)
	if len(hash) <= tokenFingerprintBytes*2 {
		return hash
	}

	return hash[:tokenFingerprintBytes*2]
}

func TokenHash(token string) string {
	if token == "" {
		return ""
	}

	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
