/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxy_svc

import "time"

type Cert struct {
	// Unique cert ID
	Id string `json:"id"`

	// When cert record was created
	CreatedAt time.Time `json:"createdAt"`

	// When cert record was last updated
	UpdatedAt time.Time `json:"updatedAt"`

	// Base64 encoded PEM certificate
	Cert string `json:"cert"`

	// Cert validity start time
	NotBefore time.Time `json:"notBefore"`

	// Cert validity end time
	NotAfter time.Time `json:"notAfter"`

	// Subject Common Name (typically domain)
	CommonName string `json:"commonName"`

	// Subject Alternative Names (covered domains)
	DNSNames []string `json:"dnsNames"`

	// Certificate issuer name (e.g., Let's Encrypt)
	Issuer string `json:"issuer"`

	// Unique certificate serial number
	SerialNumber string `json:"serialNumber"`

	// Whether cert is a CA (usually false for LE certs)
	IsCA bool `json:"isCA"`

	// Algorithm used to sign the cert (e.g., SHA256-RSA)
	SignatureAlgorithm string `json:"signatureAlgorithm"`

	// Public key algorithm (e.g., RSA, ECDSA)
	PublicKeyAlgorithm string `json:"publicKeyAlgorithm"`

	// Bit length of the public key
	PublicKeyBitLength int `json:"publicKeyBitLength"`
}

func (c *Cert) GetId() string {
	return c.Id
}

type ListCertsRequest struct {
	Ids []string `json:"ids,omitempty"`
}

type ListCertsResponse struct {
	Certs []Cert `json:"certs"`
}
