/*
1Backend

A unified backend development platform for your AI applications—microservices-based and built to scale.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SourceSvcCheckoutRepoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceSvcCheckoutRepoRequest{}

// SourceSvcCheckoutRepoRequest struct for SourceSvcCheckoutRepoRequest
type SourceSvcCheckoutRepoRequest struct {
	// Password or token for HTTPS auth
	Password *string `json:"password,omitempty"`
	// SSH private key (optional for SSH connection)
	SshKey *string `json:"ssh_key,omitempty"`
	// Password for SSH private key if encrypted (optional)
	SshKeyPwd *string `json:"ssh_key_pwd,omitempty"`
	// Token for HTTPS auth (optional for SSH)
	Token *string `json:"token,omitempty"`
	// Full repository URL (e.g., https://github.com/user/repo)
	Url *string `json:"url,omitempty"`
	// Username for HTTPS or SSH user (optional for SSH)
	Username *string `json:"username,omitempty"`
	// Branch, tag, or commit SHA
	Version *string `json:"version,omitempty"`
}

// NewSourceSvcCheckoutRepoRequest instantiates a new SourceSvcCheckoutRepoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceSvcCheckoutRepoRequest() *SourceSvcCheckoutRepoRequest {
	this := SourceSvcCheckoutRepoRequest{}
	return &this
}

// NewSourceSvcCheckoutRepoRequestWithDefaults instantiates a new SourceSvcCheckoutRepoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceSvcCheckoutRepoRequestWithDefaults() *SourceSvcCheckoutRepoRequest {
	this := SourceSvcCheckoutRepoRequest{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SourceSvcCheckoutRepoRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSvcCheckoutRepoRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SourceSvcCheckoutRepoRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SourceSvcCheckoutRepoRequest) SetPassword(v string) {
	o.Password = &v
}

// GetSshKey returns the SshKey field value if set, zero value otherwise.
func (o *SourceSvcCheckoutRepoRequest) GetSshKey() string {
	if o == nil || IsNil(o.SshKey) {
		var ret string
		return ret
	}
	return *o.SshKey
}

// GetSshKeyOk returns a tuple with the SshKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSvcCheckoutRepoRequest) GetSshKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SshKey) {
		return nil, false
	}
	return o.SshKey, true
}

// HasSshKey returns a boolean if a field has been set.
func (o *SourceSvcCheckoutRepoRequest) HasSshKey() bool {
	if o != nil && !IsNil(o.SshKey) {
		return true
	}

	return false
}

// SetSshKey gets a reference to the given string and assigns it to the SshKey field.
func (o *SourceSvcCheckoutRepoRequest) SetSshKey(v string) {
	o.SshKey = &v
}

// GetSshKeyPwd returns the SshKeyPwd field value if set, zero value otherwise.
func (o *SourceSvcCheckoutRepoRequest) GetSshKeyPwd() string {
	if o == nil || IsNil(o.SshKeyPwd) {
		var ret string
		return ret
	}
	return *o.SshKeyPwd
}

// GetSshKeyPwdOk returns a tuple with the SshKeyPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSvcCheckoutRepoRequest) GetSshKeyPwdOk() (*string, bool) {
	if o == nil || IsNil(o.SshKeyPwd) {
		return nil, false
	}
	return o.SshKeyPwd, true
}

// HasSshKeyPwd returns a boolean if a field has been set.
func (o *SourceSvcCheckoutRepoRequest) HasSshKeyPwd() bool {
	if o != nil && !IsNil(o.SshKeyPwd) {
		return true
	}

	return false
}

// SetSshKeyPwd gets a reference to the given string and assigns it to the SshKeyPwd field.
func (o *SourceSvcCheckoutRepoRequest) SetSshKeyPwd(v string) {
	o.SshKeyPwd = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SourceSvcCheckoutRepoRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSvcCheckoutRepoRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SourceSvcCheckoutRepoRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SourceSvcCheckoutRepoRequest) SetToken(v string) {
	o.Token = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SourceSvcCheckoutRepoRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSvcCheckoutRepoRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SourceSvcCheckoutRepoRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SourceSvcCheckoutRepoRequest) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SourceSvcCheckoutRepoRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSvcCheckoutRepoRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SourceSvcCheckoutRepoRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SourceSvcCheckoutRepoRequest) SetUsername(v string) {
	o.Username = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SourceSvcCheckoutRepoRequest) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSvcCheckoutRepoRequest) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SourceSvcCheckoutRepoRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SourceSvcCheckoutRepoRequest) SetVersion(v string) {
	o.Version = &v
}

func (o SourceSvcCheckoutRepoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceSvcCheckoutRepoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.SshKey) {
		toSerialize["ssh_key"] = o.SshKey
	}
	if !IsNil(o.SshKeyPwd) {
		toSerialize["ssh_key_pwd"] = o.SshKeyPwd
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableSourceSvcCheckoutRepoRequest struct {
	value *SourceSvcCheckoutRepoRequest
	isSet bool
}

func (v NullableSourceSvcCheckoutRepoRequest) Get() *SourceSvcCheckoutRepoRequest {
	return v.value
}

func (v *NullableSourceSvcCheckoutRepoRequest) Set(val *SourceSvcCheckoutRepoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceSvcCheckoutRepoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceSvcCheckoutRepoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceSvcCheckoutRepoRequest(val *SourceSvcCheckoutRepoRequest) *NullableSourceSvcCheckoutRepoRequest {
	return &NullableSourceSvcCheckoutRepoRequest{value: val, isSet: true}
}

func (v NullableSourceSvcCheckoutRepoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceSvcCheckoutRepoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


