
# UserSvcLoginRequest


## Properties

Name | Type
------------ | -------------
`appHost` | string
`contact` | [UserSvcContactInput](UserSvcContactInput.md)
`device` | string
`password` | string
`slug` | string

## Example

```typescript
import type { UserSvcLoginRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "appHost": shoes.com,
  "contact": null,
  "device": null,
  "password": null,
  "slug": null,
} satisfies UserSvcLoginRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcLoginRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


