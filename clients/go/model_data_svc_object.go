/*
1Backend

A common backend for your AI applications—microservices-based and built to scale.

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

// checks if the DataSvcObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSvcObject{}

// DataSvcObject struct for DataSvcObject
type DataSvcObject struct {
	// Authors is a list of user ID and organization ID who created the object. The authors field tracks which users or organizations created an entry, helping to prevent spam. If an organization ID is not provided, the currently active organization will be queried from the User Svc.
	Authors []string `json:"authors,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Data map[string]interface{} `json:"data"`
	// Deleters is a list of user IDs and role IDs that can delete the object. `_self` can be used to refer to the caller user's userId and `_org` can be used to refer to the user's currently active organization (if exists).
	Deleters []string `json:"deleters,omitempty"`
	Id *string `json:"id,omitempty"`
	// Readers is a list of user IDs and role IDs that can read the object. `_self` can be used to refer to the caller user's userId and `_org` can be used to refer to the user's currently active organization (if exists).
	Readers []string `json:"readers,omitempty"`
	Table string `json:"table"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// Writers is a list of user IDs and role IDs that can write the object. `_self` can be used to refer to the caller user's userId and `_org` can be used to refer to the user's currently active organization (if exists).
	Writers []string `json:"writers,omitempty"`
}

type _DataSvcObject DataSvcObject

// NewDataSvcObject instantiates a new DataSvcObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSvcObject(data map[string]interface{}, table string) *DataSvcObject {
	this := DataSvcObject{}
	this.Data = data
	this.Table = table
	return &this
}

// NewDataSvcObjectWithDefaults instantiates a new DataSvcObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSvcObjectWithDefaults() *DataSvcObject {
	this := DataSvcObject{}
	return &this
}

// GetAuthors returns the Authors field value if set, zero value otherwise.
func (o *DataSvcObject) GetAuthors() []string {
	if o == nil || IsNil(o.Authors) {
		var ret []string
		return ret
	}
	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcObject) GetAuthorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Authors) {
		return nil, false
	}
	return o.Authors, true
}

// HasAuthors returns a boolean if a field has been set.
func (o *DataSvcObject) HasAuthors() bool {
	if o != nil && !IsNil(o.Authors) {
		return true
	}

	return false
}

// SetAuthors gets a reference to the given []string and assigns it to the Authors field.
func (o *DataSvcObject) SetAuthors(v []string) {
	o.Authors = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DataSvcObject) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcObject) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DataSvcObject) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DataSvcObject) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetData returns the Data field value
func (o *DataSvcObject) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DataSvcObject) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *DataSvcObject) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetDeleters returns the Deleters field value if set, zero value otherwise.
func (o *DataSvcObject) GetDeleters() []string {
	if o == nil || IsNil(o.Deleters) {
		var ret []string
		return ret
	}
	return o.Deleters
}

// GetDeletersOk returns a tuple with the Deleters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcObject) GetDeletersOk() ([]string, bool) {
	if o == nil || IsNil(o.Deleters) {
		return nil, false
	}
	return o.Deleters, true
}

// HasDeleters returns a boolean if a field has been set.
func (o *DataSvcObject) HasDeleters() bool {
	if o != nil && !IsNil(o.Deleters) {
		return true
	}

	return false
}

// SetDeleters gets a reference to the given []string and assigns it to the Deleters field.
func (o *DataSvcObject) SetDeleters(v []string) {
	o.Deleters = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DataSvcObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DataSvcObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DataSvcObject) SetId(v string) {
	o.Id = &v
}

// GetReaders returns the Readers field value if set, zero value otherwise.
func (o *DataSvcObject) GetReaders() []string {
	if o == nil || IsNil(o.Readers) {
		var ret []string
		return ret
	}
	return o.Readers
}

// GetReadersOk returns a tuple with the Readers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcObject) GetReadersOk() ([]string, bool) {
	if o == nil || IsNil(o.Readers) {
		return nil, false
	}
	return o.Readers, true
}

// HasReaders returns a boolean if a field has been set.
func (o *DataSvcObject) HasReaders() bool {
	if o != nil && !IsNil(o.Readers) {
		return true
	}

	return false
}

// SetReaders gets a reference to the given []string and assigns it to the Readers field.
func (o *DataSvcObject) SetReaders(v []string) {
	o.Readers = v
}

// GetTable returns the Table field value
func (o *DataSvcObject) GetTable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Table
}

// GetTableOk returns a tuple with the Table field value
// and a boolean to check if the value has been set.
func (o *DataSvcObject) GetTableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Table, true
}

// SetTable sets field value
func (o *DataSvcObject) SetTable(v string) {
	o.Table = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DataSvcObject) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcObject) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DataSvcObject) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *DataSvcObject) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetWriters returns the Writers field value if set, zero value otherwise.
func (o *DataSvcObject) GetWriters() []string {
	if o == nil || IsNil(o.Writers) {
		var ret []string
		return ret
	}
	return o.Writers
}

// GetWritersOk returns a tuple with the Writers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcObject) GetWritersOk() ([]string, bool) {
	if o == nil || IsNil(o.Writers) {
		return nil, false
	}
	return o.Writers, true
}

// HasWriters returns a boolean if a field has been set.
func (o *DataSvcObject) HasWriters() bool {
	if o != nil && !IsNil(o.Writers) {
		return true
	}

	return false
}

// SetWriters gets a reference to the given []string and assigns it to the Writers field.
func (o *DataSvcObject) SetWriters(v []string) {
	o.Writers = v
}

func (o DataSvcObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSvcObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authors) {
		toSerialize["authors"] = o.Authors
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	toSerialize["data"] = o.Data
	if !IsNil(o.Deleters) {
		toSerialize["deleters"] = o.Deleters
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Readers) {
		toSerialize["readers"] = o.Readers
	}
	toSerialize["table"] = o.Table
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Writers) {
		toSerialize["writers"] = o.Writers
	}
	return toSerialize, nil
}

func (o *DataSvcObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"table",
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

	varDataSvcObject := _DataSvcObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataSvcObject)

	if err != nil {
		return err
	}

	*o = DataSvcObject(varDataSvcObject)

	return err
}

type NullableDataSvcObject struct {
	value *DataSvcObject
	isSet bool
}

func (v NullableDataSvcObject) Get() *DataSvcObject {
	return v.value
}

func (v *NullableDataSvcObject) Set(val *DataSvcObject) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSvcObject) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSvcObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSvcObject(val *DataSvcObject) *NullableDataSvcObject {
	return &NullableDataSvcObject{value: val, isSet: true}
}

func (v NullableDataSvcObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSvcObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


