
# PolicySvcInstance


## Properties

Name | Type
------------ | -------------
`endpoint` | string
`id` | string
`parameters` | [PolicySvcParameters](PolicySvcParameters.md)
`templateId` | [PolicySvcTemplateId](PolicySvcTemplateId.md)

## Example

```typescript
import type { PolicySvcInstance } from ''

// TODO: Update the object below with actual values
const example = {
  "endpoint": /user-svc/register,
  "id": null,
  "parameters": null,
  "templateId": null,
} satisfies PolicySvcInstance

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PolicySvcInstance
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


