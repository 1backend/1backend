/* tslint:disable */
/* eslint-disable */
/**
 * Basic Svc
 * An example service for bootstrapping.
 *
 * The version of the OpenAPI document: 0.3.0-rc.8
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
 * @interface BasicSvcSavePetRequest
 */
export interface BasicSvcSavePetRequest {
    /**
     * 
     * @type {string}
     * @memberof BasicSvcSavePetRequest
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof BasicSvcSavePetRequest
     */
    name: string;
}

/**
 * Check if a given object implements the BasicSvcSavePetRequest interface.
 */
export function instanceOfBasicSvcSavePetRequest(value: object): value is BasicSvcSavePetRequest {
    if (!('name' in value) || value['name'] === undefined) return false;
    return true;
}

export function BasicSvcSavePetRequestFromJSON(json: any): BasicSvcSavePetRequest {
    return BasicSvcSavePetRequestFromJSONTyped(json, false);
}

export function BasicSvcSavePetRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): BasicSvcSavePetRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'],
    };
}

export function BasicSvcSavePetRequestToJSON(json: any): BasicSvcSavePetRequest {
    return BasicSvcSavePetRequestToJSONTyped(json, false);
}

export function BasicSvcSavePetRequestToJSONTyped(value?: BasicSvcSavePetRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'name': value['name'],
    };
}

