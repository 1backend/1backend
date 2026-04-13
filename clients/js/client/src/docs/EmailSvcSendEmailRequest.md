
# EmailSvcSendEmailRequest


## Properties

Name | Type
------------ | -------------
`attachments` | [Array&lt;EmailSvcAttachment&gt;](EmailSvcAttachment.md)
`bcc` | Array&lt;string&gt;
`body` | string
`cc` | Array&lt;string&gt;
`contentType` | string
`fromEmail` | string
`fromName` | string
`id` | string
`subject` | string
`to` | Array&lt;string&gt;

## Example

```typescript
import type { EmailSvcSendEmailRequest } from ''

// TODO: Update the object below with actual values
const example = {
  "attachments": null,
  "bcc": null,
  "body": null,
  "cc": null,
  "contentType": null,
  "fromEmail": null,
  "fromName": null,
  "id": null,
  "subject": null,
  "to": null,
} satisfies EmailSvcSendEmailRequest

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as EmailSvcSendEmailRequest
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


