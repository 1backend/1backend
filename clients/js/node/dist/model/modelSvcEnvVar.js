/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ModelSvcEnvVar {
    static getAttributeTypeMap() {
        return ModelSvcEnvVar.attributeTypeMap;
    }
}
ModelSvcEnvVar.discriminator = undefined;
ModelSvcEnvVar.attributeTypeMap = [
    {
        "name": "key",
        "baseName": "key",
        "type": "string"
    },
    {
        "name": "value",
        "baseName": "value",
        "type": "string"
    }
];
