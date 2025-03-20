/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class RegistrySvcAPISpec {
    static getAttributeTypeMap() {
        return RegistrySvcAPISpec.attributeTypeMap;
    }
}
RegistrySvcAPISpec.discriminator = undefined;
RegistrySvcAPISpec.attributeTypeMap = [
    {
        "name": "metadata",
        "baseName": "metadata",
        "type": "{ [key: string]: string; }"
    },
    {
        "name": "protocolType",
        "baseName": "protocolType",
        "type": "string"
    },
    {
        "name": "url",
        "baseName": "url",
        "type": "string"
    },
    {
        "name": "version",
        "baseName": "version",
        "type": "string"
    }
];
