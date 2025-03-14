/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class SecretSvcDecryptValueResponse {
    static getAttributeTypeMap() {
        return SecretSvcDecryptValueResponse.attributeTypeMap;
    }
}
SecretSvcDecryptValueResponse.discriminator = undefined;
SecretSvcDecryptValueResponse.attributeTypeMap = [
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
