/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcListGrantsRequest {
    static getAttributeTypeMap() {
        return UserSvcListGrantsRequest.attributeTypeMap;
    }
}
UserSvcListGrantsRequest.discriminator = undefined;
UserSvcListGrantsRequest.attributeTypeMap = [
    {
        "name": "permission",
        "baseName": "permission",
        "type": "string"
    },
    {
        "name": "slug",
        "baseName": "slug",
        "type": "string"
    }
];
