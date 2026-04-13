
# ModelSvcContainer


## Properties

Name | Type
------------ | -------------
`envars` | [Array&lt;ModelSvcEnvVar&gt;](ModelSvcEnvVar.md)
`imageTemplate` | string
`keeps` | [Array&lt;ModelSvcKeep&gt;](ModelSvcKeep.md)
`port` | number

## Example

```typescript
import type { ModelSvcContainer } from ''

// TODO: Update the object below with actual values
const example = {
  "envars": null,
  "imageTemplate": null,
  "keeps": null,
  "port": null,
} satisfies ModelSvcContainer

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ModelSvcContainer
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


