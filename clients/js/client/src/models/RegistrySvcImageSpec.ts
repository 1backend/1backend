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
/**
 * 
 * @export
 * @interface RegistrySvcImageSpec
 */
export interface RegistrySvcImageSpec {
    /**
     * InternalPorts are the ports the container will listen on internally
     * @type {Array<number>}
     * @memberof RegistrySvcImageSpec
     */
    internalPorts: Array<number>;
    /**
     * Name is the container image name/URL to use for the container
     * @type {string}
     * @memberof RegistrySvcImageSpec
     */
    name: string;
}

/**
 * Check if a given object implements the RegistrySvcImageSpec interface.
 */
export function instanceOfRegistrySvcImageSpec(value: object): value is RegistrySvcImageSpec {
    if (!('internalPorts' in value) || value['internalPorts'] === undefined) return false;
    if (!('name' in value) || value['name'] === undefined) return false;
    return true;
}

export function RegistrySvcImageSpecFromJSON(json: any): RegistrySvcImageSpec {
    return RegistrySvcImageSpecFromJSONTyped(json, false);
}

export function RegistrySvcImageSpecFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcImageSpec {
    if (json == null) {
        return json;
    }
    return {
        
        'internalPorts': json['internalPorts'],
        'name': json['name'],
    };
}

export function RegistrySvcImageSpecToJSON(json: any): RegistrySvcImageSpec {
    return RegistrySvcImageSpecToJSONTyped(json, false);
}

export function RegistrySvcImageSpecToJSONTyped(value?: RegistrySvcImageSpec | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'internalPorts': value['internalPorts'],
        'name': value['name'],
    };
}

