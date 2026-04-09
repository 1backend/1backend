
# ContainerSvcBuildImageRequest


## Properties

Name | Type
------------ | -------------
`contextPath` | string
`dockerfilePath` | string
`name` | string

## Example

```typescript
import type { ContainerSvcBuildImageRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "contextPath": .,
  "dockerfilePath": Dockerfile,
  "name": nginx:latest,
} satisfies ContainerSvcBuildImageRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ContainerSvcBuildImageRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


