/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.39
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SecretSvcSecret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretSvcSecret{}

// SecretSvcSecret struct for SecretSvcSecret
type SecretSvcSecret struct {
	// Slugs of services/users who can change the deleters list
	CanChangeDeleters []string `json:"canChangeDeleters,omitempty"`
	// Slugs of services/users who can change the readers list
	CanChangeReaders []string `json:"canChangeReaders,omitempty"`
	// Slugs of services/users who can change the writers list
	CanChangeWriters []string `json:"canChangeWriters,omitempty"`
	// Checksum of the secret value
	Checksum *string `json:"checksum,omitempty"`
	// Algorithm used for the checksum (e.g., \"CRC32\")
	ChecksumAlgorithm *SecretSvcChecksumAlgorithm `json:"checksumAlgorithm,omitempty"`
	// Slugs of services/users who can delete the secret
	Deleters []string `json:"deleters,omitempty"`
	// Whether the secret is encrypted All secrets are encrypted before written to the DB. This really only exists for write requests to know if the secret is already encrypted. Ie: while most `secret save [key] [value]` commands are probably not encrypted, File based saves, eg. `secret save secretA.yaml` are probably encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
	// Id of the secret
	Id *string `json:"id,omitempty"`
	// Envar or slug-like key of the secret
	Key *string `json:"key,omitempty"`
	// Namespace of the secret
	Namespace *string `json:"namespace,omitempty"`
	// Slugs of services/users who can read the secret
	Readers []string `json:"readers,omitempty"`
	// Secret Value
	Value *string `json:"value,omitempty"`
	// Slugs of services/users who can modify the secret
	Writers []string `json:"writers,omitempty"`
}

// NewSecretSvcSecret instantiates a new SecretSvcSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretSvcSecret() *SecretSvcSecret {
	this := SecretSvcSecret{}
	return &this
}

// NewSecretSvcSecretWithDefaults instantiates a new SecretSvcSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretSvcSecretWithDefaults() *SecretSvcSecret {
	this := SecretSvcSecret{}
	return &this
}

// GetCanChangeDeleters returns the CanChangeDeleters field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetCanChangeDeleters() []string {
	if o == nil || IsNil(o.CanChangeDeleters) {
		var ret []string
		return ret
	}
	return o.CanChangeDeleters
}

// GetCanChangeDeletersOk returns a tuple with the CanChangeDeleters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetCanChangeDeletersOk() ([]string, bool) {
	if o == nil || IsNil(o.CanChangeDeleters) {
		return nil, false
	}
	return o.CanChangeDeleters, true
}

// HasCanChangeDeleters returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasCanChangeDeleters() bool {
	if o != nil && !IsNil(o.CanChangeDeleters) {
		return true
	}

	return false
}

// SetCanChangeDeleters gets a reference to the given []string and assigns it to the CanChangeDeleters field.
func (o *SecretSvcSecret) SetCanChangeDeleters(v []string) {
	o.CanChangeDeleters = v
}

// GetCanChangeReaders returns the CanChangeReaders field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetCanChangeReaders() []string {
	if o == nil || IsNil(o.CanChangeReaders) {
		var ret []string
		return ret
	}
	return o.CanChangeReaders
}

// GetCanChangeReadersOk returns a tuple with the CanChangeReaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetCanChangeReadersOk() ([]string, bool) {
	if o == nil || IsNil(o.CanChangeReaders) {
		return nil, false
	}
	return o.CanChangeReaders, true
}

// HasCanChangeReaders returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasCanChangeReaders() bool {
	if o != nil && !IsNil(o.CanChangeReaders) {
		return true
	}

	return false
}

// SetCanChangeReaders gets a reference to the given []string and assigns it to the CanChangeReaders field.
func (o *SecretSvcSecret) SetCanChangeReaders(v []string) {
	o.CanChangeReaders = v
}

// GetCanChangeWriters returns the CanChangeWriters field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetCanChangeWriters() []string {
	if o == nil || IsNil(o.CanChangeWriters) {
		var ret []string
		return ret
	}
	return o.CanChangeWriters
}

// GetCanChangeWritersOk returns a tuple with the CanChangeWriters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetCanChangeWritersOk() ([]string, bool) {
	if o == nil || IsNil(o.CanChangeWriters) {
		return nil, false
	}
	return o.CanChangeWriters, true
}

// HasCanChangeWriters returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasCanChangeWriters() bool {
	if o != nil && !IsNil(o.CanChangeWriters) {
		return true
	}

	return false
}

// SetCanChangeWriters gets a reference to the given []string and assigns it to the CanChangeWriters field.
func (o *SecretSvcSecret) SetCanChangeWriters(v []string) {
	o.CanChangeWriters = v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetChecksum() string {
	if o == nil || IsNil(o.Checksum) {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *SecretSvcSecret) SetChecksum(v string) {
	o.Checksum = &v
}

// GetChecksumAlgorithm returns the ChecksumAlgorithm field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetChecksumAlgorithm() SecretSvcChecksumAlgorithm {
	if o == nil || IsNil(o.ChecksumAlgorithm) {
		var ret SecretSvcChecksumAlgorithm
		return ret
	}
	return *o.ChecksumAlgorithm
}

// GetChecksumAlgorithmOk returns a tuple with the ChecksumAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetChecksumAlgorithmOk() (*SecretSvcChecksumAlgorithm, bool) {
	if o == nil || IsNil(o.ChecksumAlgorithm) {
		return nil, false
	}
	return o.ChecksumAlgorithm, true
}

// HasChecksumAlgorithm returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasChecksumAlgorithm() bool {
	if o != nil && !IsNil(o.ChecksumAlgorithm) {
		return true
	}

	return false
}

// SetChecksumAlgorithm gets a reference to the given SecretSvcChecksumAlgorithm and assigns it to the ChecksumAlgorithm field.
func (o *SecretSvcSecret) SetChecksumAlgorithm(v SecretSvcChecksumAlgorithm) {
	o.ChecksumAlgorithm = &v
}

// GetDeleters returns the Deleters field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetDeleters() []string {
	if o == nil || IsNil(o.Deleters) {
		var ret []string
		return ret
	}
	return o.Deleters
}

// GetDeletersOk returns a tuple with the Deleters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetDeletersOk() ([]string, bool) {
	if o == nil || IsNil(o.Deleters) {
		return nil, false
	}
	return o.Deleters, true
}

// HasDeleters returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasDeleters() bool {
	if o != nil && !IsNil(o.Deleters) {
		return true
	}

	return false
}

// SetDeleters gets a reference to the given []string and assigns it to the Deleters field.
func (o *SecretSvcSecret) SetDeleters(v []string) {
	o.Deleters = v
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetEncrypted() bool {
	if o == nil || IsNil(o.Encrypted) {
		var ret bool
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.Encrypted) {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasEncrypted() bool {
	if o != nil && !IsNil(o.Encrypted) {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given bool and assigns it to the Encrypted field.
func (o *SecretSvcSecret) SetEncrypted(v bool) {
	o.Encrypted = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SecretSvcSecret) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SecretSvcSecret) SetKey(v string) {
	o.Key = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *SecretSvcSecret) SetNamespace(v string) {
	o.Namespace = &v
}

// GetReaders returns the Readers field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetReaders() []string {
	if o == nil || IsNil(o.Readers) {
		var ret []string
		return ret
	}
	return o.Readers
}

// GetReadersOk returns a tuple with the Readers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetReadersOk() ([]string, bool) {
	if o == nil || IsNil(o.Readers) {
		return nil, false
	}
	return o.Readers, true
}

// HasReaders returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasReaders() bool {
	if o != nil && !IsNil(o.Readers) {
		return true
	}

	return false
}

// SetReaders gets a reference to the given []string and assigns it to the Readers field.
func (o *SecretSvcSecret) SetReaders(v []string) {
	o.Readers = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SecretSvcSecret) SetValue(v string) {
	o.Value = &v
}

// GetWriters returns the Writers field value if set, zero value otherwise.
func (o *SecretSvcSecret) GetWriters() []string {
	if o == nil || IsNil(o.Writers) {
		var ret []string
		return ret
	}
	return o.Writers
}

// GetWritersOk returns a tuple with the Writers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSecret) GetWritersOk() ([]string, bool) {
	if o == nil || IsNil(o.Writers) {
		return nil, false
	}
	return o.Writers, true
}

// HasWriters returns a boolean if a field has been set.
func (o *SecretSvcSecret) HasWriters() bool {
	if o != nil && !IsNil(o.Writers) {
		return true
	}

	return false
}

// SetWriters gets a reference to the given []string and assigns it to the Writers field.
func (o *SecretSvcSecret) SetWriters(v []string) {
	o.Writers = v
}

func (o SecretSvcSecret) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretSvcSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanChangeDeleters) {
		toSerialize["canChangeDeleters"] = o.CanChangeDeleters
	}
	if !IsNil(o.CanChangeReaders) {
		toSerialize["canChangeReaders"] = o.CanChangeReaders
	}
	if !IsNil(o.CanChangeWriters) {
		toSerialize["canChangeWriters"] = o.CanChangeWriters
	}
	if !IsNil(o.Checksum) {
		toSerialize["checksum"] = o.Checksum
	}
	if !IsNil(o.ChecksumAlgorithm) {
		toSerialize["checksumAlgorithm"] = o.ChecksumAlgorithm
	}
	if !IsNil(o.Deleters) {
		toSerialize["deleters"] = o.Deleters
	}
	if !IsNil(o.Encrypted) {
		toSerialize["encrypted"] = o.Encrypted
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Readers) {
		toSerialize["readers"] = o.Readers
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Writers) {
		toSerialize["writers"] = o.Writers
	}
	return toSerialize, nil
}

type NullableSecretSvcSecret struct {
	value *SecretSvcSecret
	isSet bool
}

func (v NullableSecretSvcSecret) Get() *SecretSvcSecret {
	return v.value
}

func (v *NullableSecretSvcSecret) Set(val *SecretSvcSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretSvcSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretSvcSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretSvcSecret(val *SecretSvcSecret) *NullableSecretSvcSecret {
	return &NullableSecretSvcSecret{value: val, isSet: true}
}

func (v NullableSecretSvcSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretSvcSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


