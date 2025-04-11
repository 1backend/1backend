/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcSaveOrganizationRequest {
    static getAttributeTypeMap() {
        return UserSvcSaveOrganizationRequest.attributeTypeMap;
    }
}
UserSvcSaveOrganizationRequest.discriminator = undefined;
UserSvcSaveOrganizationRequest.attributeTypeMap = [
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
        "name": "thumbnailFileId",
        "baseName": "thumbnailFileId",
        "type": "string"
    }
];
