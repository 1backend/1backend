/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package secret_svc

type ErrorResponse struct {
	Error string `json:"error"`
}

type ChecksumAlgorithm string

const (
	ChecksumAlgorithmUnspecified ChecksumAlgorithm = ""
	ChecksumAlgorithmCRC32       ChecksumAlgorithm = "CRC32"
	ChecksumAlgorithmBlake2s     ChecksumAlgorithm = "BLAKE2s"
	ChecksumAlgorithmSha256      ChecksumAlgorithm = "SHA-256"
	ChecksumAlgorithmSha512      ChecksumAlgorithm = "SHA-512"
)

type Secret struct {
	InternalId string `json:"internalId,omitempty" swagger:"ignore"`

	App string `json:"app" binding:"required"`

	// Envar- or slug-like id of the secret
	Id string `json:"id" binding:"required"`

	Value string `json:"value" binding:"required"` // Secret Value

	Readers  []string `json:"readers"`  // Slugs of services/users who can read the secret
	Writers  []string `json:"writers"`  // Slugs of services/users who can modify the secret
	Deleters []string `json:"deleters"` // Slugs of services/users who can delete the secret

	CanChangeReaders  []string `json:"canChangeReaders"`  // Slugs of services/users who can change the readers list
	CanChangeWriters  []string `json:"canChangeWriters"`  // Slugs of services/users who can change the writers list
	CanChangeDeleters []string `json:"canChangeDeleters"` // Slugs of services/users who can change the deleters list

	// Whether the secret is encrypted
	// All secrets are encrypted before written to the DB.
	// This really only exists for write requests to know if the secret is already encrypted.
	// Ie: while most `secret save [id] [value]` commands are probably not encrypted,
	// File based saves, eg. `secret save secretA.yaml` are probably encrypted.
	Encrypted bool `json:"encrypted,omitempty"`

	Checksum          string            `json:"checksum,omitempty"`                          // Checksum of the secret value
	ChecksumAlgorithm ChecksumAlgorithm `json:"checksumAlgorithm,omitempty" example:"CRC32"` // Algorithm used for the checksum (e.g., "CRC32")
}

func (s *Secret) GetId() string {
	return s.InternalId
}

type SecretInput struct {
	App string `json:"app,omitempty"`

	// Envar- or slug-like id of the secret
	Id string `json:"id" binding:"required"`

	Value string `json:"value,omitempty"` // Secret Value

	Readers  []string `json:"readers"`  // Slugs of services/users who can read the secret
	Writers  []string `json:"writers"`  // Slugs of services/users who can modify the secret
	Deleters []string `json:"deleters"` // Slugs of services/users who can delete the secret

	CanChangeReaders  []string `json:"canChangeReaders"`  // Slugs of services/users who can change the readers list
	CanChangeWriters  []string `json:"canChangeWriters"`  // Slugs of services/users who can change the writers list
	CanChangeDeleters []string `json:"canChangeDeleters"` // Slugs of services/users who can change the deleters list

	// Whether the secret is encrypted
	// All secrets are encrypted before written to the DB.
	// This really only exists for write requests to know if the secret is already encrypted.
	// Ie: while most `secret save [id] [value]` commands are probably not encrypted,
	// File based saves, eg. `secret save secretA.yaml` are probably encrypted.
	Encrypted bool `json:"encrypted,omitempty"`

	Checksum          string            `json:"checksum,omitempty"`                          // Checksum of the secret value
	ChecksumAlgorithm ChecksumAlgorithm `json:"checksumAlgorithm,omitempty" example:"CRC32"` // Algorithm used for the checksum (e.g., "CRC32")
}

type ListSecretsRequest struct {
	Id  string   `json:"id"`
	Ids []string `json:"ids"`
}

type ListSecretsResponse struct {
	Secrets []*Secret `json:"secrets"`
}

type SaveSecretsRequest struct {
	Secrets []*SecretInput `json:"secrets"`
}

type SaveSecretsResponse struct {
}

type EncryptValueRequest struct {
	Value  string   `json:"value,omitempty"`
	Values []string `json:"values,omitempty"`
}

type EncryptValueResponse struct {
	Value  string   `json:"value,omitempty"`
	Values []string `json:"values,omitempty"`
}

type DecryptValueRequest struct {
	Value  string   `json:"value,omitempty"`
	Values []string `json:"values,omitempty"`
}

type DecryptValueResponse struct {
	Value  string   `json:"value,omitempty"`
	Values []string `json:"values,omitempty"`
}

type RemoveSecretsRequest struct {
	Id  string   `json:"id"`
	Ids []string `json:"ids"`
}

type RemoveSecretsResponse struct {
}

type IsSecureResponse struct {
	IsSecure bool `json:"isSecure" binding:"required"`
}
