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
export class UserSvcListEnrollsResponse {
    static getAttributeTypeMap() {
        return UserSvcListEnrollsResponse.attributeTypeMap;
    }
}
UserSvcListEnrollsResponse.discriminator = undefined;
UserSvcListEnrollsResponse.attributeTypeMap = [
    {
        "name": "enrolls",
        "baseName": "enrolls",
        "type": "Array<UserSvcEnroll>"
    }
];
