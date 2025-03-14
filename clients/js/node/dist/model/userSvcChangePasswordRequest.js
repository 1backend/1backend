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
export class UserSvcChangePasswordRequest {
    static getAttributeTypeMap() {
        return UserSvcChangePasswordRequest.attributeTypeMap;
    }
}
UserSvcChangePasswordRequest.discriminator = undefined;
UserSvcChangePasswordRequest.attributeTypeMap = [
    {
        "name": "currentPassword",
        "baseName": "currentPassword",
        "type": "string"
    },
    {
        "name": "newPassword",
        "baseName": "newPassword",
        "type": "string"
    },
    {
        "name": "slug",
        "baseName": "slug",
        "type": "string"
    }
];
