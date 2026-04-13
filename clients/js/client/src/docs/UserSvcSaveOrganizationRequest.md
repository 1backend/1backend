
# UserSvcSaveOrganizationRequest


## Properties

Name | Type
------------ | -------------
`activate` | boolean
`assignCaller` | boolean
`id` | string
`name` | string
`slug` | string
`thumbnailFileId` | string

## Example

```typescript
import type { UserSvcSaveOrganizationRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "activate": true,
  "assignCaller": true,
  "id": null,
  "name": null,
  "slug": null,
  "thumbnailFileId": file_fQDxusW8og,
} satisfies UserSvcSaveOrganizationRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcSaveOrganizationRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


