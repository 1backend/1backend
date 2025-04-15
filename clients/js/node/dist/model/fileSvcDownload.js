/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class FileSvcDownload {
    static getAttributeTypeMap() {
        return FileSvcDownload.attributeTypeMap;
    }
}
FileSvcDownload.discriminator = undefined;
FileSvcDownload.attributeTypeMap = [
    {
        "name": "createdAt",
        "baseName": "createdAt",
        "type": "string"
    },
    {
        "name": "downloadedBytes",
        "baseName": "downloadedBytes",
        "type": "number"
    },
    {
        "name": "error",
        "baseName": "error",
        "type": "string"
    },
    {
        "name": "fileName",
        "baseName": "fileName",
        "type": "string"
    },
    {
        "name": "filePath",
        "baseName": "filePath",
        "type": "string"
    },
    {
        "name": "fileSize",
        "baseName": "fileSize",
        "type": "number"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "progress",
        "baseName": "progress",
        "type": "number"
    },
    {
        "name": "status",
        "baseName": "status",
        "type": "string"
    },
    {
        "name": "updatedAt",
        "baseName": "updatedAt",
        "type": "string"
    },
    {
        "name": "url",
        "baseName": "url",
        "type": "string"
    }
];
