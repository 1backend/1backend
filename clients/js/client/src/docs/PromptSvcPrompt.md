
# PromptSvcPrompt


## Properties

Name | Type
------------ | -------------
`createdAt` | string
`engineParameters` | [PromptSvcEngineParameters](PromptSvcEngineParameters.md)
`error` | string
`id` | string
`lastRun` | string
`maxRetries` | number
`modelId` | string
`parameters` | [PromptSvcParameters](PromptSvcParameters.md)
`prompt` | string
`requestMessageId` | string
`responseMessageId` | string
`runCount` | number
`status` | [PromptSvcPromptStatus](PromptSvcPromptStatus.md)
`sync` | boolean
`threadId` | string
`type` | [PromptSvcPromptType](PromptSvcPromptType.md)
`updatedAt` | string
`userId` | string

## Example

```typescript
import type { PromptSvcPrompt } from ''

// TODO: Update the object below with actual values
const example = {
  "createdAt": null,
  "engineParameters": null,
  "error": null,
  "id": null,
  "lastRun": null,
  "maxRetries": 10,
  "modelId": huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf,
  "parameters": null,
  "prompt": What's a banana?,
  "requestMessageId": null,
  "responseMessageId": null,
  "runCount": null,
  "status": null,
  "sync": null,
  "threadId": null,
  "type": null,
  "updatedAt": null,
  "userId": null,
} satisfies PromptSvcPrompt

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PromptSvcPrompt
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


