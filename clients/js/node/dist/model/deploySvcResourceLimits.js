/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class DeploySvcResourceLimits {
    static getAttributeTypeMap() {
        return DeploySvcResourceLimits.attributeTypeMap;
    }
}
DeploySvcResourceLimits.discriminator = undefined;
DeploySvcResourceLimits.attributeTypeMap = [
    {
        "name": "cpu",
        "baseName": "cpu",
        "type": "string"
    },
    {
        "name": "memory",
        "baseName": "memory",
        "type": "string"
    },
    {
        "name": "vram",
        "baseName": "vram",
        "type": "string"
    }
];
