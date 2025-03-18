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
export class FileSvcGetDownloadResponse {
    static getAttributeTypeMap() {
        return FileSvcGetDownloadResponse.attributeTypeMap;
    }
}
FileSvcGetDownloadResponse.discriminator = undefined;
FileSvcGetDownloadResponse.attributeTypeMap = [
    {
        "name": "download",
        "baseName": "download",
        "type": "FileSvcDownload"
    },
    {
        "name": "exists",
        "baseName": "exists",
        "type": "boolean"
    }
];
