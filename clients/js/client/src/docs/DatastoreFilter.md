
# DatastoreFilter


## Properties

Name | Type
------------ | -------------
`fields` | Array&lt;string&gt;
`op` | [DatastoreOp](DatastoreOp.md)
`subFilters` | [Array&lt;DatastoreFilter&gt;](DatastoreFilter.md)
`valuesJson` | string

## Example

```typescript
import type { DatastoreFilter } from ''

// TODO: Update the object below with actual values
const example = {
  "fields": null,
  "op": null,
  "subFilters": null,
  "valuesJson": null,
} satisfies DatastoreFilter

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as DatastoreFilter
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


