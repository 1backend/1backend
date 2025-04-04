/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcListUsersResponse {
    static getAttributeTypeMap() {
        return UserSvcListUsersResponse.attributeTypeMap;
    }
}
UserSvcListUsersResponse.discriminator = undefined;
UserSvcListUsersResponse.attributeTypeMap = [
    {
        "name": "after",
        "baseName": "after",
        "type": "string"
    },
    {
        "name": "count",
        "baseName": "count",
        "type": "number"
    },
    {
        "name": "users",
        "baseName": "users",
        "type": "Array<UserSvcUserRecord>"
    }
];
