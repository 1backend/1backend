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

import { RequestFile } from './models';
import { DatastoreFilter } from './datastoreFilter';

export class DataSvcDeleteObjectRequest {
    'filters'?: Array<DatastoreFilter>;
    'table'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "filters",
            "baseName": "filters",
            "type": "Array<DatastoreFilter>"
        },
        {
            "name": "table",
            "baseName": "table",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return DataSvcDeleteObjectRequest.attributeTypeMap;
    }
}

