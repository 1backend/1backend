
# RegistrySvcAPISpec


## Properties

Name | Type
------------ | -------------
`metadata` | { [key: string]: string; }
`protocolType` | string
`url` | string
`version` | string

## Example

```typescript
import type { RegistrySvcAPISpec } from ''

// TODO: Update the object below with actual values
const example = {
  "metadata": null,
  "protocolType": null,
  "url": null,
  "version": null,
} satisfies RegistrySvcAPISpec

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RegistrySvcAPISpec
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


