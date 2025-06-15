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
	// Id is the host which this cert is for, e.g., "example.com" or "www.example.com"
	Id string `json:"id" binding:"required" example:"example.com"`

	// When cert record was created
	CreatedAt time.Time `json:"createdAt" binding:"required"`

	// When cert record was last updated
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	// PEM-encoded certificate bundle
	//
	//  -----BEGIN EC PARAMETERS-----
	//  BggqhkjOPQMBBw==
	//  -----END EC PARAMETERS-----
	//  -----BEGIN EC PRIVATE KEY-----
	//  MHcCAQEEIDC3+7pySTQl6WRBuef...
	//  -----END EC PRIVATE KEY-----
	//  -----BEGIN CERTIFICATE-----
	//  MIIBhTCCASugAwIBAgIUQYwE...
	//  -----END CERTIFICATE-----
	Cert string `json:"cert" binding:"required"`

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

type CertInput struct {
	// Id is the host which this cert is for, e.g., "example.com" or "www.example.com"
	Id string `json:"id" binding:"required" example:"example.com"`

	// PEM-encoded certificate bundle
	//
	//  -----BEGIN EC PARAMETERS-----
	//  BggqhkjOPQMBBw==
	//  -----END EC PARAMETERS-----
	//  -----BEGIN EC PRIVATE KEY-----
	//  MHcCAQEEIDC3+7pySTQl6WRBuef...
	//  -----END EC PRIVATE KEY-----
	//  -----BEGIN CERTIFICATE-----
	//  MIIBhTCCASugAwIBAgIUQYwE...
	//  -----END CERTIFICATE-----
	Cert string `json:"cert" binding:"required"`
}

type ListCertsRequest struct {
	Ids []string `json:"ids,omitempty"`
}

type ListCertsResponse struct {
	Certs []Cert `json:"certs"`
}

type SaveCertsRequest struct {
	Certs []CertInput `json:"certs"`
}

type SaveCertsResponse struct {
}
