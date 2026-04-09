
# ConfigSvcListVersionsRequest


## Properties

Name | Type
------------ | -------------
`afterJson` | string
`appHost` | string
`branch` | string
`ids` | Array&lt;string&gt;
`limit` | number
`selector` | { [key: string]: Array&lt;string&gt;; }
`versionIds` | Array&lt;string&gt;

## Example

```typescript
import type { ConfigSvcListVersionsRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "afterJson": null,
  "appHost": shoes.com,
  "branch": null,
  "ids": null,
  "limit": null,
  "selector": null,
  "versionIds": null,
} satisfies ConfigSvcListVersionsRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ConfigSvcListVersionsRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


