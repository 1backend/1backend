/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcContactInput {
    static getAttributeTypeMap() {
        return UserSvcContactInput.attributeTypeMap;
    }
}
UserSvcContactInput.discriminator = undefined;
UserSvcContactInput.attributeTypeMap = [
    {
        "name": "handle",
        "baseName": "handle",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "platform",
        "baseName": "platform",
        "type": "string"
    }
];
