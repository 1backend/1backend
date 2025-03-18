/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcOrganization {
    static getAttributeTypeMap() {
        return UserSvcOrganization.attributeTypeMap;
    }
}
UserSvcOrganization.discriminator = undefined;
UserSvcOrganization.attributeTypeMap = [
    {
        "name": "createdAt",
        "baseName": "createdAt",
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
        "name": "name",
        "baseName": "name",
        "type": "string"
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
