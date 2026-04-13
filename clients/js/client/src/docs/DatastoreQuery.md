
# DatastoreQuery


## Properties

Name | Type
------------ | -------------
`afterJson` | string
`count` | boolean
`filters` | [Array&lt;DatastoreFilter&gt;](DatastoreFilter.md)
`limit` | number
`orderBys` | [Array&lt;DatastoreOrderBy&gt;](DatastoreOrderBy.md)

## Example

```typescript
import type { DatastoreQuery } from ''

// TODO: Update the object below with actual values
const example = {
  "afterJson": null,
  "count": null,
  "filters": null,
  "limit": null,
  "orderBys": null,
} satisfies DatastoreQuery

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as DatastoreQuery
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


