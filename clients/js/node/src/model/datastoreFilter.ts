/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { DatastoreOp } from './datastoreOp';

export class DatastoreFilter {
    'fields'?: Array<string>;
    /**
    * JSONValues is a JSON marshalled array of values. It\'s JSON marhalled due to the limitations of the Swaggo -> OpenAPI 2.0 -> OpenAPI Go generator toolchain.
    */
    'jsonValues'?: string;
    'op'?: DatastoreOp;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
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
        }    ];

    static getAttributeTypeMap() {
        return DatastoreFilter.attributeTypeMap;
    }
}

export namespace DatastoreFilter {
}
