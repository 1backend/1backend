/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
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
 * @interface RegistrySvcEnvVar
 */
export interface RegistrySvcEnvVar {
    /**
     * Key is the environment variable name.
     * @type {string}
     * @memberof RegistrySvcEnvVar
     */
    key: string;
    /**
     * Value is the environment variable value.
     * @type {string}
     * @memberof RegistrySvcEnvVar
     */
    value: string;
}

/**
 * Check if a given object implements the RegistrySvcEnvVar interface.
 */
export function instanceOfRegistrySvcEnvVar(value: object): value is RegistrySvcEnvVar {
    if (!('key' in value) || value['key'] === undefined) return false;
    if (!('value' in value) || value['value'] === undefined) return false;
    return true;
}

export function RegistrySvcEnvVarFromJSON(json: any): RegistrySvcEnvVar {
    return RegistrySvcEnvVarFromJSONTyped(json, false);
}

export function RegistrySvcEnvVarFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcEnvVar {
    if (json == null) {
        return json;
    }
    return {
        
        'key': json['key'],
        'value': json['value'],
    };
}

export function RegistrySvcEnvVarToJSON(json: any): RegistrySvcEnvVar {
    return RegistrySvcEnvVarToJSONTyped(json, false);
}

export function RegistrySvcEnvVarToJSONTyped(value?: RegistrySvcEnvVar | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'key': value['key'],
        'value': value['value'],
    };
}

