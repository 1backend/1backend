
# UserSvcSaveSelfRequest


## Properties

Name | Type
------------ | -------------
`labels` | { [key: string]: string; }
`name` | string
`thumbnailFileId` | string

## Example

```typescript
import type { UserSvcSaveSelfRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "labels": null,
  "name": null,
  "thumbnailFileId": file_fQDxusW8og,
} satisfies UserSvcSaveSelfRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcSaveSelfRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


