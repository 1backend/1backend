/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package prompt_svc

import (
	"fmt"
	"time"

	"github.com/openorch/openorch/sdk/go/clients/stable_diffusion"
	"github.com/openorch/openorch/sdk/go/datastore"
	streammanager "github.com/openorch/openorch/server/internal/services/prompt/stream"

	chat "github.com/openorch/openorch/server/internal/services/chat/types"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type PromptStatus string

const (
	PromptStatusScheduled PromptStatus = "scheduled"
	PromptStatusRunning   PromptStatus = "running"
	PromptStatusCompleted PromptStatus = "completed"
	// Errored means it will be still retried
	PromptStatusErrored   PromptStatus = "errored"
	PromptStatusAbandoned PromptStatus = "abandoned"
	PromptStatusCanceled  PromptStatus = "canceled"
)

type PromptType string

// Scan implementations like this will cause issues
// once we start using generated openapi types cross services
// (as we should in a microservices setting).
func (pt *PromptType) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*pt = PromptType(string(v))
	case string:
		*pt = PromptType(v)
	default:
		return fmt.Errorf("unexpected PromptType type: '%T'", value)
	}

	return nil
}

const (
	// Multimodal prompts
	PromptTypeImageTextToText              PromptType = "Image-Text-to-Text"
	PromptTypeVisualQuestionAnswering      PromptType = "Visual Question Answering"
	PromptTypeDocumentQuestionAnswering    PromptType = "Document Question Answering"
	PromptTypeTextToImage                  PromptType = "Text-to-Image"
	PromptTypeImageToImage                 PromptType = "Image-to-Image"
	PromptTypeImageToVideo                 PromptType = "Image-to-Video"
	PromptTypeUnconditionalImageGeneration PromptType = "Unconditional Image Generation"
	PromptTypeTextToVideo                  PromptType = "Text-to-Video"
	PromptTypeZeroShotImageClassification  PromptType = "Zero-Shot Image Classification"
	PromptTypeZeroShotObjectDetection      PromptType = "Zero-Shot Object Detection"
	PromptTypeTextTo3D                     PromptType = "Text-to-3D"
	PromptTypeImageTo3D                    PromptType = "Image-to-3D"
	PromptTypeImageFeatureExtraction       PromptType = "Image Feature Extraction"
	PromptTypeKeypointDetection            PromptType = "Keypoint Detection"

	// NLP prompts
	PromptTypeTextToText        PromptType = "Text-to-Text"
	PromptTypeQuestionAnswering PromptType = "Question Answering"
	PromptTypeTranslation       PromptType = "Translation"
	PromptTypeSummarization     PromptType = "Summarization"
	PromptTypeTextGeneration    PromptType = "Text Generation"
	PromptTypeFillMask          PromptType = "Fill-Mask"

	// Audio prompts
	PromptTypeTextToSpeech               PromptType = "Text-to-Speech"
	PromptTypeTextToAudio                PromptType = "Text-to-Audio"
	PromptTypeAutomaticSpeechRecognition PromptType = "Automatic Speech Recognition"
	PromptTypeAudioToAudio               PromptType = "Audio-to-Audio"
	PromptTypeAudioClassification        PromptType = "Audio Classification"

	// Other
	PromptTypeReinforcementLearning PromptType = "Reinforcement Learning"
	PromptTypeRobotics              PromptType = "Robotics"
	PromptTypeGraphMachineLearning  PromptType = "Graph Machine Learning"
)

type Prompt struct {
	// Id is the unique ID of the prompt.
	Id string `json:"id"`

	// Prompt is the message itself eg. "What's a banana?
	Prompt string `json:"prompt" example:"What's a banana?" binding:"required"`

	// Sync drives whether prompt add request should wait and hang until
	// the prompt is done executing. By default the prompt just gets put on a queue
	// and the client will just subscribe to a Thread Stream.
	// For quick and dirty scripting however it's often times easier to do things syncronously.
	// In those cases set Sync to true.
	Sync bool `json:"sync"`

	// Type is inferred from the `Parameters` or `EngineParameters` field.
	// Eg. A LLamaCpp prompt will be "Text-to-Text",
	// a Stabel Diffusion one will be "Text-to-Image" etc.
	Type PromptType `json:"type,omitempty"`

	// ThreadId is the ID of the thread a prompt belongs to.
	// Clients subscribe to Thread Streams to see the answer to a prompt,
	// or set `prompt.sync` to true for a blocking answer.
	ThreadId string `json:"threadId"`

	RequestMessageId  string `json:"requestMessageId"`
	ResponseMessageId string `json:"responseMessageId"`

	// ModelId is just the OpenOrch internal ID of the model.
	ModelId string `json:"modelId,omitempty" example:"huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf"`

	// MaxRetries specified how many times the system should retry a prompt when it keeps erroring.
	MaxRetries int `json:"maxRetries,omitempty" example:"10"`

	// CreatedAt is the time of the prompt creation.
	CreatedAt time.Time `json:"createdAt"`

	// UpdatedAt is the last time the prompt was updated.
	UpdatedAt time.Time `json:"updatedAt"`

	// Status of the prompt.
	Status PromptStatus `json:"status,omitempty"`

	// LastRun is the time of the last prompt run.
	LastRun time.Time `json:"lastRun,omitempty"`

	// RunCount is the number of times the prompt was retried due to errors
	RunCount int `json:"runCount,omitempty"`

	// Error that arose during prompt execution, if any.
	Error string `json:"error,omitempty"`

	// UserId contains the ID of the user who submitted the prompt.
	UserId string `json:"userId"`

	// AI engine/platform (eg. LlamaCpp, Stable Diffusion) specific parameters
	EngineParameters *EngineParameters `json:"engineParameters,omitempty"`

	// AI engine/platform (eg. LlamaCpp, Stable Diffusion) agnostic parameters.
	// Use these high level parameters when you don't care about the actual engine, only
	// the functionality (eg. text to image, image to image) it provides.
	Parameters *Parameters `json:"parameters,omitempty"`
}

type Parameters struct {
	TextToText  *TextToTextParameters  `json:"textToText,omitempty"`
	TextToImage *TextToImageParameters `json:"textToImage,omitempty"`
}

// TextToTextParameters represents a high-level, abstract text to text AI engine.
type TextToTextParameters struct {
	// Template of the prompt. Optional. If not present it's derived from ModelId.
	Template string `json:"template" example:"[INST]{prompt}[/INST]"`
}

// TextToImageParameters represents a high-level, abstract text to image AI engine.
type TextToImageParameters struct {
	// The primary prompt for generating the image.
	// Defaults to the top-level prompt if not specified.
	// If both are provided (which should be avoided), this field takes precedence.
	Prompt string `json:"prompt,omitempty"`

	// A negative prompt to specify what should be avoided in the image.
	NegativePrompt string `json:"negativePrompt,omitempty"`

	// List of artistic styles or themes to apply.
	Styles []string `json:"styles,omitempty"`

	// Optional seed for reproducibility. If not set, a random seed is used.
	Seed *int `json:"seed,omitempty"`

	// Number of images to generate per batch.
	BatchSize int `json:"batchSize,omitempty"`

	// Number of batches to generate.
	NumIterations int `json:"numIterations,omitempty"`

	// Number of inference steps for image generation.
	Steps int `json:"steps,omitempty"`

	// How closely the output should follow the prompt.
	GuidanceScale float64 `json:"guidanceScale,omitempty"`

	// Image dimensions (width and height in pixels).
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`

	// Alternative way to specify dimensions (e.g., "16:9", "1:1").
	AspectRatio string `json:"aspectRatio,omitempty"`

	// Controls how much variation is introduced in image modifications.
	DenoisingStrength float64 `json:"denoisingStrength,omitempty"`

	// Whether to apply AI-based upscaling.
	EnableUpscaling bool `json:"enableUpscaling,omitempty"`

	// Whether to enhance facial details for portraits.
	RestoreFaces bool `json:"restoreFaces,omitempty"`

	// Specifies the sampling method used during generation.
	Scheduler string `json:"scheduler,omitempty"`

	// Preset quality settings (e.g., Low, Medium, High, Ultra).
	QualityPreset string `json:"qualityPreset,omitempty"`

	// Output format for the generated image (png, jpg, webp, etc.).
	Format string `json:"format,omitempty"`
}

type EngineParameters struct {
	LlamaCpp        *LlamaCppParameters        `json:"llamaCppParameters,omitempty"`
	StableDiffusion *StableDiffusionParameters `json:"stableDiffusion,omitempty"`
}

type LlamaCppParameters struct {
	// Template of the prompt. Optional. If not present it's derived from ModelId.
	Template string `json:"template" example:"[INST]{prompt}[/INST]"`
}

type StableDiffusionParameters struct {
	// Text to image parameters
	Txt2Img *stable_diffusion.Txt2ImgRequest `json:"txt2Img,omitempty"`
}

func (c *Prompt) GetId() string {
	return c.Id
}

func (c *Prompt) GetUpdatedAt() string {
	return c.Id
}

type PromptRequest struct {
	// Id is the unique ID of the prompt.
	Id string `json:"id"`

	// Prompt is the message itself eg. "What's a banana?
	Prompt string `json:"prompt" example:"What's a banana?" binding:"required"`

	// Sync drives whether prompt add request should wait and hang until
	// the prompt is done executing. By default the prompt just gets put on a queue
	// and the client will just subscribe to a Thread Stream.
	// For quick and dirty scripting however it's often times easier to do things synchronously.
	// In those cases set Sync to true.
	Sync bool `json:"sync"`

	// ThreadId is the ID of the thread a prompt belongs to.
	// Clients subscribe to Thread Streams to see the answer to a prompt,
	// or set `prompt.sync` to true for a blocking answer.
	ThreadId string `json:"threadId"`

	// ModelId is just the OpenOrch internal ID of the model.
	ModelId string `json:"modelId,omitempty" example:"huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf"`

	// MaxRetries specified how many times the system should retry a prompt when it keeps erroring.
	MaxRetries int `json:"maxRetries,omitempty" example:"10"`

	// AI engine/platform (eg. Llama, Stable Diffusion) specific parameters
	EngineParameters *EngineParameters `json:"engineParameters,omitempty"`

	// AI engine/platform (eg. Llama, Stable Diffusion) agnostic parameters.
	// Use these high level parameters when you don't care about the actual engine, only
	// the functionality (eg. text to image, image to image) it provides.
	Parameters *Parameters `json:"parameters,omitempty"`
}

type PromptResponse struct {
	// Response message contains the response text and files.
	// This field is populated only for synchronous prompts (`sync = true`).
	// For asynchronous prompts, the response will provided in the associated
	// message identified by the `responseMessageId` of the `promptSvc.prompt` object once the prompt completes.
	ResponseMessage *chat.Message `json:"responseMessage,omitempty"`

	// Prompt contains the details of the prompt that was just created by this request.
	// This includes the ID, prompt text, status, and other associated metadata.
	Prompt *Prompt `json:"prompt,omitempty"`
}

type ListPromptsRequest struct {
	Query *datastore.Query `json:"query"`
}

type ListPromptsResponse struct {
	Prompts []*Prompt   `json:"prompts"`
	After   interface{} `json:"after,omitempty"`
	Count   int64       `json:"count"`
}

type RemovePromptRequest struct {
	PromptId string `json:"promptId"`
}

type RemovePromptResponse struct{}

type ListPromptOptions struct {
	Query *datastore.Query `json:"query"`
}

type TypesRequest struct {
}

type TypesResponse struct {
	Chunk streammanager.Chunk
}
