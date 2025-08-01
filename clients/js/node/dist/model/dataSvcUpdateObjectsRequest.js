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
export class DataSvcUpdateObjectsRequest {
    static getAttributeTypeMap() {
        return DataSvcUpdateObjectsRequest.attributeTypeMap;
    }
}
DataSvcUpdateObjectsRequest.discriminator = undefined;
DataSvcUpdateObjectsRequest.attributeTypeMap = [
    {
        "name": "filters",
        "baseName": "filters",
        "type": "Array<DatastoreFilter>"
    },
    {
        "name": "object",
        "baseName": "object",
        "type": "DataSvcObject"
    },
    {
        "name": "table",
        "baseName": "table",
        "type": "string"
    }
];
