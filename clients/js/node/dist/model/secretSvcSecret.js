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
export class SecretSvcSecret {
    static getAttributeTypeMap() {
        return SecretSvcSecret.attributeTypeMap;
    }
}
SecretSvcSecret.discriminator = undefined;
SecretSvcSecret.attributeTypeMap = [
    {
        "name": "app",
        "baseName": "app",
        "type": "string"
    },
    {
        "name": "canChangeDeleters",
        "baseName": "canChangeDeleters",
        "type": "Array<string>"
    },
    {
        "name": "canChangeReaders",
        "baseName": "canChangeReaders",
        "type": "Array<string>"
    },
    {
        "name": "canChangeWriters",
        "baseName": "canChangeWriters",
        "type": "Array<string>"
    },
    {
        "name": "checksum",
        "baseName": "checksum",
        "type": "string"
    },
    {
        "name": "checksumAlgorithm",
        "baseName": "checksumAlgorithm",
        "type": "SecretSvcChecksumAlgorithm"
    },
    {
        "name": "deleters",
        "baseName": "deleters",
        "type": "Array<string>"
    },
    {
        "name": "encrypted",
        "baseName": "encrypted",
        "type": "boolean"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "key",
        "baseName": "key",
        "type": "string"
    },
    {
        "name": "readers",
        "baseName": "readers",
        "type": "Array<string>"
    },
    {
        "name": "value",
        "baseName": "value",
        "type": "string"
    },
    {
        "name": "writers",
        "baseName": "writers",
        "type": "Array<string>"
    }
];
