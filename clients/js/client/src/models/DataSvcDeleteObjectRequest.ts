/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { DatastoreFilter } from './DatastoreFilter';
import {
    DatastoreFilterFromJSON,
    DatastoreFilterFromJSONTyped,
    DatastoreFilterToJSON,
    DatastoreFilterToJSONTyped,
} from './DatastoreFilter';

/**
 * 
 * @export
 * @interface DataSvcDeleteObjectRequest
 */
export interface DataSvcDeleteObjectRequest {
    /**
     * 
     * @type {Array<DatastoreFilter>}
     * @memberof DataSvcDeleteObjectRequest
     */
    filters?: Array<DatastoreFilter>;
    /**
     * 
     * @type {string}
     * @memberof DataSvcDeleteObjectRequest
     */
    table?: string;
}

/**
 * Check if a given object implements the DataSvcDeleteObjectRequest interface.
 */
export function instanceOfDataSvcDeleteObjectRequest(value: object): value is DataSvcDeleteObjectRequest {
    return true;
}

export function DataSvcDeleteObjectRequestFromJSON(json: any): DataSvcDeleteObjectRequest {
    return DataSvcDeleteObjectRequestFromJSONTyped(json, false);
}

export function DataSvcDeleteObjectRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcDeleteObjectRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'filters': json['filters'] == null ? undefined : ((json['filters'] as Array<any>).map(DatastoreFilterFromJSON)),
        'table': json['table'] == null ? undefined : json['table'],
    };
}

export function DataSvcDeleteObjectRequestToJSON(json: any): DataSvcDeleteObjectRequest {
    return DataSvcDeleteObjectRequestToJSONTyped(json, false);
}

export function DataSvcDeleteObjectRequestToJSONTyped(value?: DataSvcDeleteObjectRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'filters': value['filters'] == null ? undefined : ((value['filters'] as Array<any>).map(DatastoreFilterToJSON)),
        'table': value['table'],
    };
}

