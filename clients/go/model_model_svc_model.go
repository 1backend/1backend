/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.30
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelSvcModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSvcModel{}

// ModelSvcModel struct for ModelSvcModel
type ModelSvcModel struct {
	Assets []ModelSvcAsset `json:"assets,omitempty"`
	Bits *int32 `json:"bits,omitempty"`
	Description *string `json:"description,omitempty"`
	Extension *string `json:"extension,omitempty"`
	Flavour *string `json:"flavour,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	Id string `json:"id"`
	MaxBits *int32 `json:"maxBits,omitempty"`
	MaxRam *float32 `json:"maxRam,omitempty"`
	Mirrors []string `json:"mirrors,omitempty"`
	Name string `json:"name"`
	Parameters *string `json:"parameters,omitempty"`
	PlatformId string `json:"platformId"`
	PromptTemplate *string `json:"promptTemplate,omitempty"`
	Quality *string `json:"quality,omitempty"`
	QuantComment *string `json:"quantComment,omitempty"`
	Size *float32 `json:"size,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Uncensored *bool `json:"uncensored,omitempty"`
	Version *string `json:"version,omitempty"`
}

type _ModelSvcModel ModelSvcModel

// NewModelSvcModel instantiates a new ModelSvcModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSvcModel(id string, name string, platformId string) *ModelSvcModel {
	this := ModelSvcModel{}
	this.Id = id
	this.Name = name
	this.PlatformId = platformId
	return &this
}

// NewModelSvcModelWithDefaults instantiates a new ModelSvcModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSvcModelWithDefaults() *ModelSvcModel {
	this := ModelSvcModel{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *ModelSvcModel) GetAssets() []ModelSvcAsset {
	if o == nil || IsNil(o.Assets) {
		var ret []ModelSvcAsset
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetAssetsOk() ([]ModelSvcAsset, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *ModelSvcModel) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []ModelSvcAsset and assigns it to the Assets field.
func (o *ModelSvcModel) SetAssets(v []ModelSvcAsset) {
	o.Assets = v
}

// GetBits returns the Bits field value if set, zero value otherwise.
func (o *ModelSvcModel) GetBits() int32 {
	if o == nil || IsNil(o.Bits) {
		var ret int32
		return ret
	}
	return *o.Bits
}

// GetBitsOk returns a tuple with the Bits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetBitsOk() (*int32, bool) {
	if o == nil || IsNil(o.Bits) {
		return nil, false
	}
	return o.Bits, true
}

// HasBits returns a boolean if a field has been set.
func (o *ModelSvcModel) HasBits() bool {
	if o != nil && !IsNil(o.Bits) {
		return true
	}

	return false
}

// SetBits gets a reference to the given int32 and assigns it to the Bits field.
func (o *ModelSvcModel) SetBits(v int32) {
	o.Bits = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelSvcModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelSvcModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelSvcModel) SetDescription(v string) {
	o.Description = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *ModelSvcModel) GetExtension() string {
	if o == nil || IsNil(o.Extension) {
		var ret string
		return ret
	}
	return *o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *ModelSvcModel) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given string and assigns it to the Extension field.
func (o *ModelSvcModel) SetExtension(v string) {
	o.Extension = &v
}

// GetFlavour returns the Flavour field value if set, zero value otherwise.
func (o *ModelSvcModel) GetFlavour() string {
	if o == nil || IsNil(o.Flavour) {
		var ret string
		return ret
	}
	return *o.Flavour
}

// GetFlavourOk returns a tuple with the Flavour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetFlavourOk() (*string, bool) {
	if o == nil || IsNil(o.Flavour) {
		return nil, false
	}
	return o.Flavour, true
}

// HasFlavour returns a boolean if a field has been set.
func (o *ModelSvcModel) HasFlavour() bool {
	if o != nil && !IsNil(o.Flavour) {
		return true
	}

	return false
}

// SetFlavour gets a reference to the given string and assigns it to the Flavour field.
func (o *ModelSvcModel) SetFlavour(v string) {
	o.Flavour = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *ModelSvcModel) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *ModelSvcModel) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *ModelSvcModel) SetFullName(v string) {
	o.FullName = &v
}

// GetId returns the Id field value
func (o *ModelSvcModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelSvcModel) SetId(v string) {
	o.Id = v
}

// GetMaxBits returns the MaxBits field value if set, zero value otherwise.
func (o *ModelSvcModel) GetMaxBits() int32 {
	if o == nil || IsNil(o.MaxBits) {
		var ret int32
		return ret
	}
	return *o.MaxBits
}

// GetMaxBitsOk returns a tuple with the MaxBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetMaxBitsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxBits) {
		return nil, false
	}
	return o.MaxBits, true
}

// HasMaxBits returns a boolean if a field has been set.
func (o *ModelSvcModel) HasMaxBits() bool {
	if o != nil && !IsNil(o.MaxBits) {
		return true
	}

	return false
}

// SetMaxBits gets a reference to the given int32 and assigns it to the MaxBits field.
func (o *ModelSvcModel) SetMaxBits(v int32) {
	o.MaxBits = &v
}

// GetMaxRam returns the MaxRam field value if set, zero value otherwise.
func (o *ModelSvcModel) GetMaxRam() float32 {
	if o == nil || IsNil(o.MaxRam) {
		var ret float32
		return ret
	}
	return *o.MaxRam
}

// GetMaxRamOk returns a tuple with the MaxRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetMaxRamOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxRam) {
		return nil, false
	}
	return o.MaxRam, true
}

// HasMaxRam returns a boolean if a field has been set.
func (o *ModelSvcModel) HasMaxRam() bool {
	if o != nil && !IsNil(o.MaxRam) {
		return true
	}

	return false
}

// SetMaxRam gets a reference to the given float32 and assigns it to the MaxRam field.
func (o *ModelSvcModel) SetMaxRam(v float32) {
	o.MaxRam = &v
}

// GetMirrors returns the Mirrors field value if set, zero value otherwise.
func (o *ModelSvcModel) GetMirrors() []string {
	if o == nil || IsNil(o.Mirrors) {
		var ret []string
		return ret
	}
	return o.Mirrors
}

// GetMirrorsOk returns a tuple with the Mirrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetMirrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Mirrors) {
		return nil, false
	}
	return o.Mirrors, true
}

// HasMirrors returns a boolean if a field has been set.
func (o *ModelSvcModel) HasMirrors() bool {
	if o != nil && !IsNil(o.Mirrors) {
		return true
	}

	return false
}

// SetMirrors gets a reference to the given []string and assigns it to the Mirrors field.
func (o *ModelSvcModel) SetMirrors(v []string) {
	o.Mirrors = v
}

// GetName returns the Name field value
func (o *ModelSvcModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelSvcModel) SetName(v string) {
	o.Name = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ModelSvcModel) GetParameters() string {
	if o == nil || IsNil(o.Parameters) {
		var ret string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetParametersOk() (*string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ModelSvcModel) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given string and assigns it to the Parameters field.
func (o *ModelSvcModel) SetParameters(v string) {
	o.Parameters = &v
}

// GetPlatformId returns the PlatformId field value
func (o *ModelSvcModel) GetPlatformId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformId
}

// GetPlatformIdOk returns a tuple with the PlatformId field value
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetPlatformIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlatformId, true
}

// SetPlatformId sets field value
func (o *ModelSvcModel) SetPlatformId(v string) {
	o.PlatformId = v
}

// GetPromptTemplate returns the PromptTemplate field value if set, zero value otherwise.
func (o *ModelSvcModel) GetPromptTemplate() string {
	if o == nil || IsNil(o.PromptTemplate) {
		var ret string
		return ret
	}
	return *o.PromptTemplate
}

// GetPromptTemplateOk returns a tuple with the PromptTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetPromptTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.PromptTemplate) {
		return nil, false
	}
	return o.PromptTemplate, true
}

// HasPromptTemplate returns a boolean if a field has been set.
func (o *ModelSvcModel) HasPromptTemplate() bool {
	if o != nil && !IsNil(o.PromptTemplate) {
		return true
	}

	return false
}

// SetPromptTemplate gets a reference to the given string and assigns it to the PromptTemplate field.
func (o *ModelSvcModel) SetPromptTemplate(v string) {
	o.PromptTemplate = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *ModelSvcModel) GetQuality() string {
	if o == nil || IsNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetQualityOk() (*string, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *ModelSvcModel) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *ModelSvcModel) SetQuality(v string) {
	o.Quality = &v
}

// GetQuantComment returns the QuantComment field value if set, zero value otherwise.
func (o *ModelSvcModel) GetQuantComment() string {
	if o == nil || IsNil(o.QuantComment) {
		var ret string
		return ret
	}
	return *o.QuantComment
}

// GetQuantCommentOk returns a tuple with the QuantComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetQuantCommentOk() (*string, bool) {
	if o == nil || IsNil(o.QuantComment) {
		return nil, false
	}
	return o.QuantComment, true
}

// HasQuantComment returns a boolean if a field has been set.
func (o *ModelSvcModel) HasQuantComment() bool {
	if o != nil && !IsNil(o.QuantComment) {
		return true
	}

	return false
}

// SetQuantComment gets a reference to the given string and assigns it to the QuantComment field.
func (o *ModelSvcModel) SetQuantComment(v string) {
	o.QuantComment = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ModelSvcModel) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ModelSvcModel) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *ModelSvcModel) SetSize(v float32) {
	o.Size = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ModelSvcModel) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ModelSvcModel) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ModelSvcModel) SetTags(v []string) {
	o.Tags = v
}

// GetUncensored returns the Uncensored field value if set, zero value otherwise.
func (o *ModelSvcModel) GetUncensored() bool {
	if o == nil || IsNil(o.Uncensored) {
		var ret bool
		return ret
	}
	return *o.Uncensored
}

// GetUncensoredOk returns a tuple with the Uncensored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetUncensoredOk() (*bool, bool) {
	if o == nil || IsNil(o.Uncensored) {
		return nil, false
	}
	return o.Uncensored, true
}

// HasUncensored returns a boolean if a field has been set.
func (o *ModelSvcModel) HasUncensored() bool {
	if o != nil && !IsNil(o.Uncensored) {
		return true
	}

	return false
}

// SetUncensored gets a reference to the given bool and assigns it to the Uncensored field.
func (o *ModelSvcModel) SetUncensored(v bool) {
	o.Uncensored = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ModelSvcModel) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcModel) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ModelSvcModel) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ModelSvcModel) SetVersion(v string) {
	o.Version = &v
}

func (o ModelSvcModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSvcModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !IsNil(o.Bits) {
		toSerialize["bits"] = o.Bits
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !IsNil(o.Flavour) {
		toSerialize["flavour"] = o.Flavour
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.MaxBits) {
		toSerialize["maxBits"] = o.MaxBits
	}
	if !IsNil(o.MaxRam) {
		toSerialize["maxRam"] = o.MaxRam
	}
	if !IsNil(o.Mirrors) {
		toSerialize["mirrors"] = o.Mirrors
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	toSerialize["platformId"] = o.PlatformId
	if !IsNil(o.PromptTemplate) {
		toSerialize["promptTemplate"] = o.PromptTemplate
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if !IsNil(o.QuantComment) {
		toSerialize["quantComment"] = o.QuantComment
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Uncensored) {
		toSerialize["uncensored"] = o.Uncensored
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

func (o *ModelSvcModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"platformId",
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

	varModelSvcModel := _ModelSvcModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelSvcModel)

	if err != nil {
		return err
	}

	*o = ModelSvcModel(varModelSvcModel)

	return err
}

type NullableModelSvcModel struct {
	value *ModelSvcModel
	isSet bool
}

func (v NullableModelSvcModel) Get() *ModelSvcModel {
	return v.value
}

func (v *NullableModelSvcModel) Set(val *ModelSvcModel) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSvcModel) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSvcModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSvcModel(val *ModelSvcModel) *NullableModelSvcModel {
	return &NullableModelSvcModel{value: val, isSet: true}
}

func (v NullableModelSvcModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSvcModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


