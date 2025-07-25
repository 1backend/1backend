/*
1Backend

AI-native microservices platform.

API version: 0.7.6
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ProxySvcCert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxySvcCert{}

// ProxySvcCert struct for ProxySvcCert
type ProxySvcCert struct {
	// PEM-encoded certificate bundle   -----BEGIN EC PARAMETERS-----  BggqhkjOPQMBBw==  -----END EC PARAMETERS-----  -----BEGIN EC PRIVATE KEY-----  MHcCAQEEIDC3+7pySTQl6WRBuef...  -----END EC PRIVATE KEY-----  -----BEGIN CERTIFICATE-----  MIIBhTCCASugAwIBAgIUQYwE...  -----END CERTIFICATE-----
	Cert string `json:"cert"`
	// Subject Common Name (typically domain)
	CommonName *string `json:"commonName,omitempty"`
	// When cert record was created
	CreatedAt string `json:"createdAt"`
	// Subject Alternative Names (covered domains)
	DnsNames []string `json:"dnsNames,omitempty"`
	// Id is the host which this cert is for, e.g., \"example.com\" or \"www.example.com\"
	Id string `json:"id"`
	// Whether cert is a CA (usually false for LE certs)
	IsCA *bool `json:"isCA,omitempty"`
	// Certificate issuer name (e.g., Let's Encrypt)
	Issuer *string `json:"issuer,omitempty"`
	// Cert validity end time
	NotAfter *string `json:"notAfter,omitempty"`
	// Cert validity start time
	NotBefore *string `json:"notBefore,omitempty"`
	// Public key algorithm (e.g., RSA, ECDSA)
	PublicKeyAlgorithm *string `json:"publicKeyAlgorithm,omitempty"`
	// Bit length of the public key
	PublicKeyBitLength *int32 `json:"publicKeyBitLength,omitempty"`
	// Unique certificate serial number
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Algorithm used to sign the cert (e.g., SHA256-RSA)
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`
	// When cert record was last updated
	UpdatedAt string `json:"updatedAt"`
}

type _ProxySvcCert ProxySvcCert

// NewProxySvcCert instantiates a new ProxySvcCert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxySvcCert(cert string, createdAt string, id string, updatedAt string) *ProxySvcCert {
	this := ProxySvcCert{}
	this.Cert = cert
	this.CreatedAt = createdAt
	this.Id = id
	this.UpdatedAt = updatedAt
	return &this
}

// NewProxySvcCertWithDefaults instantiates a new ProxySvcCert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxySvcCertWithDefaults() *ProxySvcCert {
	this := ProxySvcCert{}
	return &this
}

// GetCert returns the Cert field value
func (o *ProxySvcCert) GetCert() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cert
}

// GetCertOk returns a tuple with the Cert field value
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cert, true
}

// SetCert sets field value
func (o *ProxySvcCert) SetCert(v string) {
	o.Cert = v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *ProxySvcCert) GetCommonName() string {
	if o == nil || IsNil(o.CommonName) {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetCommonNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommonName) {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *ProxySvcCert) HasCommonName() bool {
	if o != nil && !IsNil(o.CommonName) {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *ProxySvcCert) SetCommonName(v string) {
	o.CommonName = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProxySvcCert) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProxySvcCert) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDnsNames returns the DnsNames field value if set, zero value otherwise.
func (o *ProxySvcCert) GetDnsNames() []string {
	if o == nil || IsNil(o.DnsNames) {
		var ret []string
		return ret
	}
	return o.DnsNames
}

// GetDnsNamesOk returns a tuple with the DnsNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetDnsNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsNames) {
		return nil, false
	}
	return o.DnsNames, true
}

// HasDnsNames returns a boolean if a field has been set.
func (o *ProxySvcCert) HasDnsNames() bool {
	if o != nil && !IsNil(o.DnsNames) {
		return true
	}

	return false
}

// SetDnsNames gets a reference to the given []string and assigns it to the DnsNames field.
func (o *ProxySvcCert) SetDnsNames(v []string) {
	o.DnsNames = v
}

// GetId returns the Id field value
func (o *ProxySvcCert) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProxySvcCert) SetId(v string) {
	o.Id = v
}

// GetIsCA returns the IsCA field value if set, zero value otherwise.
func (o *ProxySvcCert) GetIsCA() bool {
	if o == nil || IsNil(o.IsCA) {
		var ret bool
		return ret
	}
	return *o.IsCA
}

// GetIsCAOk returns a tuple with the IsCA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetIsCAOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCA) {
		return nil, false
	}
	return o.IsCA, true
}

// HasIsCA returns a boolean if a field has been set.
func (o *ProxySvcCert) HasIsCA() bool {
	if o != nil && !IsNil(o.IsCA) {
		return true
	}

	return false
}

// SetIsCA gets a reference to the given bool and assigns it to the IsCA field.
func (o *ProxySvcCert) SetIsCA(v bool) {
	o.IsCA = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *ProxySvcCert) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *ProxySvcCert) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *ProxySvcCert) SetIssuer(v string) {
	o.Issuer = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *ProxySvcCert) GetNotAfter() string {
	if o == nil || IsNil(o.NotAfter) {
		var ret string
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetNotAfterOk() (*string, bool) {
	if o == nil || IsNil(o.NotAfter) {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *ProxySvcCert) HasNotAfter() bool {
	if o != nil && !IsNil(o.NotAfter) {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given string and assigns it to the NotAfter field.
func (o *ProxySvcCert) SetNotAfter(v string) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *ProxySvcCert) GetNotBefore() string {
	if o == nil || IsNil(o.NotBefore) {
		var ret string
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetNotBeforeOk() (*string, bool) {
	if o == nil || IsNil(o.NotBefore) {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *ProxySvcCert) HasNotBefore() bool {
	if o != nil && !IsNil(o.NotBefore) {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given string and assigns it to the NotBefore field.
func (o *ProxySvcCert) SetNotBefore(v string) {
	o.NotBefore = &v
}

// GetPublicKeyAlgorithm returns the PublicKeyAlgorithm field value if set, zero value otherwise.
func (o *ProxySvcCert) GetPublicKeyAlgorithm() string {
	if o == nil || IsNil(o.PublicKeyAlgorithm) {
		var ret string
		return ret
	}
	return *o.PublicKeyAlgorithm
}

// GetPublicKeyAlgorithmOk returns a tuple with the PublicKeyAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetPublicKeyAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKeyAlgorithm) {
		return nil, false
	}
	return o.PublicKeyAlgorithm, true
}

// HasPublicKeyAlgorithm returns a boolean if a field has been set.
func (o *ProxySvcCert) HasPublicKeyAlgorithm() bool {
	if o != nil && !IsNil(o.PublicKeyAlgorithm) {
		return true
	}

	return false
}

// SetPublicKeyAlgorithm gets a reference to the given string and assigns it to the PublicKeyAlgorithm field.
func (o *ProxySvcCert) SetPublicKeyAlgorithm(v string) {
	o.PublicKeyAlgorithm = &v
}

// GetPublicKeyBitLength returns the PublicKeyBitLength field value if set, zero value otherwise.
func (o *ProxySvcCert) GetPublicKeyBitLength() int32 {
	if o == nil || IsNil(o.PublicKeyBitLength) {
		var ret int32
		return ret
	}
	return *o.PublicKeyBitLength
}

// GetPublicKeyBitLengthOk returns a tuple with the PublicKeyBitLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetPublicKeyBitLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.PublicKeyBitLength) {
		return nil, false
	}
	return o.PublicKeyBitLength, true
}

// HasPublicKeyBitLength returns a boolean if a field has been set.
func (o *ProxySvcCert) HasPublicKeyBitLength() bool {
	if o != nil && !IsNil(o.PublicKeyBitLength) {
		return true
	}

	return false
}

// SetPublicKeyBitLength gets a reference to the given int32 and assigns it to the PublicKeyBitLength field.
func (o *ProxySvcCert) SetPublicKeyBitLength(v int32) {
	o.PublicKeyBitLength = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *ProxySvcCert) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *ProxySvcCert) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *ProxySvcCert) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *ProxySvcCert) GetSignatureAlgorithm() string {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *ProxySvcCert) HasSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.SignatureAlgorithm) {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *ProxySvcCert) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ProxySvcCert) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ProxySvcCert) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ProxySvcCert) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o ProxySvcCert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxySvcCert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cert"] = o.Cert
	if !IsNil(o.CommonName) {
		toSerialize["commonName"] = o.CommonName
	}
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.DnsNames) {
		toSerialize["dnsNames"] = o.DnsNames
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.IsCA) {
		toSerialize["isCA"] = o.IsCA
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.NotAfter) {
		toSerialize["notAfter"] = o.NotAfter
	}
	if !IsNil(o.NotBefore) {
		toSerialize["notBefore"] = o.NotBefore
	}
	if !IsNil(o.PublicKeyAlgorithm) {
		toSerialize["publicKeyAlgorithm"] = o.PublicKeyAlgorithm
	}
	if !IsNil(o.PublicKeyBitLength) {
		toSerialize["publicKeyBitLength"] = o.PublicKeyBitLength
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !IsNil(o.SignatureAlgorithm) {
		toSerialize["signatureAlgorithm"] = o.SignatureAlgorithm
	}
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ProxySvcCert) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cert",
		"createdAt",
		"id",
		"updatedAt",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProxySvcCert := _ProxySvcCert{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProxySvcCert)

	if err != nil {
		return err
	}

	*o = ProxySvcCert(varProxySvcCert)

	return err
}

type NullableProxySvcCert struct {
	value *ProxySvcCert
	isSet bool
}

func (v NullableProxySvcCert) Get() *ProxySvcCert {
	return v.value
}

func (v *NullableProxySvcCert) Set(val *ProxySvcCert) {
	v.value = val
	v.isSet = true
}

func (v NullableProxySvcCert) IsSet() bool {
	return v.isSet
}

func (v *NullableProxySvcCert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxySvcCert(val *ProxySvcCert) *NullableProxySvcCert {
	return &NullableProxySvcCert{value: val, isSet: true}
}

func (v NullableProxySvcCert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxySvcCert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


