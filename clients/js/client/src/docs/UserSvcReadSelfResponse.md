
# UserSvcReadSelfResponse


## Properties

Name | Type
------------ | -------------
`activeOrganizationId` | string
`contacts` | [Array&lt;UserSvcContact&gt;](UserSvcContact.md)
`organizations` | [Array&lt;UserSvcOrganization&gt;](UserSvcOrganization.md)
`roles` | Array&lt;string&gt;
`tokenCount` | number
`user` | [UserSvcUser](UserSvcUser.md)

## Example

```typescript
import type { UserSvcReadSelfResponse } from ''

// TODO: Update the object below with actual values
const example = {
  "activeOrganizationId": null,
  "contacts": null,
  "organizations": null,
  "roles": null,
  "tokenCount": null,
  "user": null,
} satisfies UserSvcReadSelfResponse

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcReadSelfResponse
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


