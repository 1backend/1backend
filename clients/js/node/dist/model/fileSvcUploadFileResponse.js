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
export class FileSvcUploadFileResponse {
    static getAttributeTypeMap() {
        return FileSvcUploadFileResponse.attributeTypeMap;
    }
}
FileSvcUploadFileResponse.discriminator = undefined;
FileSvcUploadFileResponse.attributeTypeMap = [
    {
        "name": "upload",
        "baseName": "upload",
        "type": "FileSvcUpload"
    }
];
