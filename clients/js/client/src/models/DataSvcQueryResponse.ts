/* tslint:disable */
/* eslint-disable */
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
 * @interface DataSvcQueryResponse
 */
export interface DataSvcQueryResponse {
    /**
     * 
     * @type {Array<DataSvcObject>}
     * @memberof DataSvcQueryResponse
     */
    objects?: Array<DataSvcObject>;
}

/**
 * Check if a given object implements the DataSvcQueryResponse interface.
 */
export function instanceOfDataSvcQueryResponse(value: object): value is DataSvcQueryResponse {
    return true;
}

export function DataSvcQueryResponseFromJSON(json: any): DataSvcQueryResponse {
    return DataSvcQueryResponseFromJSONTyped(json, false);
}

export function DataSvcQueryResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcQueryResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'objects': json['objects'] == null ? undefined : ((json['objects'] as Array<any>).map(DataSvcObjectFromJSON)),
    };
}

export function DataSvcQueryResponseToJSON(json: any): DataSvcQueryResponse {
    return DataSvcQueryResponseToJSONTyped(json, false);
}

export function DataSvcQueryResponseToJSONTyped(value?: DataSvcQueryResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'objects': value['objects'] == null ? undefined : ((value['objects'] as Array<any>).map(DataSvcObjectToJSON)),
    };
}

