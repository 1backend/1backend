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
)

// checks if the PromptSvcTextToImageParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptSvcTextToImageParameters{}

// PromptSvcTextToImageParameters struct for PromptSvcTextToImageParameters
type PromptSvcTextToImageParameters struct {
	// Alternative way to specify dimensions (e.g., \"16:9\", \"1:1\").
	AspectRatio *string `json:"aspectRatio,omitempty"`
	// Number of images to generate per batch.
	BatchSize *int32 `json:"batchSize,omitempty"`
	// Controls how much variation is introduced in image modifications.
	DenoisingStrength *float32 `json:"denoisingStrength,omitempty"`
	// Whether to apply AI-based upscaling.
	EnableUpscaling *bool `json:"enableUpscaling,omitempty"`
	// Output format for the generated image (png, jpg, webp, etc.).
	Format *string `json:"format,omitempty"`
	// How closely the output should follow the prompt.
	GuidanceScale *float32 `json:"guidanceScale,omitempty"`
	Height *int32 `json:"height,omitempty"`
	// A negative prompt to specify what should be avoided in the image.
	NegativePrompt *string `json:"negativePrompt,omitempty"`
	// Number of batches to generate.
	NumIterations *int32 `json:"numIterations,omitempty"`
	// The primary prompt for generating the image. Defaults to the top-level prompt if not specified. If both are provided (which should be avoided), this field takes precedence.
	Prompt *string `json:"prompt,omitempty"`
	// Preset quality settings (e.g., Low, Medium, High, Ultra).
	QualityPreset *string `json:"qualityPreset,omitempty"`
	// Whether to enhance facial details for portraits.
	RestoreFaces *bool `json:"restoreFaces,omitempty"`
	// Specifies the sampling method used during generation.
	Scheduler *string `json:"scheduler,omitempty"`
	// Optional seed for reproducibility. If not set, a random seed is used.
	Seed *int32 `json:"seed,omitempty"`
	// Number of inference steps for image generation.
	Steps *int32 `json:"steps,omitempty"`
	// List of artistic styles or themes to apply.
	Styles []string `json:"styles,omitempty"`
	// Image dimensions (width and height in pixels).
	Width *int32 `json:"width,omitempty"`
}

// NewPromptSvcTextToImageParameters instantiates a new PromptSvcTextToImageParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptSvcTextToImageParameters() *PromptSvcTextToImageParameters {
	this := PromptSvcTextToImageParameters{}
	return &this
}

// NewPromptSvcTextToImageParametersWithDefaults instantiates a new PromptSvcTextToImageParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptSvcTextToImageParametersWithDefaults() *PromptSvcTextToImageParameters {
	this := PromptSvcTextToImageParameters{}
	return &this
}

// GetAspectRatio returns the AspectRatio field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetAspectRatio() string {
	if o == nil || IsNil(o.AspectRatio) {
		var ret string
		return ret
	}
	return *o.AspectRatio
}

// GetAspectRatioOk returns a tuple with the AspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetAspectRatioOk() (*string, bool) {
	if o == nil || IsNil(o.AspectRatio) {
		return nil, false
	}
	return o.AspectRatio, true
}

// HasAspectRatio returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasAspectRatio() bool {
	if o != nil && !IsNil(o.AspectRatio) {
		return true
	}

	return false
}

// SetAspectRatio gets a reference to the given string and assigns it to the AspectRatio field.
func (o *PromptSvcTextToImageParameters) SetAspectRatio(v string) {
	o.AspectRatio = &v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetBatchSize() int32 {
	if o == nil || IsNil(o.BatchSize) {
		var ret int32
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetBatchSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.BatchSize) {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasBatchSize() bool {
	if o != nil && !IsNil(o.BatchSize) {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given int32 and assigns it to the BatchSize field.
func (o *PromptSvcTextToImageParameters) SetBatchSize(v int32) {
	o.BatchSize = &v
}

// GetDenoisingStrength returns the DenoisingStrength field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetDenoisingStrength() float32 {
	if o == nil || IsNil(o.DenoisingStrength) {
		var ret float32
		return ret
	}
	return *o.DenoisingStrength
}

// GetDenoisingStrengthOk returns a tuple with the DenoisingStrength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetDenoisingStrengthOk() (*float32, bool) {
	if o == nil || IsNil(o.DenoisingStrength) {
		return nil, false
	}
	return o.DenoisingStrength, true
}

// HasDenoisingStrength returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasDenoisingStrength() bool {
	if o != nil && !IsNil(o.DenoisingStrength) {
		return true
	}

	return false
}

// SetDenoisingStrength gets a reference to the given float32 and assigns it to the DenoisingStrength field.
func (o *PromptSvcTextToImageParameters) SetDenoisingStrength(v float32) {
	o.DenoisingStrength = &v
}

// GetEnableUpscaling returns the EnableUpscaling field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetEnableUpscaling() bool {
	if o == nil || IsNil(o.EnableUpscaling) {
		var ret bool
		return ret
	}
	return *o.EnableUpscaling
}

// GetEnableUpscalingOk returns a tuple with the EnableUpscaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetEnableUpscalingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableUpscaling) {
		return nil, false
	}
	return o.EnableUpscaling, true
}

// HasEnableUpscaling returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasEnableUpscaling() bool {
	if o != nil && !IsNil(o.EnableUpscaling) {
		return true
	}

	return false
}

// SetEnableUpscaling gets a reference to the given bool and assigns it to the EnableUpscaling field.
func (o *PromptSvcTextToImageParameters) SetEnableUpscaling(v bool) {
	o.EnableUpscaling = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *PromptSvcTextToImageParameters) SetFormat(v string) {
	o.Format = &v
}

// GetGuidanceScale returns the GuidanceScale field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetGuidanceScale() float32 {
	if o == nil || IsNil(o.GuidanceScale) {
		var ret float32
		return ret
	}
	return *o.GuidanceScale
}

// GetGuidanceScaleOk returns a tuple with the GuidanceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetGuidanceScaleOk() (*float32, bool) {
	if o == nil || IsNil(o.GuidanceScale) {
		return nil, false
	}
	return o.GuidanceScale, true
}

// HasGuidanceScale returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasGuidanceScale() bool {
	if o != nil && !IsNil(o.GuidanceScale) {
		return true
	}

	return false
}

// SetGuidanceScale gets a reference to the given float32 and assigns it to the GuidanceScale field.
func (o *PromptSvcTextToImageParameters) SetGuidanceScale(v float32) {
	o.GuidanceScale = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *PromptSvcTextToImageParameters) SetHeight(v int32) {
	o.Height = &v
}

// GetNegativePrompt returns the NegativePrompt field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetNegativePrompt() string {
	if o == nil || IsNil(o.NegativePrompt) {
		var ret string
		return ret
	}
	return *o.NegativePrompt
}

// GetNegativePromptOk returns a tuple with the NegativePrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetNegativePromptOk() (*string, bool) {
	if o == nil || IsNil(o.NegativePrompt) {
		return nil, false
	}
	return o.NegativePrompt, true
}

// HasNegativePrompt returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasNegativePrompt() bool {
	if o != nil && !IsNil(o.NegativePrompt) {
		return true
	}

	return false
}

// SetNegativePrompt gets a reference to the given string and assigns it to the NegativePrompt field.
func (o *PromptSvcTextToImageParameters) SetNegativePrompt(v string) {
	o.NegativePrompt = &v
}

// GetNumIterations returns the NumIterations field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetNumIterations() int32 {
	if o == nil || IsNil(o.NumIterations) {
		var ret int32
		return ret
	}
	return *o.NumIterations
}

// GetNumIterationsOk returns a tuple with the NumIterations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetNumIterationsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumIterations) {
		return nil, false
	}
	return o.NumIterations, true
}

// HasNumIterations returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasNumIterations() bool {
	if o != nil && !IsNil(o.NumIterations) {
		return true
	}

	return false
}

// SetNumIterations gets a reference to the given int32 and assigns it to the NumIterations field.
func (o *PromptSvcTextToImageParameters) SetNumIterations(v int32) {
	o.NumIterations = &v
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetPrompt() string {
	if o == nil || IsNil(o.Prompt) {
		var ret string
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetPromptOk() (*string, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasPrompt() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given string and assigns it to the Prompt field.
func (o *PromptSvcTextToImageParameters) SetPrompt(v string) {
	o.Prompt = &v
}

// GetQualityPreset returns the QualityPreset field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetQualityPreset() string {
	if o == nil || IsNil(o.QualityPreset) {
		var ret string
		return ret
	}
	return *o.QualityPreset
}

// GetQualityPresetOk returns a tuple with the QualityPreset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetQualityPresetOk() (*string, bool) {
	if o == nil || IsNil(o.QualityPreset) {
		return nil, false
	}
	return o.QualityPreset, true
}

// HasQualityPreset returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasQualityPreset() bool {
	if o != nil && !IsNil(o.QualityPreset) {
		return true
	}

	return false
}

// SetQualityPreset gets a reference to the given string and assigns it to the QualityPreset field.
func (o *PromptSvcTextToImageParameters) SetQualityPreset(v string) {
	o.QualityPreset = &v
}

// GetRestoreFaces returns the RestoreFaces field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetRestoreFaces() bool {
	if o == nil || IsNil(o.RestoreFaces) {
		var ret bool
		return ret
	}
	return *o.RestoreFaces
}

// GetRestoreFacesOk returns a tuple with the RestoreFaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetRestoreFacesOk() (*bool, bool) {
	if o == nil || IsNil(o.RestoreFaces) {
		return nil, false
	}
	return o.RestoreFaces, true
}

// HasRestoreFaces returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasRestoreFaces() bool {
	if o != nil && !IsNil(o.RestoreFaces) {
		return true
	}

	return false
}

// SetRestoreFaces gets a reference to the given bool and assigns it to the RestoreFaces field.
func (o *PromptSvcTextToImageParameters) SetRestoreFaces(v bool) {
	o.RestoreFaces = &v
}

// GetScheduler returns the Scheduler field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetScheduler() string {
	if o == nil || IsNil(o.Scheduler) {
		var ret string
		return ret
	}
	return *o.Scheduler
}

// GetSchedulerOk returns a tuple with the Scheduler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetSchedulerOk() (*string, bool) {
	if o == nil || IsNil(o.Scheduler) {
		return nil, false
	}
	return o.Scheduler, true
}

// HasScheduler returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasScheduler() bool {
	if o != nil && !IsNil(o.Scheduler) {
		return true
	}

	return false
}

// SetScheduler gets a reference to the given string and assigns it to the Scheduler field.
func (o *PromptSvcTextToImageParameters) SetScheduler(v string) {
	o.Scheduler = &v
}

// GetSeed returns the Seed field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetSeed() int32 {
	if o == nil || IsNil(o.Seed) {
		var ret int32
		return ret
	}
	return *o.Seed
}

// GetSeedOk returns a tuple with the Seed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetSeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Seed) {
		return nil, false
	}
	return o.Seed, true
}

// HasSeed returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasSeed() bool {
	if o != nil && !IsNil(o.Seed) {
		return true
	}

	return false
}

// SetSeed gets a reference to the given int32 and assigns it to the Seed field.
func (o *PromptSvcTextToImageParameters) SetSeed(v int32) {
	o.Seed = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetSteps() int32 {
	if o == nil || IsNil(o.Steps) {
		var ret int32
		return ret
	}
	return *o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetStepsOk() (*int32, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given int32 and assigns it to the Steps field.
func (o *PromptSvcTextToImageParameters) SetSteps(v int32) {
	o.Steps = &v
}

// GetStyles returns the Styles field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetStyles() []string {
	if o == nil || IsNil(o.Styles) {
		var ret []string
		return ret
	}
	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetStylesOk() ([]string, bool) {
	if o == nil || IsNil(o.Styles) {
		return nil, false
	}
	return o.Styles, true
}

// HasStyles returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasStyles() bool {
	if o != nil && !IsNil(o.Styles) {
		return true
	}

	return false
}

// SetStyles gets a reference to the given []string and assigns it to the Styles field.
func (o *PromptSvcTextToImageParameters) SetStyles(v []string) {
	o.Styles = v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *PromptSvcTextToImageParameters) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcTextToImageParameters) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *PromptSvcTextToImageParameters) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *PromptSvcTextToImageParameters) SetWidth(v int32) {
	o.Width = &v
}

func (o PromptSvcTextToImageParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptSvcTextToImageParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AspectRatio) {
		toSerialize["aspectRatio"] = o.AspectRatio
	}
	if !IsNil(o.BatchSize) {
		toSerialize["batchSize"] = o.BatchSize
	}
	if !IsNil(o.DenoisingStrength) {
		toSerialize["denoisingStrength"] = o.DenoisingStrength
	}
	if !IsNil(o.EnableUpscaling) {
		toSerialize["enableUpscaling"] = o.EnableUpscaling
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.GuidanceScale) {
		toSerialize["guidanceScale"] = o.GuidanceScale
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.NegativePrompt) {
		toSerialize["negativePrompt"] = o.NegativePrompt
	}
	if !IsNil(o.NumIterations) {
		toSerialize["numIterations"] = o.NumIterations
	}
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	if !IsNil(o.QualityPreset) {
		toSerialize["qualityPreset"] = o.QualityPreset
	}
	if !IsNil(o.RestoreFaces) {
		toSerialize["restoreFaces"] = o.RestoreFaces
	}
	if !IsNil(o.Scheduler) {
		toSerialize["scheduler"] = o.Scheduler
	}
	if !IsNil(o.Seed) {
		toSerialize["seed"] = o.Seed
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if !IsNil(o.Styles) {
		toSerialize["styles"] = o.Styles
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullablePromptSvcTextToImageParameters struct {
	value *PromptSvcTextToImageParameters
	isSet bool
}

func (v NullablePromptSvcTextToImageParameters) Get() *PromptSvcTextToImageParameters {
	return v.value
}

func (v *NullablePromptSvcTextToImageParameters) Set(val *PromptSvcTextToImageParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcTextToImageParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcTextToImageParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcTextToImageParameters(val *PromptSvcTextToImageParameters) *NullablePromptSvcTextToImageParameters {
	return &NullablePromptSvcTextToImageParameters{value: val, isSet: true}
}

func (v NullablePromptSvcTextToImageParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcTextToImageParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


