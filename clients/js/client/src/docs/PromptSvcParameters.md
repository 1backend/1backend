
# PromptSvcParameters


## Properties

Name | Type
------------ | -------------
`textToImage` | [PromptSvcTextToImageParameters](PromptSvcTextToImageParameters.md)
`textToText` | [PromptSvcTextToTextParameters](PromptSvcTextToTextParameters.md)

## Example

```typescript
import type { PromptSvcParameters } from ''

// TODO: Update the object below with actual values
const example = {
  "textToImage": null,
  "textToText": null,
} satisfies PromptSvcParameters

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PromptSvcParameters
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


