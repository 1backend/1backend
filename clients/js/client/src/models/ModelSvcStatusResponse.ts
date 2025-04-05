/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { ModelSvcModelStatus } from './ModelSvcModelStatus';
import {
    ModelSvcModelStatusFromJSON,
    ModelSvcModelStatusFromJSONTyped,
    ModelSvcModelStatusToJSON,
    ModelSvcModelStatusToJSONTyped,
} from './ModelSvcModelStatus';

/**
 * 
 * @export
 * @interface ModelSvcStatusResponse
 */
export interface ModelSvcStatusResponse {
    /**
     * 
     * @type {ModelSvcModelStatus}
     * @memberof ModelSvcStatusResponse
     */
    status?: ModelSvcModelStatus;
}

/**
 * Check if a given object implements the ModelSvcStatusResponse interface.
 */
export function instanceOfModelSvcStatusResponse(value: object): value is ModelSvcStatusResponse {
    return true;
}

export function ModelSvcStatusResponseFromJSON(json: any): ModelSvcStatusResponse {
    return ModelSvcStatusResponseFromJSONTyped(json, false);
}

export function ModelSvcStatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcStatusResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'status': json['status'] == null ? undefined : ModelSvcModelStatusFromJSON(json['status']),
    };
}

export function ModelSvcStatusResponseToJSON(json: any): ModelSvcStatusResponse {
    return ModelSvcStatusResponseToJSONTyped(json, false);
}

export function ModelSvcStatusResponseToJSONTyped(value?: ModelSvcStatusResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'status': ModelSvcModelStatusToJSON(value['status']),
    };
}

