/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class PromptSvcTextToImageParameters {
    static getAttributeTypeMap() {
        return PromptSvcTextToImageParameters.attributeTypeMap;
    }
}
PromptSvcTextToImageParameters.discriminator = undefined;
PromptSvcTextToImageParameters.attributeTypeMap = [
    {
        "name": "aspectRatio",
        "baseName": "aspectRatio",
        "type": "string"
    },
    {
        "name": "batchSize",
        "baseName": "batchSize",
        "type": "number"
    },
    {
        "name": "denoisingStrength",
        "baseName": "denoisingStrength",
        "type": "number"
    },
    {
        "name": "enableUpscaling",
        "baseName": "enableUpscaling",
        "type": "boolean"
    },
    {
        "name": "format",
        "baseName": "format",
        "type": "string"
    },
    {
        "name": "guidanceScale",
        "baseName": "guidanceScale",
        "type": "number"
    },
    {
        "name": "height",
        "baseName": "height",
        "type": "number"
    },
    {
        "name": "negativePrompt",
        "baseName": "negativePrompt",
        "type": "string"
    },
    {
        "name": "numIterations",
        "baseName": "numIterations",
        "type": "number"
    },
    {
        "name": "prompt",
        "baseName": "prompt",
        "type": "string"
    },
    {
        "name": "qualityPreset",
        "baseName": "qualityPreset",
        "type": "string"
    },
    {
        "name": "restoreFaces",
        "baseName": "restoreFaces",
        "type": "boolean"
    },
    {
        "name": "scheduler",
        "baseName": "scheduler",
        "type": "string"
    },
    {
        "name": "seed",
        "baseName": "seed",
        "type": "number"
    },
    {
        "name": "steps",
        "baseName": "steps",
        "type": "number"
    },
    {
        "name": "styles",
        "baseName": "styles",
        "type": "Array<string>"
    },
    {
        "name": "width",
        "baseName": "width",
        "type": "number"
    }
];
