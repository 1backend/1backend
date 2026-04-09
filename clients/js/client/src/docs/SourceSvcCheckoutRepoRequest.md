
# SourceSvcCheckoutRepoRequest


## Properties

Name | Type
------------ | -------------
`password` | string
`ssh_key` | string
`ssh_key_pwd` | string
`token` | string
`url` | string
`username` | string
`version` | string

## Example

```typescript
import type { SourceSvcCheckoutRepoRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "password": null,
  "ssh_key": null,
  "ssh_key_pwd": null,
  "token": null,
  "url": null,
  "username": null,
  "version": null,
} satisfies SourceSvcCheckoutRepoRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SourceSvcCheckoutRepoRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


