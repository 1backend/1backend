
# RegistrySvcNode


## Properties

Name | Type
------------ | -------------
`availabilityZone` | string
`gpus` | [Array&lt;RegistrySvcGPU&gt;](RegistrySvcGPU.md)
`id` | string
`lastHeartbeat` | string
`region` | string
`url` | string
`usage` | [RegistrySvcResourceUsage](RegistrySvcResourceUsage.md)

## Example

```typescript
import type { RegistrySvcNode } from ''

// TODO: Update the object below with actual values
const example = {
  "availabilityZone": null,
  "gpus": null,
  "id": node_di9riJEvH2,
  "lastHeartbeat": null,
  "region": null,
  "url": null,
  "usage": null,
} satisfies RegistrySvcNode

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RegistrySvcNode
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


