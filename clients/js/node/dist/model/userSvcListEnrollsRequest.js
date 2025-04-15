/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcListEnrollsRequest {
    static getAttributeTypeMap() {
        return UserSvcListEnrollsRequest.attributeTypeMap;
    }
}
UserSvcListEnrollsRequest.discriminator = undefined;
UserSvcListEnrollsRequest.attributeTypeMap = [
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
