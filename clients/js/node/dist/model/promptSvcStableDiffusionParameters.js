/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class PromptSvcStableDiffusionParameters {
    static getAttributeTypeMap() {
        return PromptSvcStableDiffusionParameters.attributeTypeMap;
    }
}
PromptSvcStableDiffusionParameters.discriminator = undefined;
PromptSvcStableDiffusionParameters.attributeTypeMap = [
    {
        "name": "txt2Img",
        "baseName": "txt2Img",
        "type": "StableDiffusionTxt2ImgRequest"
    }
];
