
# PromptSvcPromptResponse


## Properties

Name | Type
------------ | -------------
`prompt` | [PromptSvcPrompt](PromptSvcPrompt.md)
`responseMessage` | [ChatSvcMessage](ChatSvcMessage.md)

## Example

```typescript
import type { PromptSvcPromptResponse } from ''

// TODO: Update the object below with actual values
const example = {
  "prompt": null,
  "responseMessage": null,
} satisfies PromptSvcPromptResponse

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PromptSvcPromptResponse
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


