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
export class SecretSvcEncryptValueRequest {
    static getAttributeTypeMap() {
        return SecretSvcEncryptValueRequest.attributeTypeMap;
    }
}
SecretSvcEncryptValueRequest.discriminator = undefined;
SecretSvcEncryptValueRequest.attributeTypeMap = [
    {
        "name": "value",
        "baseName": "value",
        "type": "string"
    },
    {
        "name": "values",
        "baseName": "values",
        "type": "Array<string>"
    }
];
