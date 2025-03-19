/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcRegisterResponse {
    static getAttributeTypeMap() {
        return UserSvcRegisterResponse.attributeTypeMap;
    }
}
UserSvcRegisterResponse.discriminator = undefined;
UserSvcRegisterResponse.attributeTypeMap = [
    {
        "name": "token",
        "baseName": "token",
        "type": "UserSvcAuthToken"
    }
];
