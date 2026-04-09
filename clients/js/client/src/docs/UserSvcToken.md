
# UserSvcToken


## Properties

Name | Type
------------ | -------------
`active` | boolean
`app` | [UserSvcApp](UserSvcApp.md)
`appId` | string
`createdAt` | string
`deletedAt` | string
`device` | string
`expiresAt` | string
`id` | string
`internalId` | string
`lastRefreshedAt` | string
`token` | string
`updatedAt` | string
`userId` | string

## Example

```typescript
import type { UserSvcToken } from ''

// TODO: Update the object below with actual values
const example = {
  "active": null,
  "app": null,
  "appId": null,
  "createdAt": null,
  "deletedAt": null,
  "device": null,
  "expiresAt": null,
  "id": null,
  "internalId": null,
  "lastRefreshedAt": null,
  "token": null,
  "updatedAt": null,
  "userId": null,
} satisfies UserSvcToken

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcToken
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


