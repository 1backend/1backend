/*
1Backend

AI-native microservices platform.

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

// checks if the PromptSvcPromptRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptSvcPromptRequest{}

// PromptSvcPromptRequest struct for PromptSvcPromptRequest
type PromptSvcPromptRequest struct {
	// AI engine/platform (eg. Llama, Stable Diffusion) specific parameters
	EngineParameters *PromptSvcEngineParameters `json:"engineParameters,omitempty"`
	// Id is the unique ID of the prompt.
	Id *string `json:"id,omitempty"`
	// MaxRetries specified how many times the system should retry a prompt when it keeps erroring.
	MaxRetries *int32 `json:"maxRetries,omitempty"`
	// ModelId is just the 1Backend internal ID of the model.
	ModelId *string `json:"modelId,omitempty"`
	// AI engine/platform (eg. Llama, Stable Diffusion) agnostic parameters. Use these high level parameters when you don't care about the actual engine, only the functionality (eg. text to image, image to image) it provides.
	Parameters *PromptSvcParameters `json:"parameters,omitempty"`
	// Prompt is the message itself eg. \"What's a banana?
	Prompt string `json:"prompt"`
	// Sync drives whether prompt add request should wait and hang until the prompt is done executing. By default the prompt just gets put on a queue and the client will just subscribe to a Thread Stream. For quick and dirty scripting however it's often times easier to do things synchronously. In those cases set Sync to true.
	Sync *bool `json:"sync,omitempty"`
	// ThreadId is the ID of the thread a prompt belongs to. Clients subscribe to Thread Streams to see the answer to a prompt, or set `prompt.sync` to true for a blocking answer.
	ThreadId *string `json:"threadId,omitempty"`
}

type _PromptSvcPromptRequest PromptSvcPromptRequest

// NewPromptSvcPromptRequest instantiates a new PromptSvcPromptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptSvcPromptRequest(prompt string) *PromptSvcPromptRequest {
	this := PromptSvcPromptRequest{}
	this.Prompt = prompt
	return &this
}

// NewPromptSvcPromptRequestWithDefaults instantiates a new PromptSvcPromptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptSvcPromptRequestWithDefaults() *PromptSvcPromptRequest {
	this := PromptSvcPromptRequest{}
	return &this
}

// GetEngineParameters returns the EngineParameters field value if set, zero value otherwise.
func (o *PromptSvcPromptRequest) GetEngineParameters() PromptSvcEngineParameters {
	if o == nil || IsNil(o.EngineParameters) {
		var ret PromptSvcEngineParameters
		return ret
	}
	return *o.EngineParameters
}

// GetEngineParametersOk returns a tuple with the EngineParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPromptRequest) GetEngineParametersOk() (*PromptSvcEngineParameters, bool) {
	if o == nil || IsNil(o.EngineParameters) {
		return nil, false
	}
	return o.EngineParameters, true
}

// HasEngineParameters returns a boolean if a field has been set.
func (o *PromptSvcPromptRequest) HasEngineParameters() bool {
	if o != nil && !IsNil(o.EngineParameters) {
		return true
	}

	return false
}

// SetEngineParameters gets a reference to the given PromptSvcEngineParameters and assigns it to the EngineParameters field.
func (o *PromptSvcPromptRequest) SetEngineParameters(v PromptSvcEngineParameters) {
	o.EngineParameters = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PromptSvcPromptRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPromptRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PromptSvcPromptRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PromptSvcPromptRequest) SetId(v string) {
	o.Id = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *PromptSvcPromptRequest) GetMaxRetries() int32 {
	if o == nil || IsNil(o.MaxRetries) {
		var ret int32
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPromptRequest) GetMaxRetriesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRetries) {
		return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *PromptSvcPromptRequest) HasMaxRetries() bool {
	if o != nil && !IsNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int32 and assigns it to the MaxRetries field.
func (o *PromptSvcPromptRequest) SetMaxRetries(v int32) {
	o.MaxRetries = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *PromptSvcPromptRequest) GetModelId() string {
	if o == nil || IsNil(o.ModelId) {
		var ret string
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPromptRequest) GetModelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ModelId) {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *PromptSvcPromptRequest) HasModelId() bool {
	if o != nil && !IsNil(o.ModelId) {
		return true
	}

	return false
}

// SetModelId gets a reference to the given string and assigns it to the ModelId field.
func (o *PromptSvcPromptRequest) SetModelId(v string) {
	o.ModelId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *PromptSvcPromptRequest) GetParameters() PromptSvcParameters {
	if o == nil || IsNil(o.Parameters) {
		var ret PromptSvcParameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPromptRequest) GetParametersOk() (*PromptSvcParameters, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *PromptSvcPromptRequest) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given PromptSvcParameters and assigns it to the Parameters field.
func (o *PromptSvcPromptRequest) SetParameters(v PromptSvcParameters) {
	o.Parameters = &v
}

// GetPrompt returns the Prompt field value
func (o *PromptSvcPromptRequest) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *PromptSvcPromptRequest) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *PromptSvcPromptRequest) SetPrompt(v string) {
	o.Prompt = v
}

// GetSync returns the Sync field value if set, zero value otherwise.
func (o *PromptSvcPromptRequest) GetSync() bool {
	if o == nil || IsNil(o.Sync) {
		var ret bool
		return ret
	}
	return *o.Sync
}

// GetSyncOk returns a tuple with the Sync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPromptRequest) GetSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.Sync) {
		return nil, false
	}
	return o.Sync, true
}

// HasSync returns a boolean if a field has been set.
func (o *PromptSvcPromptRequest) HasSync() bool {
	if o != nil && !IsNil(o.Sync) {
		return true
	}

	return false
}

// SetSync gets a reference to the given bool and assigns it to the Sync field.
func (o *PromptSvcPromptRequest) SetSync(v bool) {
	o.Sync = &v
}

// GetThreadId returns the ThreadId field value if set, zero value otherwise.
func (o *PromptSvcPromptRequest) GetThreadId() string {
	if o == nil || IsNil(o.ThreadId) {
		var ret string
		return ret
	}
	return *o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPromptRequest) GetThreadIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThreadId) {
		return nil, false
	}
	return o.ThreadId, true
}

// HasThreadId returns a boolean if a field has been set.
func (o *PromptSvcPromptRequest) HasThreadId() bool {
	if o != nil && !IsNil(o.ThreadId) {
		return true
	}

	return false
}

// SetThreadId gets a reference to the given string and assigns it to the ThreadId field.
func (o *PromptSvcPromptRequest) SetThreadId(v string) {
	o.ThreadId = &v
}

func (o PromptSvcPromptRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptSvcPromptRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EngineParameters) {
		toSerialize["engineParameters"] = o.EngineParameters
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MaxRetries) {
		toSerialize["maxRetries"] = o.MaxRetries
	}
	if !IsNil(o.ModelId) {
		toSerialize["modelId"] = o.ModelId
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	toSerialize["prompt"] = o.Prompt
	if !IsNil(o.Sync) {
		toSerialize["sync"] = o.Sync
	}
	if !IsNil(o.ThreadId) {
		toSerialize["threadId"] = o.ThreadId
	}
	return toSerialize, nil
}

func (o *PromptSvcPromptRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prompt",
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

	varPromptSvcPromptRequest := _PromptSvcPromptRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPromptSvcPromptRequest)

	if err != nil {
		return err
	}

	*o = PromptSvcPromptRequest(varPromptSvcPromptRequest)

	return err
}

type NullablePromptSvcPromptRequest struct {
	value *PromptSvcPromptRequest
	isSet bool
}

func (v NullablePromptSvcPromptRequest) Get() *PromptSvcPromptRequest {
	return v.value
}

func (v *NullablePromptSvcPromptRequest) Set(val *PromptSvcPromptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcPromptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcPromptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcPromptRequest(val *PromptSvcPromptRequest) *NullablePromptSvcPromptRequest {
	return &NullablePromptSvcPromptRequest{value: val, isSet: true}
}

func (v NullablePromptSvcPromptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcPromptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


