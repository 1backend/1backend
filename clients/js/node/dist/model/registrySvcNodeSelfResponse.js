/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class RegistrySvcNodeSelfResponse {
    static getAttributeTypeMap() {
        return RegistrySvcNodeSelfResponse.attributeTypeMap;
    }
}
RegistrySvcNodeSelfResponse.discriminator = undefined;
RegistrySvcNodeSelfResponse.attributeTypeMap = [
    {
        "name": "node",
        "baseName": "node",
        "type": "RegistrySvcNode"
    }
];
