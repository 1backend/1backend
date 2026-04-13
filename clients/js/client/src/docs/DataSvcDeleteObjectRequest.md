
# DataSvcDeleteObjectRequest


## Properties

Name | Type
------------ | -------------
`filters` | [Array&lt;DatastoreFilter&gt;](DatastoreFilter.md)
`table` | string

## Example

```typescript
import type { DataSvcDeleteObjectRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "filters": null,
  "table": null,
} satisfies DataSvcDeleteObjectRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as DataSvcDeleteObjectRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


