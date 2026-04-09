
# ModelSvcPlatform


## Properties

Name | Type
------------ | -------------
`architectures` | [ModelSvcArchitectures](ModelSvcArchitectures.md)
`id` | string
`name` | string
`types` | [Array&lt;PromptSvcPromptType&gt;](PromptSvcPromptType.md)
`version` | number

## Example

```typescript
import type { ModelSvcPlatform } from ''

// TODO: Update the object below with actual values
const example = {
  "architectures": null,
  "id": null,
  "name": null,
  "types": null,
  "version": null,
} satisfies ModelSvcPlatform

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ModelSvcPlatform
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


