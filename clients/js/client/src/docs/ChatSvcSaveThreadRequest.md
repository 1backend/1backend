
# ChatSvcSaveThreadRequest


## Properties

Name | Type
------------ | -------------
`id` | string
`title` | string
`topicIds` | Array&lt;string&gt;
`userIds` | Array&lt;string&gt;

## Example

```typescript
import type { ChatSvcSaveThreadRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "id": thr_emSQnpJbhG,
  "title": null,
  "topicIds": null,
  "userIds": null,
} satisfies ChatSvcSaveThreadRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ChatSvcSaveThreadRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


