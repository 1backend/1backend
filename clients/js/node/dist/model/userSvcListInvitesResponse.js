/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcListInvitesResponse {
    static getAttributeTypeMap() {
        return UserSvcListInvitesResponse.attributeTypeMap;
    }
}
UserSvcListInvitesResponse.discriminator = undefined;
UserSvcListInvitesResponse.attributeTypeMap = [
    {
        "name": "invites",
        "baseName": "invites",
        "type": "Array<UserSvcInvite>"
    }
];
