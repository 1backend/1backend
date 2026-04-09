
# ModelSvcGetModelResponse


## Properties

Name | Type
------------ | -------------
`_exists` | boolean
`model` | [ModelSvcModel](ModelSvcModel.md)
`platform` | [ModelSvcPlatform](ModelSvcPlatform.md)

## Example

```typescript
import type { ModelSvcGetModelResponse } from ''

// TODO: Update the object below with actual values
const example = {
  "_exists": null,
  "model": null,
  "platform": null,
} satisfies ModelSvcGetModelResponse

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ModelSvcGetModelResponse
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


