/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class FileSvcDownloadFileRequest {
    static getAttributeTypeMap() {
        return FileSvcDownloadFileRequest.attributeTypeMap;
    }
}
FileSvcDownloadFileRequest.discriminator = undefined;
FileSvcDownloadFileRequest.attributeTypeMap = [
    {
        "name": "folderPath",
        "baseName": "folderPath",
        "type": "string"
    },
    {
        "name": "url",
        "baseName": "url",
        "type": "string"
    }
];
