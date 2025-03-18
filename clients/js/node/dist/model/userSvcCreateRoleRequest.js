/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcCreateRoleRequest {
    static getAttributeTypeMap() {
        return UserSvcCreateRoleRequest.attributeTypeMap;
    }
}
UserSvcCreateRoleRequest.discriminator = undefined;
UserSvcCreateRoleRequest.attributeTypeMap = [
    {
        "name": "description",
        "baseName": "description",
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
        "name": "permissionIds",
        "baseName": "permissionIds",
        "type": "Array<string>"
    }
];
