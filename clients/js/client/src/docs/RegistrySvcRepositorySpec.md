
# RegistrySvcRepositorySpec


## Properties

Name | Type
------------ | -------------
`buildContext` | string
`containerFile` | string
`internalPorts` | Array&lt;number&gt;
`url` | string
`version` | string

## Example

```typescript
import type { RegistrySvcRepositorySpec } from ''

// TODO: Update the object below with actual values
const example = {
  "buildContext": path/to/subfolder,
  "containerFile": docker/Dockerfile,
  "internalPorts": null,
  "url": https://github.com/1backend/1backend.git,
  "version": v1.0.0,
} satisfies RegistrySvcRepositorySpec

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RegistrySvcRepositorySpec
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


