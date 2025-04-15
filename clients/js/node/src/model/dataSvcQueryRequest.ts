/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { DatastoreQuery } from './datastoreQuery';

export class DataSvcQueryRequest {
    'query'?: DatastoreQuery;
    'readers'?: Array<string>;
    'table'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
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
        }    ];

    static getAttributeTypeMap() {
        return DataSvcQueryRequest.attributeTypeMap;
    }
}

