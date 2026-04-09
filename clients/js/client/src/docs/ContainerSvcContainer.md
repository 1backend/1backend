
# ContainerSvcContainer


## Properties

Name | Type
------------ | -------------
`assets` | [Array&lt;ContainerSvcAsset&gt;](ContainerSvcAsset.md)
`capabilities` | [ContainerSvcCapabilities](ContainerSvcCapabilities.md)
`envs` | [Array&lt;ContainerSvcEnvVar&gt;](ContainerSvcEnvVar.md)
`hash` | string
`id` | string
`image` | string
`keeps` | [Array&lt;ContainerSvcKeep&gt;](ContainerSvcKeep.md)
`labels` | [Array&lt;ContainerSvcLabel&gt;](ContainerSvcLabel.md)
`names` | Array&lt;string&gt;
`network` | [ContainerSvcNetwork](ContainerSvcNetwork.md)
`nodeId` | string
`ports` | [Array&lt;ContainerSvcPortMapping&gt;](ContainerSvcPortMapping.md)
`resources` | [ContainerSvcResources](ContainerSvcResources.md)
`runtime` | string
`status` | string
`volumes` | [Array&lt;ContainerSvcVolume&gt;](ContainerSvcVolume.md)

## Example

```typescript
import type { ContainerSvcContainer } from ''

// TODO: Update the object below with actual values
const example = {
  "assets": null,
  "capabilities": null,
  "envs": null,
  "hash": null,
  "id": null,
  "image": null,
  "keeps": null,
  "labels": null,
  "names": null,
  "network": null,
  "nodeId": null,
  "ports": null,
  "resources": null,
  "runtime": null,
  "status": null,
  "volumes": null,
} satisfies ContainerSvcContainer

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ContainerSvcContainer
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


