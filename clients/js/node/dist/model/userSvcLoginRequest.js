/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcLoginRequest {
    static getAttributeTypeMap() {
        return UserSvcLoginRequest.attributeTypeMap;
    }
}
UserSvcLoginRequest.discriminator = undefined;
UserSvcLoginRequest.attributeTypeMap = [
    {
        "name": "app",
        "baseName": "app",
        "type": "string"
    },
    {
        "name": "contact",
        "baseName": "contact",
        "type": "string"
    },
    {
        "name": "device",
        "baseName": "device",
        "type": "string"
    },
    {
        "name": "password",
        "baseName": "password",
        "type": "string"
    },
    {
        "name": "slug",
        "baseName": "slug",
        "type": "string"
    }
];
