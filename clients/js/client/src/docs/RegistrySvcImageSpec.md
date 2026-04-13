
# RegistrySvcImageSpec


## Properties

Name | Type
------------ | -------------
`internalPorts` | Array&lt;number&gt;
`name` | string

## Example

```typescript
import type { RegistrySvcImageSpec } from ''

// TODO: Update the object below with actual values
const example = {
  "internalPorts": null,
  "name": nginx:latest,
} satisfies RegistrySvcImageSpec

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as RegistrySvcImageSpec
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


