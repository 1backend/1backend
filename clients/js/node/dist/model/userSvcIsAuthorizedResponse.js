/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcIsAuthorizedResponse {
    static getAttributeTypeMap() {
        return UserSvcIsAuthorizedResponse.attributeTypeMap;
    }
}
UserSvcIsAuthorizedResponse.discriminator = undefined;
UserSvcIsAuthorizedResponse.attributeTypeMap = [
    {
        "name": "authorized",
        "baseName": "authorized",
        "type": "boolean"
    },
    {
        "name": "user",
        "baseName": "user",
        "type": "UserSvcUser"
    }
];
