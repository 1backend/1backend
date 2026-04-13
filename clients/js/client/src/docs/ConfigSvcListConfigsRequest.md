
# ConfigSvcListConfigsRequest


## Properties

Name | Type
------------ | -------------
`appHost` | string
`branch` | string
`ids` | Array&lt;string&gt;
`scope` | [ConfigSvcListConfigsScope](ConfigSvcListConfigsScope.md)
`selector` | { [key: string]: Array&lt;string&gt;; }

## Example

```typescript
import type { ConfigSvcListConfigsRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "appHost": shoes.com,
  "branch": null,
  "ids": null,
  "scope": null,
  "selector": null,
} satisfies ConfigSvcListConfigsRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ConfigSvcListConfigsRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


