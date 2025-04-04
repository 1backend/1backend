/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ModelSvcGetModelResponse {
    static getAttributeTypeMap() {
        return ModelSvcGetModelResponse.attributeTypeMap;
    }
}
ModelSvcGetModelResponse.discriminator = undefined;
ModelSvcGetModelResponse.attributeTypeMap = [
    {
        "name": "exists",
        "baseName": "exists",
        "type": "boolean"
    },
    {
        "name": "model",
        "baseName": "model",
        "type": "ModelSvcModel"
    },
    {
        "name": "platform",
        "baseName": "platform",
        "type": "ModelSvcPlatform"
    }
];
