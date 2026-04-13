
# RegistrySvcDefinition


## Properties

Name | Type
------------ | -------------
`apiSpecs` | [Array&lt;RegistrySvcAPISpec&gt;](RegistrySvcAPISpec.md)
`clients` | [Array&lt;RegistrySvcClient&gt;](RegistrySvcClient.md)
`envars` | [Array&lt;RegistrySvcEnvVar&gt;](RegistrySvcEnvVar.md)
`id` | string
`image` | [RegistrySvcImageSpec](RegistrySvcImageSpec.md)
`ports` | [Array&lt;RegistrySvcPortMapping&gt;](RegistrySvcPortMapping.md)
`repository` | [RegistrySvcRepositorySpec](RegistrySvcRepositorySpec.md)

## Example

```typescript
import type { RegistrySvcDefinition } from ''

// TODO: Update the object below with actual values
const example = {
  "apiSpecs": null,
  "clients": null,
  "envars": null,
  "id": null,
  "image": null,
  "ports": null,
  "repository": null,
} satisfies RegistrySvcDefinition

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RegistrySvcDefinition
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


