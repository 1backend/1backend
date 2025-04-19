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
export class UserSvcEnroll {
    static getAttributeTypeMap() {
        return UserSvcEnroll.attributeTypeMap;
    }
}
UserSvcEnroll.discriminator = undefined;
UserSvcEnroll.attributeTypeMap = [
    {
        "name": "contactId",
        "baseName": "contactId",
        "type": "string"
    },
    {
        "name": "createdAt",
        "baseName": "createdAt",
        "type": "string"
    },
    {
        "name": "createdBy",
        "baseName": "createdBy",
        "type": "string"
    },
    {
        "name": "deletedAt",
        "baseName": "deletedAt",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "role",
        "baseName": "role",
        "type": "string"
    },
    {
        "name": "updatedAt",
        "baseName": "updatedAt",
        "type": "string"
    },
    {
        "name": "userId",
        "baseName": "userId",
        "type": "string"
    }
];
