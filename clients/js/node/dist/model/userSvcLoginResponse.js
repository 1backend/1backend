/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcLoginResponse {
    static getAttributeTypeMap() {
        return UserSvcLoginResponse.attributeTypeMap;
    }
}
UserSvcLoginResponse.discriminator = undefined;
UserSvcLoginResponse.attributeTypeMap = [
    {
        "name": "token",
        "baseName": "token",
        "type": "UserSvcAuthToken"
    }
];
