/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class PromptSvcEngineParameters {
    static getAttributeTypeMap() {
        return PromptSvcEngineParameters.attributeTypeMap;
    }
}
PromptSvcEngineParameters.discriminator = undefined;
PromptSvcEngineParameters.attributeTypeMap = [
    {
        "name": "llamaCppParameters",
        "baseName": "llamaCppParameters",
        "type": "PromptSvcLlamaCppParameters"
    },
    {
        "name": "stableDiffusion",
        "baseName": "stableDiffusion",
        "type": "PromptSvcStableDiffusionParameters"
    }
];
