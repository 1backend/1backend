
# ContainerSvcRunContainerResponse


## Properties

Name | Type
------------ | -------------
`ports` | [Array&lt;ContainerSvcPortMapping&gt;](ContainerSvcPortMapping.md)
`started` | boolean

## Example

```typescript
import type { ContainerSvcRunContainerResponse } from ''

// TODO: Update the object below with actual values
const example = {
  "ports": null,
  "started": null,
} satisfies ContainerSvcRunContainerResponse

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ContainerSvcRunContainerResponse
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


