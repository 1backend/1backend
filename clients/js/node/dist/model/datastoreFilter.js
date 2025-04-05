/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class DatastoreFilter {
    static getAttributeTypeMap() {
        return DatastoreFilter.attributeTypeMap;
    }
}
DatastoreFilter.discriminator = undefined;
DatastoreFilter.attributeTypeMap = [
    {
        "name": "fields",
        "baseName": "fields",
        "type": "Array<string>"
    },
    {
        "name": "jsonValues",
        "baseName": "jsonValues",
        "type": "string"
    },
    {
        "name": "op",
        "baseName": "op",
        "type": "DatastoreOp"
    },
    {
        "name": "subFilters",
        "baseName": "subFilters",
        "type": "Array<DatastoreFilter>"
    }
];
