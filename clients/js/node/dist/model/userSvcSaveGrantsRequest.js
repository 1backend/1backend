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
export class UserSvcSaveGrantsRequest {
    static getAttributeTypeMap() {
        return UserSvcSaveGrantsRequest.attributeTypeMap;
    }
}
UserSvcSaveGrantsRequest.discriminator = undefined;
UserSvcSaveGrantsRequest.attributeTypeMap = [
    {
        "name": "grants",
        "baseName": "grants",
        "type": "Array<UserSvcGrant>"
    }
];
