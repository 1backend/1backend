
# ModelSvcCudaParameters


## Properties

Name | Type
------------ | -------------
`container` | [ModelSvcContainer](ModelSvcContainer.md)
`cudaVersionPrecision` | number
`defaultCudaVersion` | string
`defaultCudnnVersion` | string

## Example

```typescript
import type { ModelSvcCudaParameters } from ''

// TODO: Update the object below with actual values
const example = {
  "container": null,
  "cudaVersionPrecision": null,
  "defaultCudaVersion": null,
  "defaultCudnnVersion": null,
} satisfies ModelSvcCudaParameters

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ModelSvcCudaParameters
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


