
# ChatSvcMessage


## Properties

Name | Type
------------ | -------------
`createdAt` | string
`fileIds` | Array&lt;string&gt;
`id` | string
`meta` | { [key: string]: any; }
`text` | string
`threadId` | string
`updatedAt` | string
`userId` | string

## Example

```typescript
import type { ChatSvcMessage } from ''

// TODO: Update the object below with actual values
const example = {
  "createdAt": null,
  "fileIds": null,
  "id": msg_emSOPlW58o,
  "meta": null,
  "text": null,
  "threadId": thr_emSOeEUWAg,
  "updatedAt": null,
  "userId": null,
} satisfies ChatSvcMessage

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ChatSvcMessage
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


