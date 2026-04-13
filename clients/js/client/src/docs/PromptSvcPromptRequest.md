
# PromptSvcPromptRequest


## Properties

Name | Type
------------ | -------------
`engineParameters` | [PromptSvcEngineParameters](PromptSvcEngineParameters.md)
`id` | string
`maxRetries` | number
`modelId` | string
`parameters` | [PromptSvcParameters](PromptSvcParameters.md)
`prompt` | string
`sync` | boolean
`threadId` | string

## Example

```typescript
import type { PromptSvcPromptRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "engineParameters": null,
  "id": null,
  "maxRetries": 10,
  "modelId": huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf,
  "parameters": null,
  "prompt": What's a banana?,
  "sync": null,
  "threadId": null,
} satisfies PromptSvcPromptRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PromptSvcPromptRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


