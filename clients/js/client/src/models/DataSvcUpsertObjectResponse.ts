/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
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
 * @interface DataSvcUpsertObjectResponse
 */
export interface DataSvcUpsertObjectResponse {
    /**
     * 
     * @type {DataSvcObject}
     * @memberof DataSvcUpsertObjectResponse
     */
    object?: DataSvcObject;
}

/**
 * Check if a given object implements the DataSvcUpsertObjectResponse interface.
 */
export function instanceOfDataSvcUpsertObjectResponse(value: object): value is DataSvcUpsertObjectResponse {
    return true;
}

export function DataSvcUpsertObjectResponseFromJSON(json: any): DataSvcUpsertObjectResponse {
    return DataSvcUpsertObjectResponseFromJSONTyped(json, false);
}

export function DataSvcUpsertObjectResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcUpsertObjectResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'object': json['object'] == null ? undefined : DataSvcObjectFromJSON(json['object']),
    };
}

export function DataSvcUpsertObjectResponseToJSON(json: any): DataSvcUpsertObjectResponse {
    return DataSvcUpsertObjectResponseToJSONTyped(json, false);
}

export function DataSvcUpsertObjectResponseToJSONTyped(value?: DataSvcUpsertObjectResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'object': DataSvcObjectToJSON(value['object']),
    };
}

