
# UserSvcListUsersRequest


## Properties

Name | Type
------------ | -------------
`afterTime` | string
`contactId` | string
`count` | boolean
`ids` | Array&lt;string&gt;
`limit` | number
`order` | [UserSvcOrderDirection](UserSvcOrderDirection.md)
`orderBy` | [UserSvcListUsersOrderBy](UserSvcListUsersOrderBy.md)
`search` | string

## Example

```typescript
import type { UserSvcListUsersRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "afterTime": null,
  "contactId": null,
  "count": false,
  "ids": null,
  "limit": 10,
  "order": null,
  "orderBy": null,
  "search": null,
} satisfies UserSvcListUsersRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcListUsersRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


