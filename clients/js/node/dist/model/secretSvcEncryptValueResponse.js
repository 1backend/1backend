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
export class SecretSvcEncryptValueResponse {
    static getAttributeTypeMap() {
        return SecretSvcEncryptValueResponse.attributeTypeMap;
    }
}
SecretSvcEncryptValueResponse.discriminator = undefined;
SecretSvcEncryptValueResponse.attributeTypeMap = [
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
