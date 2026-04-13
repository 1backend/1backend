
# RegistrySvcRegisterInstanceRequest


## Properties

Name | Type
------------ | -------------
`deploymentId` | string
`host` | string
`id` | string
`ip` | string
`path` | string
`port` | number
`scheme` | string
`url` | string

## Example

```typescript
import type { RegistrySvcRegisterInstanceRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "deploymentId": depl_deBUCtJirc,
  "host": myserver.com,
  "id": inst_di9riJEvH2,
  "ip": 8.8.8.8,
  "path": /your-svc,
  "port": 8080,
  "scheme": https,
  "url": https://myserver.com:5981,
} satisfies RegistrySvcRegisterInstanceRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RegistrySvcRegisterInstanceRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


