
# ContainerSvcRunContainerRequest


## Properties

Name | Type
------------ | -------------
`assets` | [Array&lt;ContainerSvcAsset&gt;](ContainerSvcAsset.md)
`capabilities` | [ContainerSvcCapabilities](ContainerSvcCapabilities.md)
`envs` | [Array&lt;ContainerSvcEnvVar&gt;](ContainerSvcEnvVar.md)
`hash` | string
`image` | string
`keeps` | [Array&lt;ContainerSvcKeep&gt;](ContainerSvcKeep.md)
`labels` | [Array&lt;ContainerSvcLabel&gt;](ContainerSvcLabel.md)
`names` | Array&lt;string&gt;
`ports` | [Array&lt;ContainerSvcPortMapping&gt;](ContainerSvcPortMapping.md)

## Example

```typescript
import type { ContainerSvcRunContainerRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "assets": null,
  "capabilities": null,
  "envs": null,
  "hash": null,
  "image": nginx:latest,
  "keeps": null,
  "labels": null,
  "names": null,
  "ports": null,
} satisfies ContainerSvcRunContainerRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ContainerSvcRunContainerRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


