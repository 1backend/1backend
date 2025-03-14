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
 * @interface ContainerSvcEnvVar
 */
export interface ContainerSvcEnvVar {
    /**
     * Key is the environment variable name.
     * @type {string}
     * @memberof ContainerSvcEnvVar
     */
    key: string;
    /**
     * Value is the environment variable value.
     * @type {string}
     * @memberof ContainerSvcEnvVar
     */
    value: string;
}

/**
 * Check if a given object implements the ContainerSvcEnvVar interface.
 */
export function instanceOfContainerSvcEnvVar(value: object): value is ContainerSvcEnvVar {
    if (!('key' in value) || value['key'] === undefined) return false;
    if (!('value' in value) || value['value'] === undefined) return false;
    return true;
}

export function ContainerSvcEnvVarFromJSON(json: any): ContainerSvcEnvVar {
    return ContainerSvcEnvVarFromJSONTyped(json, false);
}

export function ContainerSvcEnvVarFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcEnvVar {
    if (json == null) {
        return json;
    }
    return {
        
        'key': json['key'],
        'value': json['value'],
    };
}

export function ContainerSvcEnvVarToJSON(json: any): ContainerSvcEnvVar {
    return ContainerSvcEnvVarToJSONTyped(json, false);
}

export function ContainerSvcEnvVarToJSONTyped(value?: ContainerSvcEnvVar | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'key': value['key'],
        'value': value['value'],
    };
}

