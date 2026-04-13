
# RegistrySvcInstance


## Properties

Name | Type
------------ | -------------
`deploymentId` | string
`details` | string
`host` | string
`id` | string
`ip` | string
`lastHeartbeat` | string
`nodeUrl` | string
`path` | string
`port` | number
`scheme` | string
`slug` | string
`status` | [RegistrySvcInstanceStatus](RegistrySvcInstanceStatus.md)
`tags` | Array&lt;string&gt;
`url` | string

## Example

```typescript
import type { RegistrySvcInstance } from ''

// TODO: Update the object below with actual values
const example = {
  "deploymentId": depl_deBUCtJirc,
  "details": Instance is healthy,
  "host": myserver.com,
  "id": inst_di9riJEvH2,
  "ip": 8.8.8.8,
  "lastHeartbeat": null,
  "nodeUrl": https://myserver.com:11337,
  "path": /your-svc,
  "port": 8080,
  "scheme": https,
  "slug": my-svc,
  "status": null,
  "tags": [tag1, tag2],
  "url": https://myserver.com:5981,
} satisfies RegistrySvcInstance

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RegistrySvcInstance
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


