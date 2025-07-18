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
	"fmt"
)

// PromptSvcPromptType the model 'PromptSvcPromptType'
type PromptSvcPromptType string

// List of prompt_svc.PromptType
const (
	PromptTypeImageTextToText PromptSvcPromptType = "Image-Text-to-Text"
	PromptTypeVisualQuestionAnswering PromptSvcPromptType = "Visual Question Answering"
	PromptTypeDocumentQuestionAnswering PromptSvcPromptType = "Document Question Answering"
	PromptTypeTextToImage PromptSvcPromptType = "Text-to-Image"
	PromptTypeImageToImage PromptSvcPromptType = "Image-to-Image"
	PromptTypeImageToVideo PromptSvcPromptType = "Image-to-Video"
	PromptTypeUnconditionalImageGeneration PromptSvcPromptType = "Unconditional Image Generation"
	PromptTypeTextToVideo PromptSvcPromptType = "Text-to-Video"
	PromptTypeZeroShotImageClassification PromptSvcPromptType = "Zero-Shot Image Classification"
	PromptTypeZeroShotObjectDetection PromptSvcPromptType = "Zero-Shot Object Detection"
	PromptTypeTextTo3D PromptSvcPromptType = "Text-to-3D"
	PromptTypeImageTo3D PromptSvcPromptType = "Image-to-3D"
	PromptTypeImageFeatureExtraction PromptSvcPromptType = "Image Feature Extraction"
	PromptTypeKeypointDetection PromptSvcPromptType = "Keypoint Detection"
	PromptTypeTextToText PromptSvcPromptType = "Text-to-Text"
	PromptTypeQuestionAnswering PromptSvcPromptType = "Question Answering"
	PromptTypeTranslation PromptSvcPromptType = "Translation"
	PromptTypeSummarization PromptSvcPromptType = "Summarization"
	PromptTypeTextGeneration PromptSvcPromptType = "Text Generation"
	PromptTypeFillMask PromptSvcPromptType = "Fill-Mask"
	PromptTypeTextToSpeech PromptSvcPromptType = "Text-to-Speech"
	PromptTypeTextToAudio PromptSvcPromptType = "Text-to-Audio"
	PromptTypeAutomaticSpeechRecognition PromptSvcPromptType = "Automatic Speech Recognition"
	PromptTypeAudioToAudio PromptSvcPromptType = "Audio-to-Audio"
	PromptTypeAudioClassification PromptSvcPromptType = "Audio Classification"
	PromptTypeReinforcementLearning PromptSvcPromptType = "Reinforcement Learning"
	PromptTypeRobotics PromptSvcPromptType = "Robotics"
	PromptTypeGraphMachineLearning PromptSvcPromptType = "Graph Machine Learning"
)

// All allowed values of PromptSvcPromptType enum
var AllowedPromptSvcPromptTypeEnumValues = []PromptSvcPromptType{
	"Image-Text-to-Text",
	"Visual Question Answering",
	"Document Question Answering",
	"Text-to-Image",
	"Image-to-Image",
	"Image-to-Video",
	"Unconditional Image Generation",
	"Text-to-Video",
	"Zero-Shot Image Classification",
	"Zero-Shot Object Detection",
	"Text-to-3D",
	"Image-to-3D",
	"Image Feature Extraction",
	"Keypoint Detection",
	"Text-to-Text",
	"Question Answering",
	"Translation",
	"Summarization",
	"Text Generation",
	"Fill-Mask",
	"Text-to-Speech",
	"Text-to-Audio",
	"Automatic Speech Recognition",
	"Audio-to-Audio",
	"Audio Classification",
	"Reinforcement Learning",
	"Robotics",
	"Graph Machine Learning",
}

func (v *PromptSvcPromptType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PromptSvcPromptType(value)
	for _, existing := range AllowedPromptSvcPromptTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PromptSvcPromptType", value)
}

// NewPromptSvcPromptTypeFromValue returns a pointer to a valid PromptSvcPromptType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromptSvcPromptTypeFromValue(v string) (*PromptSvcPromptType, error) {
	ev := PromptSvcPromptType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromptSvcPromptType: valid values are %v", v, AllowedPromptSvcPromptTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromptSvcPromptType) IsValid() bool {
	for _, existing := range AllowedPromptSvcPromptTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to prompt_svc.PromptType value
func (v PromptSvcPromptType) Ptr() *PromptSvcPromptType {
	return &v
}

type NullablePromptSvcPromptType struct {
	value *PromptSvcPromptType
	isSet bool
}

func (v NullablePromptSvcPromptType) Get() *PromptSvcPromptType {
	return v.value
}

func (v *NullablePromptSvcPromptType) Set(val *PromptSvcPromptType) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcPromptType) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcPromptType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcPromptType(val *PromptSvcPromptType) *NullablePromptSvcPromptType {
	return &NullablePromptSvcPromptType{value: val, isSet: true}
}

func (v NullablePromptSvcPromptType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcPromptType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

