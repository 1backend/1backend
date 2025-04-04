/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { DatastoreQuery } from './DatastoreQuery';
import {
    DatastoreQueryFromJSON,
    DatastoreQueryFromJSONTyped,
    DatastoreQueryToJSON,
    DatastoreQueryToJSONTyped,
} from './DatastoreQuery';

/**
 * 
 * @export
 * @interface DataSvcQueryRequest
 */
export interface DataSvcQueryRequest {
    /**
     * 
     * @type {DatastoreQuery}
     * @memberof DataSvcQueryRequest
     */
    query?: DatastoreQuery;
    /**
     * 
     * @type {Array<string>}
     * @memberof DataSvcQueryRequest
     */
    readers?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof DataSvcQueryRequest
     */
    table?: string;
}

/**
 * Check if a given object implements the DataSvcQueryRequest interface.
 */
export function instanceOfDataSvcQueryRequest(value: object): value is DataSvcQueryRequest {
    return true;
}

export function DataSvcQueryRequestFromJSON(json: any): DataSvcQueryRequest {
    return DataSvcQueryRequestFromJSONTyped(json, false);
}

export function DataSvcQueryRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcQueryRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'query': json['query'] == null ? undefined : DatastoreQueryFromJSON(json['query']),
        'readers': json['readers'] == null ? undefined : json['readers'],
        'table': json['table'] == null ? undefined : json['table'],
    };
}

export function DataSvcQueryRequestToJSON(json: any): DataSvcQueryRequest {
    return DataSvcQueryRequestToJSONTyped(json, false);
}

export function DataSvcQueryRequestToJSONTyped(value?: DataSvcQueryRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'query': DatastoreQueryToJSON(value['query']),
        'readers': value['readers'],
        'table': value['table'],
    };
}

