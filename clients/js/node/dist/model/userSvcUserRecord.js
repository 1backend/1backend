/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcUserRecord {
    static getAttributeTypeMap() {
        return UserSvcUserRecord.attributeTypeMap;
    }
}
UserSvcUserRecord.discriminator = undefined;
UserSvcUserRecord.attributeTypeMap = [
    {
        "name": "contactIds",
        "baseName": "contactIds",
        "type": "Array<string>"
    },
    {
        "name": "createdAt",
        "baseName": "createdAt",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "name",
        "baseName": "name",
        "type": "string"
    },
    {
        "name": "roleIds",
        "baseName": "roleIds",
        "type": "Array<string>"
    },
    {
        "name": "slug",
        "baseName": "slug",
        "type": "string"
    },
    {
        "name": "updatedAt",
        "baseName": "updatedAt",
        "type": "string"
    }
];
