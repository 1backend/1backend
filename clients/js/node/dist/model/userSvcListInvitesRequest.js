/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcListInvitesRequest {
    static getAttributeTypeMap() {
        return UserSvcListInvitesRequest.attributeTypeMap;
    }
}
UserSvcListInvitesRequest.discriminator = undefined;
UserSvcListInvitesRequest.attributeTypeMap = [
    {
        "name": "contactId",
        "baseName": "contactId",
        "type": "string"
    },
    {
        "name": "role",
        "baseName": "role",
        "type": "string"
    },
    {
        "name": "userId",
        "baseName": "userId",
        "type": "string"
    }
];
