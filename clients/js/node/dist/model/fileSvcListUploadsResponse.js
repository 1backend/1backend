/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class FileSvcListUploadsResponse {
    static getAttributeTypeMap() {
        return FileSvcListUploadsResponse.attributeTypeMap;
    }
}
FileSvcListUploadsResponse.discriminator = undefined;
FileSvcListUploadsResponse.attributeTypeMap = [
    {
        "name": "uploads",
        "baseName": "uploads",
        "type": "Array<FileSvcUpload>"
    }
];
