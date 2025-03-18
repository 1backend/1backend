/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class SecretSvcRemoveSecretsRequest {
    static getAttributeTypeMap() {
        return SecretSvcRemoveSecretsRequest.attributeTypeMap;
    }
}
SecretSvcRemoveSecretsRequest.discriminator = undefined;
SecretSvcRemoveSecretsRequest.attributeTypeMap = [
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "ids",
        "baseName": "ids",
        "type": "Array<string>"
    },
    {
        "name": "key",
        "baseName": "key",
        "type": "string"
    },
    {
        "name": "keys",
        "baseName": "keys",
        "type": "Array<string>"
    }
];
