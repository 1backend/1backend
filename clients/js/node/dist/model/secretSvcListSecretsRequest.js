/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class SecretSvcListSecretsRequest {
    static getAttributeTypeMap() {
        return SecretSvcListSecretsRequest.attributeTypeMap;
    }
}
SecretSvcListSecretsRequest.discriminator = undefined;
SecretSvcListSecretsRequest.attributeTypeMap = [
    {
        "name": "key",
        "baseName": "key",
        "type": "string"
    },
    {
        "name": "keys",
        "baseName": "keys",
        "type": "Array<string>"
    },
    {
        "name": "namespace",
        "baseName": "namespace",
        "type": "string"
    }
];
