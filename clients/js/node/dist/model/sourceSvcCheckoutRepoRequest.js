/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class SourceSvcCheckoutRepoRequest {
    static getAttributeTypeMap() {
        return SourceSvcCheckoutRepoRequest.attributeTypeMap;
    }
}
SourceSvcCheckoutRepoRequest.discriminator = undefined;
SourceSvcCheckoutRepoRequest.attributeTypeMap = [
    {
        "name": "password",
        "baseName": "password",
        "type": "string"
    },
    {
        "name": "ssh_key",
        "baseName": "ssh_key",
        "type": "string"
    },
    {
        "name": "ssh_key_pwd",
        "baseName": "ssh_key_pwd",
        "type": "string"
    },
    {
        "name": "token",
        "baseName": "token",
        "type": "string"
    },
    {
        "name": "url",
        "baseName": "url",
        "type": "string"
    },
    {
        "name": "username",
        "baseName": "username",
        "type": "string"
    },
    {
        "name": "version",
        "baseName": "version",
        "type": "string"
    }
];
