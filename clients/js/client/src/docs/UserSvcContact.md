
# UserSvcContact


## Properties

Name | Type
------------ | -------------
`createdAt` | string
`deletedAt` | string
`handle` | string
`id` | string
`isPrimary` | boolean
`platform` | string
`updatedAt` | string
`userId` | string
`verified` | boolean

## Example

```typescript
import type { UserSvcContact } from ''

// TODO: Update the object below with actual values
const example = {
  "createdAt": null,
  "deletedAt": null,
  "handle": thejoe,
  "id": twitter.com/thejoe,
  "isPrimary": null,
  "platform": twitter,
  "updatedAt": null,
  "userId": null,
  "verified": null,
} satisfies UserSvcContact

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcContact
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


