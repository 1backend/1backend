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
 * @interface DataSvcCreateObjectRequest
 */
export interface DataSvcCreateObjectRequest {
    /**
     * 
     * @type {DataSvcCreateObjectFields}
     * @memberof DataSvcCreateObjectRequest
     */
    object?: DataSvcCreateObjectFields;
}

/**
 * Check if a given object implements the DataSvcCreateObjectRequest interface.
 */
export function instanceOfDataSvcCreateObjectRequest(value: object): value is DataSvcCreateObjectRequest {
    return true;
}

export function DataSvcCreateObjectRequestFromJSON(json: any): DataSvcCreateObjectRequest {
    return DataSvcCreateObjectRequestFromJSONTyped(json, false);
}

export function DataSvcCreateObjectRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcCreateObjectRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'object': json['object'] == null ? undefined : DataSvcCreateObjectFieldsFromJSON(json['object']),
    };
}

export function DataSvcCreateObjectRequestToJSON(json: any): DataSvcCreateObjectRequest {
    return DataSvcCreateObjectRequestToJSONTyped(json, false);
}

export function DataSvcCreateObjectRequestToJSONTyped(value?: DataSvcCreateObjectRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'object': DataSvcCreateObjectFieldsToJSON(value['object']),
    };
}

