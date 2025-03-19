/*
1Backend

A unified backend development platform for microservices-based AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RegistrySvcRegisterInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcRegisterInstanceRequest{}

// RegistrySvcRegisterInstanceRequest struct for RegistrySvcRegisterInstanceRequest
type RegistrySvcRegisterInstanceRequest struct {
	// The ID of the deployment that this instance is an instance of.
	DeploymentId *string `json:"deploymentId,omitempty"`
	// Host of the instance address. Required if URL is not provided
	Host *string `json:"host,omitempty"`
	Id *string `json:"id,omitempty"`
	// IP of the instance address. Optional: to register by IP instead of host
	Ip *string `json:"ip,omitempty"`
	// Path of the instance address. Optional (e.g., \"/api\")
	Path *string `json:"path,omitempty"`
	// Port of the instance address. Required if URL is not provided
	Port *int32 `json:"port,omitempty"`
	// Scheme of the instance address. Required if URL is not provided.
	Scheme *string `json:"scheme,omitempty"`
	// Full address URL of the instance.
	Url string `json:"url"`
}

type _RegistrySvcRegisterInstanceRequest RegistrySvcRegisterInstanceRequest

// NewRegistrySvcRegisterInstanceRequest instantiates a new RegistrySvcRegisterInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcRegisterInstanceRequest(url string) *RegistrySvcRegisterInstanceRequest {
	this := RegistrySvcRegisterInstanceRequest{}
	this.Url = url
	return &this
}

// NewRegistrySvcRegisterInstanceRequestWithDefaults instantiates a new RegistrySvcRegisterInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcRegisterInstanceRequestWithDefaults() *RegistrySvcRegisterInstanceRequest {
	this := RegistrySvcRegisterInstanceRequest{}
	return &this
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *RegistrySvcRegisterInstanceRequest) GetDeploymentId() string {
	if o == nil || IsNil(o.DeploymentId) {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcRegisterInstanceRequest) GetDeploymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentId) {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *RegistrySvcRegisterInstanceRequest) HasDeploymentId() bool {
	if o != nil && !IsNil(o.DeploymentId) {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *RegistrySvcRegisterInstanceRequest) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *RegistrySvcRegisterInstanceRequest) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcRegisterInstanceRequest) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *RegistrySvcRegisterInstanceRequest) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *RegistrySvcRegisterInstanceRequest) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegistrySvcRegisterInstanceRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcRegisterInstanceRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegistrySvcRegisterInstanceRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegistrySvcRegisterInstanceRequest) SetId(v string) {
	o.Id = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *RegistrySvcRegisterInstanceRequest) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcRegisterInstanceRequest) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *RegistrySvcRegisterInstanceRequest) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *RegistrySvcRegisterInstanceRequest) SetIp(v string) {
	o.Ip = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *RegistrySvcRegisterInstanceRequest) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcRegisterInstanceRequest) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *RegistrySvcRegisterInstanceRequest) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *RegistrySvcRegisterInstanceRequest) SetPath(v string) {
	o.Path = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *RegistrySvcRegisterInstanceRequest) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcRegisterInstanceRequest) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *RegistrySvcRegisterInstanceRequest) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *RegistrySvcRegisterInstanceRequest) SetPort(v int32) {
	o.Port = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *RegistrySvcRegisterInstanceRequest) GetScheme() string {
	if o == nil || IsNil(o.Scheme) {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcRegisterInstanceRequest) GetSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *RegistrySvcRegisterInstanceRequest) HasScheme() bool {
	if o != nil && !IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *RegistrySvcRegisterInstanceRequest) SetScheme(v string) {
	o.Scheme = &v
}

// GetUrl returns the Url field value
func (o *RegistrySvcRegisterInstanceRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcRegisterInstanceRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RegistrySvcRegisterInstanceRequest) SetUrl(v string) {
	o.Url = v
}

func (o RegistrySvcRegisterInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcRegisterInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeploymentId) {
		toSerialize["deploymentId"] = o.DeploymentId
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *RegistrySvcRegisterInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varRegistrySvcRegisterInstanceRequest := _RegistrySvcRegisterInstanceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrySvcRegisterInstanceRequest)

	if err != nil {
		return err
	}

	*o = RegistrySvcRegisterInstanceRequest(varRegistrySvcRegisterInstanceRequest)

	return err
}

type NullableRegistrySvcRegisterInstanceRequest struct {
	value *RegistrySvcRegisterInstanceRequest
	isSet bool
}

func (v NullableRegistrySvcRegisterInstanceRequest) Get() *RegistrySvcRegisterInstanceRequest {
	return v.value
}

func (v *NullableRegistrySvcRegisterInstanceRequest) Set(val *RegistrySvcRegisterInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcRegisterInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcRegisterInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcRegisterInstanceRequest(val *RegistrySvcRegisterInstanceRequest) *NullableRegistrySvcRegisterInstanceRequest {
	return &NullableRegistrySvcRegisterInstanceRequest{value: val, isSet: true}
}

func (v NullableRegistrySvcRegisterInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcRegisterInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


