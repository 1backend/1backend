
# FileSvcDownload


## Properties

Name | Type
------------ | -------------
`createdAt` | string
`downloadedBytes` | number
`error` | string
`fileName` | string
`filePath` | string
`fileSize` | number
`id` | string
`progress` | number
`status` | string
`updatedAt` | string
`url` | string

## Example

```typescript
import type { FileSvcDownload } from ''

// TODO: Update the object below with actual values
const example = {
  "createdAt": null,
  "downloadedBytes": null,
  "error": null,
  "fileName": null,
  "filePath": null,
  "fileSize": null,
  "id": null,
  "progress": null,
  "status": null,
  "updatedAt": null,
  "url": null,
} satisfies FileSvcDownload

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as FileSvcDownload
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


