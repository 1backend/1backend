
# DataSvcObject


## Properties

Name | Type
------------ | -------------
`authors` | Array&lt;string&gt;
`createdAt` | string
`data` | { [key: string]: any; }
`deleters` | Array&lt;string&gt;
`id` | string
`readers` | Array&lt;string&gt;
`table` | string
`updatedAt` | string
`writers` | Array&lt;string&gt;

## Example

```typescript
import type { DataSvcObject } from ''

// TODO: Update the object below with actual values
const example = {
  "authors": [[],
  "createdAt": null,
  "data": null,
  "deleters": [[],
  "id": null,
  "readers": [[],
  "table": null,
  "updatedAt": null,
  "writers": [[],
} satisfies DataSvcObject

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as DataSvcObject
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


