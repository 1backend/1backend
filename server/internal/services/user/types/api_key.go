/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import "time"

type ApiKey struct {
	InternalId string `json:"internalId,omitempty" swagger:"ignore"`

	Id     string `json:"id" binding:"required"`
	AppId  string `json:"appId" binding:"required"`
	UserId string `json:"userId" binding:"required"`

	Name string `json:"name,omitempty"`

	// Prefix is safe for display and audit logs. The full key is only returned
	// once at creation time.
	Prefix string `json:"prefix,omitempty"`

	// SecretHash is the SHA-256 hash of the random secret portion of the key.
	// The raw API key is never stored.
	SecretHash string `json:"secretHash,omitempty" swagger:"ignore"`

	ActiveOrganizationId string `json:"activeOrganizationId,omitempty"`

	CreatedAt  time.Time  `json:"createdAt" binding:"required"`
	UpdatedAt  time.Time  `json:"updatedAt" binding:"required"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	ExpiresAt  *time.Time `json:"expiresAt,omitempty"`
	RevokedAt  *time.Time `json:"revokedAt,omitempty"`
}

func (k *ApiKey) GetId() string {
	return k.Id
}

type ApiKeyView struct {
	Id                   string     `json:"id" binding:"required"`
	AppId                string     `json:"appId" binding:"required"`
	UserId               string     `json:"userId" binding:"required"`
	Name                 string     `json:"name,omitempty"`
	Prefix               string     `json:"prefix,omitempty"`
	ActiveOrganizationId string     `json:"activeOrganizationId,omitempty"`
	CreatedAt            time.Time  `json:"createdAt" binding:"required"`
	UpdatedAt            time.Time  `json:"updatedAt" binding:"required"`
	LastUsedAt           *time.Time `json:"lastUsedAt,omitempty"`
	ExpiresAt            *time.Time `json:"expiresAt,omitempty"`
	RevokedAt            *time.Time `json:"revokedAt,omitempty"`
}

type CreateApiKeyRequest struct {
	AppHost string `json:"appHost,omitempty"`
	AppId   string `json:"appId,omitempty"`
	UserId  string `json:"userId,omitempty"`
	Name    string `json:"name,omitempty"`

	// ActiveOrganizationId is optional. If omitted, the current token's active
	// organization is used. If explicitly set to an empty string, the API key
	// authenticates without an active organization.
	ActiveOrganizationId *string `json:"activeOrganizationId,omitempty"`

	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

type CreateApiKeyResponse struct {
	ApiKey ApiKeyView `json:"apiKey" binding:"required"`

	// Key is the raw bearer credential. It is returned only once.
	Key string `json:"key" binding:"required"`
}

type ListApiKeysRequest struct {
	AppHost        string   `json:"appHost,omitempty"`
	AppId          string   `json:"appId,omitempty"`
	UserId         string   `json:"userId,omitempty"`
	Id             string   `json:"id,omitempty"`
	Ids            []string `json:"ids,omitempty"`
	IncludeRevoked bool     `json:"includeRevoked,omitempty"`
}

type ListApiKeysResponse struct {
	ApiKeys []ApiKeyView `json:"apiKeys" binding:"required"`
}

type RevokeApiKeysRequest struct {
	UserId string   `json:"userId,omitempty"`
	Id     string   `json:"id,omitempty"`
	Ids    []string `json:"ids,omitempty"`
}

type RevokeApiKeysResponse struct{}

type ExchangeApiKeyResponse struct {
	Token *Token `json:"token" binding:"required"`
}
