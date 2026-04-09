
# ModelSvcModel


## Properties

Name | Type
------------ | -------------
`assets` | [Array&lt;ModelSvcAsset&gt;](ModelSvcAsset.md)
`bits` | number
`description` | string
`extension` | string
`flavour` | string
`fullName` | string
`id` | string
`maxBits` | number
`maxRam` | number
`mirrors` | Array&lt;string&gt;
`name` | string
`parameters` | string
`platformId` | string
`promptTemplate` | string
`quality` | string
`quantComment` | string
`size` | number
`tags` | Array&lt;string&gt;
`uncensored` | boolean
`version` | string

## Example

```typescript
import type { ModelSvcModel } from ''

// TODO: Update the object below with actual values
const example = {
  "assets": null,
  "bits": null,
  "description": null,
  "extension": null,
  "flavour": null,
  "fullName": null,
  "id": null,
  "maxBits": null,
  "maxRam": null,
  "mirrors": null,
  "name": null,
  "parameters": null,
  "platformId": null,
  "promptTemplate": null,
  "quality": null,
  "quantComment": null,
  "size": null,
  "tags": null,
  "uncensored": null,
  "version": null,
} satisfies ModelSvcModel

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ModelSvcModel
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


