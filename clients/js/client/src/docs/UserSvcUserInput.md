
# UserSvcUserInput


## Properties

Name | Type
------------ | -------------
`id` | string
`labels` | { [key: string]: string; }
`name` | string
`slug` | string
`thumbnailFileId` | string

## Example

```typescript
import type { UserSvcUserInput } from ''

// TODO: Update the object below with actual values
const example = {
  "id": null,
  "labels": null,
  "name": Jane Doe,
  "slug": jane-doe,
  "thumbnailFileId": file_fQDyi1xdHK,
} satisfies UserSvcUserInput

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcUserInput
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


