
# PolicySvcParameters


## Properties

Name | Type
------------ | -------------
`blocklist` | [PolicySvcBlocklistParameters](PolicySvcBlocklistParameters.md)
`rateLimit` | [PolicySvcRateLimitParameters](PolicySvcRateLimitParameters.md)

## Example

```typescript
import type { PolicySvcParameters } from ''

// TODO: Update the object below with actual values
const example = {
  "blocklist": null,
  "rateLimit": null,
} satisfies PolicySvcParameters

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PolicySvcParameters
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


