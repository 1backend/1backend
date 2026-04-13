
# UserSvcUserRecord


## Properties

Name | Type
------------ | -------------
`contactIds` | Array&lt;string&gt;
`createdAt` | string
`id` | string
`name` | string
`roles` | Array&lt;string&gt;
`slug` | string
`updatedAt` | string

## Example

```typescript
import type { UserSvcUserRecord } from ''

// TODO: Update the object below with actual values
const example = {
  "contactIds": null,
  "createdAt": null,
  "id": null,
  "name": Jane Doe,
  "roles": null,
  "slug": jane-doe,
  "updatedAt": null,
} satisfies UserSvcUserRecord

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcUserRecord
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


