/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class DatastoreQuery {
    static getAttributeTypeMap() {
        return DatastoreQuery.attributeTypeMap;
    }
}
DatastoreQuery.discriminator = undefined;
DatastoreQuery.attributeTypeMap = [
    {
        "name": "count",
        "baseName": "count",
        "type": "boolean"
    },
    {
        "name": "filters",
        "baseName": "filters",
        "type": "Array<DatastoreFilter>"
    },
    {
        "name": "jsonAfter",
        "baseName": "jsonAfter",
        "type": "string"
    },
    {
        "name": "limit",
        "baseName": "limit",
        "type": "number"
    },
    {
        "name": "orderBys",
        "baseName": "orderBys",
        "type": "Array<DatastoreOrderBy>"
    }
];
