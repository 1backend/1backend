
# ProxySvcCert


## Properties

Name | Type
------------ | -------------
`cert` | string
`commonName` | string
`createdAt` | string
`dnsNames` | Array&lt;string&gt;
`id` | string
`isCA` | boolean
`issuer` | string
`notAfter` | string
`notBefore` | string
`publicKeyAlgorithm` | string
`publicKeyBitLength` | number
`serialNumber` | string
`signatureAlgorithm` | string
`updatedAt` | string

## Example

```typescript
import type { ProxySvcCert } from ''

// TODO: Update the object below with actual values
const example = {
  "cert": null,
  "commonName": null,
  "createdAt": null,
  "dnsNames": null,
  "id": example.com,
  "isCA": null,
  "issuer": null,
  "notAfter": null,
  "notBefore": null,
  "publicKeyAlgorithm": null,
  "publicKeyBitLength": null,
  "serialNumber": null,
  "signatureAlgorithm": null,
  "updatedAt": null,
} satisfies ProxySvcCert

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as ProxySvcCert
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


