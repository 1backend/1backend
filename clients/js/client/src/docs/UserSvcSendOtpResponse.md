
# UserSvcSendOtpResponse


## Properties

Name | Type
------------ | -------------
`body` | string
`code` | string
`contentType` | string
`fromEmail` | string
`fromName` | string
`otpId` | string
`subject` | string

## Example

```typescript
import type { UserSvcSendOtpResponse } from ''

// TODO: Update the object below with actual values
const example = {
  "body": null,
  "code": null,
  "contentType": null,
  "fromEmail": null,
  "fromName": null,
  "otpId": null,
  "subject": null,
} satisfies UserSvcSendOtpResponse

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcSendOtpResponse
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


