/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcRegisterRequest {
    static getAttributeTypeMap() {
        return UserSvcRegisterRequest.attributeTypeMap;
    }
}
UserSvcRegisterRequest.discriminator = undefined;
UserSvcRegisterRequest.attributeTypeMap = [
    {
        "name": "contact",
        "baseName": "contact",
        "type": "UserSvcContact"
    },
    {
        "name": "name",
        "baseName": "name",
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
