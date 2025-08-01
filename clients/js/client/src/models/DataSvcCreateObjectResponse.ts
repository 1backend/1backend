/* tslint:disable */
/* eslint-disable */
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
 * @interface DataSvcCreateObjectResponse
 */
export interface DataSvcCreateObjectResponse {
    /**
     * 
     * @type {DataSvcObject}
     * @memberof DataSvcCreateObjectResponse
     */
    object?: DataSvcObject;
}

/**
 * Check if a given object implements the DataSvcCreateObjectResponse interface.
 */
export function instanceOfDataSvcCreateObjectResponse(value: object): value is DataSvcCreateObjectResponse {
    return true;
}

export function DataSvcCreateObjectResponseFromJSON(json: any): DataSvcCreateObjectResponse {
    return DataSvcCreateObjectResponseFromJSONTyped(json, false);
}

export function DataSvcCreateObjectResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcCreateObjectResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'object': json['object'] == null ? undefined : DataSvcObjectFromJSON(json['object']),
    };
}

export function DataSvcCreateObjectResponseToJSON(json: any): DataSvcCreateObjectResponse {
    return DataSvcCreateObjectResponseToJSONTyped(json, false);
}

export function DataSvcCreateObjectResponseToJSONTyped(value?: DataSvcCreateObjectResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'object': DataSvcObjectToJSON(value['object']),
    };
}

