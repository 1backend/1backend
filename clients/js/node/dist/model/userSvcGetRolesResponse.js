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
export class UserSvcGetRolesResponse {
    static getAttributeTypeMap() {
        return UserSvcGetRolesResponse.attributeTypeMap;
    }
}
UserSvcGetRolesResponse.discriminator = undefined;
UserSvcGetRolesResponse.attributeTypeMap = [
    {
        "name": "roles",
        "baseName": "roles",
        "type": "Array<UserSvcRole>"
    }
];
