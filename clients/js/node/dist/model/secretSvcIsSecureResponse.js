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
export class SecretSvcIsSecureResponse {
    static getAttributeTypeMap() {
        return SecretSvcIsSecureResponse.attributeTypeMap;
    }
}
SecretSvcIsSecureResponse.discriminator = undefined;
SecretSvcIsSecureResponse.attributeTypeMap = [
    {
        "name": "isSecure",
        "baseName": "isSecure",
        "type": "boolean"
    }
];
