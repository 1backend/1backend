/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ModelSvcEnvVar
 */
export interface ModelSvcEnvVar {
    /**
     * Key is the environment variable name.
     * @type {string}
     * @memberof ModelSvcEnvVar
     */
    key?: string;
    /**
     * Value is the environment variable value.
     * @type {string}
     * @memberof ModelSvcEnvVar
     */
    value?: string;
}

/**
 * Check if a given object implements the ModelSvcEnvVar interface.
 */
export function instanceOfModelSvcEnvVar(value: object): value is ModelSvcEnvVar {
    return true;
}

export function ModelSvcEnvVarFromJSON(json: any): ModelSvcEnvVar {
    return ModelSvcEnvVarFromJSONTyped(json, false);
}

export function ModelSvcEnvVarFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcEnvVar {
    if (json == null) {
        return json;
    }
    return {
        
        'key': json['key'] == null ? undefined : json['key'],
        'value': json['value'] == null ? undefined : json['value'],
    };
}

export function ModelSvcEnvVarToJSON(json: any): ModelSvcEnvVar {
    return ModelSvcEnvVarToJSONTyped(json, false);
}

export function ModelSvcEnvVarToJSONTyped(value?: ModelSvcEnvVar | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'key': value['key'],
        'value': value['value'],
    };
}

