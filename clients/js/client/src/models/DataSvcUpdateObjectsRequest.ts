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
import type { DatastoreFilter } from './DatastoreFilter';
import {
    DatastoreFilterFromJSON,
    DatastoreFilterFromJSONTyped,
    DatastoreFilterToJSON,
    DatastoreFilterToJSONTyped,
} from './DatastoreFilter';
import type { DataSvcObject } from './DataSvcObject';
import {
    DataSvcObjectFromJSON,
    DataSvcObjectFromJSONTyped,
    DataSvcObjectToJSON,
    DataSvcObjectToJSONTyped,
} from './DataSvcObject';

/**
 * 
 * @export
 * @interface DataSvcUpdateObjectsRequest
 */
export interface DataSvcUpdateObjectsRequest {
    /**
     * Filters to determine which objects will be updated.
     * Only objects matching all filters will be modified.
     * @type {Array<DatastoreFilter>}
     * @memberof DataSvcUpdateObjectsRequest
     */
    filters?: Array<DatastoreFilter>;
    /**
     * The object containing the fields to update in matching objects.
     * @type {DataSvcObject}
     * @memberof DataSvcUpdateObjectsRequest
     */
    object?: DataSvcObject;
    /**
     * 
     * @type {string}
     * @memberof DataSvcUpdateObjectsRequest
     */
    table?: string;
}

/**
 * Check if a given object implements the DataSvcUpdateObjectsRequest interface.
 */
export function instanceOfDataSvcUpdateObjectsRequest(value: object): value is DataSvcUpdateObjectsRequest {
    return true;
}

export function DataSvcUpdateObjectsRequestFromJSON(json: any): DataSvcUpdateObjectsRequest {
    return DataSvcUpdateObjectsRequestFromJSONTyped(json, false);
}

export function DataSvcUpdateObjectsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcUpdateObjectsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'filters': json['filters'] == null ? undefined : ((json['filters'] as Array<any>).map(DatastoreFilterFromJSON)),
        'object': json['object'] == null ? undefined : DataSvcObjectFromJSON(json['object']),
        'table': json['table'] == null ? undefined : json['table'],
    };
}

export function DataSvcUpdateObjectsRequestToJSON(json: any): DataSvcUpdateObjectsRequest {
    return DataSvcUpdateObjectsRequestToJSONTyped(json, false);
}

export function DataSvcUpdateObjectsRequestToJSONTyped(value?: DataSvcUpdateObjectsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'filters': value['filters'] == null ? undefined : ((value['filters'] as Array<any>).map(DatastoreFilterToJSON)),
        'object': DataSvcObjectToJSON(value['object']),
        'table': value['table'],
    };
}

