/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ModelSvcKeep {
    static getAttributeTypeMap() {
        return ModelSvcKeep.attributeTypeMap;
    }
}
ModelSvcKeep.discriminator = undefined;
ModelSvcKeep.attributeTypeMap = [
    {
        "name": "path",
        "baseName": "path",
        "type": "string"
    },
    {
        "name": "readOnly",
        "baseName": "readOnly",
        "type": "boolean"
    }
];
