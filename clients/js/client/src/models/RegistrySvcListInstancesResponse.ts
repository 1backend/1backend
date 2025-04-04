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
import type { RegistrySvcInstance } from './RegistrySvcInstance';
import {
    RegistrySvcInstanceFromJSON,
    RegistrySvcInstanceFromJSONTyped,
    RegistrySvcInstanceToJSON,
    RegistrySvcInstanceToJSONTyped,
} from './RegistrySvcInstance';

/**
 * 
 * @export
 * @interface RegistrySvcListInstancesResponse
 */
export interface RegistrySvcListInstancesResponse {
    /**
     * 
     * @type {Array<RegistrySvcInstance>}
     * @memberof RegistrySvcListInstancesResponse
     */
    instances?: Array<RegistrySvcInstance>;
}

/**
 * Check if a given object implements the RegistrySvcListInstancesResponse interface.
 */
export function instanceOfRegistrySvcListInstancesResponse(value: object): value is RegistrySvcListInstancesResponse {
    return true;
}

export function RegistrySvcListInstancesResponseFromJSON(json: any): RegistrySvcListInstancesResponse {
    return RegistrySvcListInstancesResponseFromJSONTyped(json, false);
}

export function RegistrySvcListInstancesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcListInstancesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'instances': json['instances'] == null ? undefined : ((json['instances'] as Array<any>).map(RegistrySvcInstanceFromJSON)),
    };
}

export function RegistrySvcListInstancesResponseToJSON(json: any): RegistrySvcListInstancesResponse {
    return RegistrySvcListInstancesResponseToJSONTyped(json, false);
}

export function RegistrySvcListInstancesResponseToJSONTyped(value?: RegistrySvcListInstancesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'instances': value['instances'] == null ? undefined : ((value['instances'] as Array<any>).map(RegistrySvcInstanceToJSON)),
    };
}

