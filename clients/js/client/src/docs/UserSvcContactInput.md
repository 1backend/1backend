
# UserSvcContactInput


## Properties

Name | Type
------------ | -------------
`id` | string
`otpCode` | string
`otpId` | string
`platform` | string

## Example

```typescript
import type { UserSvcContactInput } from ''

// TODO: Update the object below with actual values
const example = {
  "id": twitter.com/thejoe,
  "otpCode": null,
  "otpId": null,
  "platform": twitter,
} satisfies UserSvcContactInput

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcContactInput
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


