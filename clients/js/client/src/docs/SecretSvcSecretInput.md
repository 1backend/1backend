
# SecretSvcSecretInput


## Properties

Name | Type
------------ | -------------
`appHost` | string
`canChangeDeleters` | Array&lt;string&gt;
`canChangeReaders` | Array&lt;string&gt;
`canChangeWriters` | Array&lt;string&gt;
`checksum` | string
`checksumAlgorithm` | [SecretSvcChecksumAlgorithm](SecretSvcChecksumAlgorithm.md)
`deleters` | Array&lt;string&gt;
`encrypted` | boolean
`id` | string
`readers` | Array&lt;string&gt;
`value` | string
`writers` | Array&lt;string&gt;

## Example

```typescript
import type { SecretSvcSecretInput } from ''

// TODO: Update the object below with actual values
const example = {
  "appHost": null,
  "canChangeDeleters": null,
  "canChangeReaders": null,
  "canChangeWriters": null,
  "checksum": null,
  "checksumAlgorithm": null,
  "deleters": null,
  "encrypted": null,
  "id": null,
  "readers": null,
  "value": null,
  "writers": null,
} satisfies SecretSvcSecretInput

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as SecretSvcSecretInput
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


