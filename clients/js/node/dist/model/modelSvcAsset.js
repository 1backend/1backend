/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ModelSvcAsset {
    static getAttributeTypeMap() {
        return ModelSvcAsset.attributeTypeMap;
    }
}
ModelSvcAsset.discriminator = undefined;
ModelSvcAsset.attributeTypeMap = [
    {
        "name": "envVarKey",
        "baseName": "envVarKey",
        "type": "string"
    },
    {
        "name": "url",
        "baseName": "url",
        "type": "string"
    }
];
