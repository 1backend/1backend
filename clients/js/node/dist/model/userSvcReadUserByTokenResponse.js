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
export class UserSvcReadUserByTokenResponse {
    static getAttributeTypeMap() {
        return UserSvcReadUserByTokenResponse.attributeTypeMap;
    }
}
UserSvcReadUserByTokenResponse.discriminator = undefined;
UserSvcReadUserByTokenResponse.attributeTypeMap = [
    {
        "name": "activeOrganizationId",
        "baseName": "activeOrganizationId",
        "type": "string"
    },
    {
        "name": "organizations",
        "baseName": "organizations",
        "type": "Array<UserSvcOrganization>"
    },
    {
        "name": "user",
        "baseName": "user",
        "type": "UserSvcUser"
    }
];
