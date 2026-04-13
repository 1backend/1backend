
# ConfigSvcSaveConfigRequest


## Properties

Name | Type
------------ | -------------
`appHost` | string
`branch` | string
`data` | { [key: string]: any; }
`dataJson` | string
`id` | string

## Example

```typescript
import type { ConfigSvcSaveConfigRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "appHost": null,
  "branch": null,
  "data": null,
  "dataJson": null,
  "id": null,
} satisfies ConfigSvcSaveConfigRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ConfigSvcSaveConfigRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


