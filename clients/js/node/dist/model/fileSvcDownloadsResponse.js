/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class FileSvcDownloadsResponse {
    static getAttributeTypeMap() {
        return FileSvcDownloadsResponse.attributeTypeMap;
    }
}
FileSvcDownloadsResponse.discriminator = undefined;
FileSvcDownloadsResponse.attributeTypeMap = [
    {
        "name": "downloads",
        "baseName": "downloads",
        "type": "Array<FileSvcDownload>"
    }
];
