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
export class UserSvcGrant {
    static getAttributeTypeMap() {
        return UserSvcGrant.attributeTypeMap;
    }
}
UserSvcGrant.discriminator = undefined;
UserSvcGrant.attributeTypeMap = [
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "permissionId",
        "baseName": "permissionId",
        "type": "string"
    },
    {
        "name": "slugs",
        "baseName": "slugs",
        "type": "Array<string>"
    }
];
