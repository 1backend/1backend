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
export class RegistrySvcListNodesRequest {
    static getAttributeTypeMap() {
        return RegistrySvcListNodesRequest.attributeTypeMap;
    }
}
RegistrySvcListNodesRequest.discriminator = undefined;
RegistrySvcListNodesRequest.attributeTypeMap = [
    {
        "name": "ids",
        "baseName": "ids",
        "type": "Array<string>"
    }
];
