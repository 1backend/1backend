
# UserSvcOrganization


## Properties

Name | Type
------------ | -------------
`appId` | string
`createdAt` | string
`deletedAt` | string
`id` | string
`internalId` | string
`name` | string
`slug` | string
`thumbnailFileId` | string
`updatedAt` | string

## Example

```typescript
import type { UserSvcOrganization } from ''

// TODO: Update the object below with actual values
const example = {
  "appId": null,
  "createdAt": null,
  "deletedAt": null,
  "id": null,
  "internalId": null,
  "name": Acme Corporation,
  "slug": acme-corporation,
  "thumbnailFileId": file_fQDxusW8og,
  "updatedAt": null,
} satisfies UserSvcOrganization

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcOrganization
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


