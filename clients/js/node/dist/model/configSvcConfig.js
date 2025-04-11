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
export class ConfigSvcConfig {
    static getAttributeTypeMap() {
        return ConfigSvcConfig.attributeTypeMap;
    }
}
ConfigSvcConfig.discriminator = undefined;
ConfigSvcConfig.attributeTypeMap = [
    {
        "name": "data",
        "baseName": "data",
        "type": "{ [key: string]: any; }"
    },
    {
        "name": "dataJson",
        "baseName": "dataJson",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "namespace",
        "baseName": "namespace",
        "type": "string"
    }
];
