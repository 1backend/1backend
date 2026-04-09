
# DataSvcQueryRequest


## Properties

Name | Type
------------ | -------------
`query` | [DatastoreQuery](DatastoreQuery.md)
`readers` | Array&lt;string&gt;
`table` | string

## Example

```typescript
import type { DataSvcQueryRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "query": null,
  "readers": null,
  "table": null,
} satisfies DataSvcQueryRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as DataSvcQueryRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


