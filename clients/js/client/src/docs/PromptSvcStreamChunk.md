
# PromptSvcStreamChunk


## Properties

Name | Type
------------ | -------------
`messageId` | string
`text` | string
`type` | [PromptSvcStreamChunkType](PromptSvcStreamChunkType.md)

## Example

```typescript
import type { PromptSvcStreamChunk } from ''

// TODO: Update the object below with actual values
const example = {
  "messageId": null,
  "text": null,
  "type": null,
} satisfies PromptSvcStreamChunk

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PromptSvcStreamChunk
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


