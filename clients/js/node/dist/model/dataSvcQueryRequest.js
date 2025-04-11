/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class DataSvcQueryRequest {
    static getAttributeTypeMap() {
        return DataSvcQueryRequest.attributeTypeMap;
    }
}
DataSvcQueryRequest.discriminator = undefined;
DataSvcQueryRequest.attributeTypeMap = [
    {
        "name": "query",
        "baseName": "query",
        "type": "DatastoreQuery"
    },
    {
        "name": "readers",
        "baseName": "readers",
        "type": "Array<string>"
    },
    {
        "name": "table",
        "baseName": "table",
        "type": "string"
    }
];
