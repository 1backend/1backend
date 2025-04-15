/* tslint:disable */
/* eslint-disable */
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

import { mapValues } from '../runtime';
import type { DataSvcCreateObjectFields } from './DataSvcCreateObjectFields';
import {
    DataSvcCreateObjectFieldsFromJSON,
    DataSvcCreateObjectFieldsFromJSONTyped,
    DataSvcCreateObjectFieldsToJSON,
    DataSvcCreateObjectFieldsToJSONTyped,
} from './DataSvcCreateObjectFields';

/**
 * 
 * @export
 * @interface DataSvcUpsertObjectRequest
 */
export interface DataSvcUpsertObjectRequest {
    /**
     * 
     * @type {DataSvcCreateObjectFields}
     * @memberof DataSvcUpsertObjectRequest
     */
    object?: DataSvcCreateObjectFields;
}

/**
 * Check if a given object implements the DataSvcUpsertObjectRequest interface.
 */
export function instanceOfDataSvcUpsertObjectRequest(value: object): value is DataSvcUpsertObjectRequest {
    return true;
}

export function DataSvcUpsertObjectRequestFromJSON(json: any): DataSvcUpsertObjectRequest {
    return DataSvcUpsertObjectRequestFromJSONTyped(json, false);
}

export function DataSvcUpsertObjectRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcUpsertObjectRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'object': json['object'] == null ? undefined : DataSvcCreateObjectFieldsFromJSON(json['object']),
    };
}

export function DataSvcUpsertObjectRequestToJSON(json: any): DataSvcUpsertObjectRequest {
    return DataSvcUpsertObjectRequestToJSONTyped(json, false);
}

export function DataSvcUpsertObjectRequestToJSONTyped(value?: DataSvcUpsertObjectRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'object': DataSvcCreateObjectFieldsToJSON(value['object']),
    };
}

