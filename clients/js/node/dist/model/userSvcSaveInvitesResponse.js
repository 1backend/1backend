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
export class UserSvcSaveInvitesResponse {
    static getAttributeTypeMap() {
        return UserSvcSaveInvitesResponse.attributeTypeMap;
    }
}
UserSvcSaveInvitesResponse.discriminator = undefined;
UserSvcSaveInvitesResponse.attributeTypeMap = [
    {
        "name": "invites",
        "baseName": "invites",
        "type": "Array<UserSvcInvite>"
    }
];
