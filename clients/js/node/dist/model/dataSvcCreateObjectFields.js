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
export class DataSvcCreateObjectFields {
    static getAttributeTypeMap() {
        return DataSvcCreateObjectFields.attributeTypeMap;
    }
}
DataSvcCreateObjectFields.discriminator = undefined;
DataSvcCreateObjectFields.attributeTypeMap = [
    {
        "name": "authors",
        "baseName": "authors",
        "type": "Array<string>"
    },
    {
        "name": "data",
        "baseName": "data",
        "type": "{ [key: string]: any; }"
    },
    {
        "name": "deleters",
        "baseName": "deleters",
        "type": "Array<string>"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
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
    },
    {
        "name": "writers",
        "baseName": "writers",
        "type": "Array<string>"
    }
];
