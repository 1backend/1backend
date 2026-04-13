
# UserSvcHasPermissionResponse


## Properties

Name | Type
------------ | -------------
`app` | [UserSvcApp](UserSvcApp.md)
`appId` | string
`authorized` | boolean
`until` | string
`user` | [UserSvcUser](UserSvcUser.md)

## Example

```typescript
import type { UserSvcHasPermissionResponse } from ''

// TODO: Update the object below with actual values
const example = {
  "app": null,
  "appId": null,
  "authorized": null,
  "until": null,
  "user": null,
} satisfies UserSvcHasPermissionResponse

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as UserSvcHasPermissionResponse
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


