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

// checks if the PromptSvcPrompt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptSvcPrompt{}

// PromptSvcPrompt struct for PromptSvcPrompt
type PromptSvcPrompt struct {
	// CreatedAt is the time of the prompt creation.
	CreatedAt *string `json:"createdAt,omitempty"`
	// AI engine/platform (eg. LlamaCpp, Stable Diffusion) specific parameters
	EngineParameters *PromptSvcEngineParameters `json:"engineParameters,omitempty"`
	// Error that arose during prompt execution, if any.
	Error *string `json:"error,omitempty"`
	// Id is the unique ID of the prompt.
	Id *string `json:"id,omitempty"`
	// LastRun is the time of the last prompt run.
	LastRun *string `json:"lastRun,omitempty"`
	// MaxRetries specified how many times the system should retry a prompt when it keeps erroring.
	MaxRetries *int32 `json:"maxRetries,omitempty"`
	// ModelId is just the 1Backend internal ID of the model.
	ModelId *string `json:"modelId,omitempty"`
	// AI engine/platform (eg. LlamaCpp, Stable Diffusion) agnostic parameters. Use these high level parameters when you don't care about the actual engine, only the functionality (eg. text to image, image to image) it provides.
	Parameters *PromptSvcParameters `json:"parameters,omitempty"`
	// Prompt is the message itself eg. \"What's a banana?
	Prompt string `json:"prompt"`
	RequestMessageId *string `json:"requestMessageId,omitempty"`
	ResponseMessageId *string `json:"responseMessageId,omitempty"`
	// RunCount is the number of times the prompt was retried due to errors
	RunCount *int32 `json:"runCount,omitempty"`
	// Status of the prompt.
	Status *PromptSvcPromptStatus `json:"status,omitempty"`
	// Sync drives whether prompt add request should wait and hang until the prompt is done executing. By default the prompt just gets put on a queue and the client will just subscribe to a Thread Stream. For quick and dirty scripting however it's often times easier to do things syncronously. In those cases set Sync to true.
	Sync *bool `json:"sync,omitempty"`
	// ThreadId is the ID of the thread a prompt belongs to. Clients subscribe to Thread Streams to see the answer to a prompt, or set `prompt.sync` to true for a blocking answer.
	ThreadId *string `json:"threadId,omitempty"`
	// Type is inferred from the `Parameters` or `EngineParameters` field. Eg. A LLamaCpp prompt will be \"Text-to-Text\", a Stabel Diffusion one will be \"Text-to-Image\" etc.
	Type *PromptSvcPromptType `json:"type,omitempty"`
	// UpdatedAt is the last time the prompt was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// UserId contains the ID of the user who submitted the prompt.
	UserId *string `json:"userId,omitempty"`
}

type _PromptSvcPrompt PromptSvcPrompt

// NewPromptSvcPrompt instantiates a new PromptSvcPrompt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptSvcPrompt(prompt string) *PromptSvcPrompt {
	this := PromptSvcPrompt{}
	this.Prompt = prompt
	return &this
}

// NewPromptSvcPromptWithDefaults instantiates a new PromptSvcPrompt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptSvcPromptWithDefaults() *PromptSvcPrompt {
	this := PromptSvcPrompt{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PromptSvcPrompt) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetEngineParameters returns the EngineParameters field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetEngineParameters() PromptSvcEngineParameters {
	if o == nil || IsNil(o.EngineParameters) {
		var ret PromptSvcEngineParameters
		return ret
	}
	return *o.EngineParameters
}

// GetEngineParametersOk returns a tuple with the EngineParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetEngineParametersOk() (*PromptSvcEngineParameters, bool) {
	if o == nil || IsNil(o.EngineParameters) {
		return nil, false
	}
	return o.EngineParameters, true
}

// HasEngineParameters returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasEngineParameters() bool {
	if o != nil && !IsNil(o.EngineParameters) {
		return true
	}

	return false
}

// SetEngineParameters gets a reference to the given PromptSvcEngineParameters and assigns it to the EngineParameters field.
func (o *PromptSvcPrompt) SetEngineParameters(v PromptSvcEngineParameters) {
	o.EngineParameters = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *PromptSvcPrompt) SetError(v string) {
	o.Error = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PromptSvcPrompt) SetId(v string) {
	o.Id = &v
}

// GetLastRun returns the LastRun field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetLastRun() string {
	if o == nil || IsNil(o.LastRun) {
		var ret string
		return ret
	}
	return *o.LastRun
}

// GetLastRunOk returns a tuple with the LastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetLastRunOk() (*string, bool) {
	if o == nil || IsNil(o.LastRun) {
		return nil, false
	}
	return o.LastRun, true
}

// HasLastRun returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasLastRun() bool {
	if o != nil && !IsNil(o.LastRun) {
		return true
	}

	return false
}

// SetLastRun gets a reference to the given string and assigns it to the LastRun field.
func (o *PromptSvcPrompt) SetLastRun(v string) {
	o.LastRun = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetMaxRetries() int32 {
	if o == nil || IsNil(o.MaxRetries) {
		var ret int32
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetMaxRetriesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRetries) {
		return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasMaxRetries() bool {
	if o != nil && !IsNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int32 and assigns it to the MaxRetries field.
func (o *PromptSvcPrompt) SetMaxRetries(v int32) {
	o.MaxRetries = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetModelId() string {
	if o == nil || IsNil(o.ModelId) {
		var ret string
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetModelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ModelId) {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasModelId() bool {
	if o != nil && !IsNil(o.ModelId) {
		return true
	}

	return false
}

// SetModelId gets a reference to the given string and assigns it to the ModelId field.
func (o *PromptSvcPrompt) SetModelId(v string) {
	o.ModelId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetParameters() PromptSvcParameters {
	if o == nil || IsNil(o.Parameters) {
		var ret PromptSvcParameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetParametersOk() (*PromptSvcParameters, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given PromptSvcParameters and assigns it to the Parameters field.
func (o *PromptSvcPrompt) SetParameters(v PromptSvcParameters) {
	o.Parameters = &v
}

// GetPrompt returns the Prompt field value
func (o *PromptSvcPrompt) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *PromptSvcPrompt) SetPrompt(v string) {
	o.Prompt = v
}

// GetRequestMessageId returns the RequestMessageId field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetRequestMessageId() string {
	if o == nil || IsNil(o.RequestMessageId) {
		var ret string
		return ret
	}
	return *o.RequestMessageId
}

// GetRequestMessageIdOk returns a tuple with the RequestMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetRequestMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestMessageId) {
		return nil, false
	}
	return o.RequestMessageId, true
}

// HasRequestMessageId returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasRequestMessageId() bool {
	if o != nil && !IsNil(o.RequestMessageId) {
		return true
	}

	return false
}

// SetRequestMessageId gets a reference to the given string and assigns it to the RequestMessageId field.
func (o *PromptSvcPrompt) SetRequestMessageId(v string) {
	o.RequestMessageId = &v
}

// GetResponseMessageId returns the ResponseMessageId field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetResponseMessageId() string {
	if o == nil || IsNil(o.ResponseMessageId) {
		var ret string
		return ret
	}
	return *o.ResponseMessageId
}

// GetResponseMessageIdOk returns a tuple with the ResponseMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetResponseMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseMessageId) {
		return nil, false
	}
	return o.ResponseMessageId, true
}

// HasResponseMessageId returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasResponseMessageId() bool {
	if o != nil && !IsNil(o.ResponseMessageId) {
		return true
	}

	return false
}

// SetResponseMessageId gets a reference to the given string and assigns it to the ResponseMessageId field.
func (o *PromptSvcPrompt) SetResponseMessageId(v string) {
	o.ResponseMessageId = &v
}

// GetRunCount returns the RunCount field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetRunCount() int32 {
	if o == nil || IsNil(o.RunCount) {
		var ret int32
		return ret
	}
	return *o.RunCount
}

// GetRunCountOk returns a tuple with the RunCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetRunCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RunCount) {
		return nil, false
	}
	return o.RunCount, true
}

// HasRunCount returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasRunCount() bool {
	if o != nil && !IsNil(o.RunCount) {
		return true
	}

	return false
}

// SetRunCount gets a reference to the given int32 and assigns it to the RunCount field.
func (o *PromptSvcPrompt) SetRunCount(v int32) {
	o.RunCount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetStatus() PromptSvcPromptStatus {
	if o == nil || IsNil(o.Status) {
		var ret PromptSvcPromptStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetStatusOk() (*PromptSvcPromptStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PromptSvcPromptStatus and assigns it to the Status field.
func (o *PromptSvcPrompt) SetStatus(v PromptSvcPromptStatus) {
	o.Status = &v
}

// GetSync returns the Sync field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetSync() bool {
	if o == nil || IsNil(o.Sync) {
		var ret bool
		return ret
	}
	return *o.Sync
}

// GetSyncOk returns a tuple with the Sync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.Sync) {
		return nil, false
	}
	return o.Sync, true
}

// HasSync returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasSync() bool {
	if o != nil && !IsNil(o.Sync) {
		return true
	}

	return false
}

// SetSync gets a reference to the given bool and assigns it to the Sync field.
func (o *PromptSvcPrompt) SetSync(v bool) {
	o.Sync = &v
}

// GetThreadId returns the ThreadId field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetThreadId() string {
	if o == nil || IsNil(o.ThreadId) {
		var ret string
		return ret
	}
	return *o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetThreadIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThreadId) {
		return nil, false
	}
	return o.ThreadId, true
}

// HasThreadId returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasThreadId() bool {
	if o != nil && !IsNil(o.ThreadId) {
		return true
	}

	return false
}

// SetThreadId gets a reference to the given string and assigns it to the ThreadId field.
func (o *PromptSvcPrompt) SetThreadId(v string) {
	o.ThreadId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetType() PromptSvcPromptType {
	if o == nil || IsNil(o.Type) {
		var ret PromptSvcPromptType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetTypeOk() (*PromptSvcPromptType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PromptSvcPromptType and assigns it to the Type field.
func (o *PromptSvcPrompt) SetType(v PromptSvcPromptType) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *PromptSvcPrompt) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PromptSvcPrompt) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPrompt) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PromptSvcPrompt) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PromptSvcPrompt) SetUserId(v string) {
	o.UserId = &v
}

func (o PromptSvcPrompt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptSvcPrompt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.EngineParameters) {
		toSerialize["engineParameters"] = o.EngineParameters
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastRun) {
		toSerialize["lastRun"] = o.LastRun
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
	if !IsNil(o.RequestMessageId) {
		toSerialize["requestMessageId"] = o.RequestMessageId
	}
	if !IsNil(o.ResponseMessageId) {
		toSerialize["responseMessageId"] = o.ResponseMessageId
	}
	if !IsNil(o.RunCount) {
		toSerialize["runCount"] = o.RunCount
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Sync) {
		toSerialize["sync"] = o.Sync
	}
	if !IsNil(o.ThreadId) {
		toSerialize["threadId"] = o.ThreadId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

func (o *PromptSvcPrompt) UnmarshalJSON(data []byte) (err error) {
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

	varPromptSvcPrompt := _PromptSvcPrompt{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPromptSvcPrompt)

	if err != nil {
		return err
	}

	*o = PromptSvcPrompt(varPromptSvcPrompt)

	return err
}

type NullablePromptSvcPrompt struct {
	value *PromptSvcPrompt
	isSet bool
}

func (v NullablePromptSvcPrompt) Get() *PromptSvcPrompt {
	return v.value
}

func (v *NullablePromptSvcPrompt) Set(val *PromptSvcPrompt) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcPrompt) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcPrompt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcPrompt(val *PromptSvcPrompt) *NullablePromptSvcPrompt {
	return &NullablePromptSvcPrompt{value: val, isSet: true}
}

func (v NullablePromptSvcPrompt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcPrompt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


