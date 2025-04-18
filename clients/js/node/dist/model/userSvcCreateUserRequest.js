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
export class UserSvcCreateUserRequest {
    static getAttributeTypeMap() {
        return UserSvcCreateUserRequest.attributeTypeMap;
    }
}
UserSvcCreateUserRequest.discriminator = undefined;
UserSvcCreateUserRequest.attributeTypeMap = [
    {
        "name": "contacts",
        "baseName": "contacts",
        "type": "Array<UserSvcContact>"
    },
    {
        "name": "password",
        "baseName": "password",
        "type": "string"
    },
    {
        "name": "roleIds",
        "baseName": "roleIds",
        "type": "Array<string>"
    },
    {
        "name": "user",
        "baseName": "user",
        "type": "UserSvcUser"
    }
];
