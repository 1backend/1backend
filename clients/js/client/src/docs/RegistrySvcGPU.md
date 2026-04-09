
# RegistrySvcGPU


## Properties

Name | Type
------------ | -------------
`busId` | string
`computeMode` | string
`cudaVersion` | string
`gpuUtilization` | number
`id` | string
`intraNodeId` | number
`memoryTotal` | number
`memoryUsage` | number
`name` | string
`performanceState` | string
`powerCap` | number
`powerUsage` | number
`processDetails` | [Array&lt;RegistrySvcProcess&gt;](RegistrySvcProcess.md)
`temperature` | number

## Example

```typescript
import type { RegistrySvcGPU } from ''

// TODO: Update the object below with actual values
const example = {
  "busId": null,
  "computeMode": null,
  "cudaVersion": null,
  "gpuUtilization": null,
  "id": null,
  "intraNodeId": null,
  "memoryTotal": null,
  "memoryUsage": null,
  "name": null,
  "performanceState": null,
  "powerCap": null,
  "powerUsage": null,
  "processDetails": null,
  "temperature": null,
} satisfies RegistrySvcGPU

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RegistrySvcGPU
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


