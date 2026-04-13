
# PromptSvcTextToImageParameters


## Properties

Name | Type
------------ | -------------
`aspectRatio` | string
`batchSize` | number
`denoisingStrength` | number
`enableUpscaling` | boolean
`format` | string
`guidanceScale` | number
`height` | number
`negativePrompt` | string
`numIterations` | number
`prompt` | string
`qualityPreset` | string
`restoreFaces` | boolean
`scheduler` | string
`seed` | number
`steps` | number
`styles` | Array&lt;string&gt;
`width` | number

## Example

```typescript
import type { PromptSvcTextToImageParameters } from ''

// TODO: Update the object below with actual values
const example = {
  "aspectRatio": null,
  "batchSize": null,
  "denoisingStrength": null,
  "enableUpscaling": null,
  "format": null,
  "guidanceScale": null,
  "height": null,
  "negativePrompt": null,
  "numIterations": null,
  "prompt": null,
  "qualityPreset": null,
  "restoreFaces": null,
  "scheduler": null,
  "seed": null,
  "steps": null,
  "styles": null,
  "width": null,
} satisfies PromptSvcTextToImageParameters

console.log(example)

// Convert the instance to a JSON string
const exampleJSON: string = JSON.stringify(example)
console.log(exampleJSON)

// Parse the JSON string back to an object
const exampleParsed = JSON.parse(exampleJSON) as PromptSvcTextToImageParameters
console.log(exampleParsed)
```

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


