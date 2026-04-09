
# PromptSvcListPromptsResponse


## Properties

Name | Type
------------ | -------------
`after` | object
`count` | number
`prompts` | [Array&lt;PromptSvcPrompt&gt;](PromptSvcPrompt.md)

## Example

```typescript
import type { PromptSvcListPromptsResponse } from ''

// TODO: Update the object below with actual values
const example = {
  "after": null,
  "count": null,
  "prompts": null,
} satisfies PromptSvcListPromptsResponse

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PromptSvcListPromptsResponse
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


