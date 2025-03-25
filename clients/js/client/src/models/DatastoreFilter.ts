/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { DatastoreOp } from './DatastoreOp';
import {
    DatastoreOpFromJSON,
    DatastoreOpFromJSONTyped,
    DatastoreOpToJSON,
    DatastoreOpToJSONTyped,
} from './DatastoreOp';

/**
 * 
 * @export
 * @interface DatastoreFilter
 */
export interface DatastoreFilter {
    /**
     * 
     * @type {Array<string>}
     * @memberof DatastoreFilter
     */
    fields?: Array<string>;
    /**
     * JSONValues is a JSON marshalled array of values.
     * It's JSON marhalled due to the limitations of the
     * Swaggo -> OpenAPI 2.0 -> OpenAPI Go generator toolchain.
     * @type {string}
     * @memberof DatastoreFilter
     */
    jsonValues?: string;
    /**
     * 
     * @type {DatastoreOp}
     * @memberof DatastoreFilter
     */
    op?: DatastoreOp;
    /**
     * SubFilters is used for operations like OR where multiple filters are combined.
     * @type {Array<DatastoreFilter>}
     * @memberof DatastoreFilter
     */
    subFilters?: Array<DatastoreFilter>;
}



/**
 * Check if a given object implements the DatastoreFilter interface.
 */
export function instanceOfDatastoreFilter(value: object): value is DatastoreFilter {
    return true;
}

export function DatastoreFilterFromJSON(json: any): DatastoreFilter {
    return DatastoreFilterFromJSONTyped(json, false);
}

export function DatastoreFilterFromJSONTyped(json: any, ignoreDiscriminator: boolean): DatastoreFilter {
    if (json == null) {
        return json;
    }
    return {
        
        'fields': json['fields'] == null ? undefined : json['fields'],
        'jsonValues': json['jsonValues'] == null ? undefined : json['jsonValues'],
        'op': json['op'] == null ? undefined : DatastoreOpFromJSON(json['op']),
        'subFilters': json['subFilters'] == null ? undefined : ((json['subFilters'] as Array<any>).map(DatastoreFilterFromJSON)),
    };
}

export function DatastoreFilterToJSON(json: any): DatastoreFilter {
    return DatastoreFilterToJSONTyped(json, false);
}

export function DatastoreFilterToJSONTyped(value?: DatastoreFilter | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'fields': value['fields'],
        'jsonValues': value['jsonValues'],
        'op': DatastoreOpToJSON(value['op']),
        'subFilters': value['subFilters'] == null ? undefined : ((value['subFilters'] as Array<any>).map(DatastoreFilterToJSON)),
    };
}

