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
export class PromptSvcParameters {
    static getAttributeTypeMap() {
        return PromptSvcParameters.attributeTypeMap;
    }
}
PromptSvcParameters.discriminator = undefined;
PromptSvcParameters.attributeTypeMap = [
    {
        "name": "textToImage",
        "baseName": "textToImage",
        "type": "PromptSvcTextToImageParameters"
    },
    {
        "name": "textToText",
        "baseName": "textToText",
        "type": "PromptSvcTextToTextParameters"
    }
];
