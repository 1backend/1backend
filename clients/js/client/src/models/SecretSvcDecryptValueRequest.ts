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
/**
 * 
 * @export
 * @interface SecretSvcDecryptValueRequest
 */
export interface SecretSvcDecryptValueRequest {
    /**
     * 
     * @type {string}
     * @memberof SecretSvcDecryptValueRequest
     */
    value?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof SecretSvcDecryptValueRequest
     */
    values?: Array<string>;
}

/**
 * Check if a given object implements the SecretSvcDecryptValueRequest interface.
 */
export function instanceOfSecretSvcDecryptValueRequest(value: object): value is SecretSvcDecryptValueRequest {
    return true;
}

export function SecretSvcDecryptValueRequestFromJSON(json: any): SecretSvcDecryptValueRequest {
    return SecretSvcDecryptValueRequestFromJSONTyped(json, false);
}

export function SecretSvcDecryptValueRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcDecryptValueRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'value': json['value'] == null ? undefined : json['value'],
        'values': json['values'] == null ? undefined : json['values'],
    };
}

export function SecretSvcDecryptValueRequestToJSON(json: any): SecretSvcDecryptValueRequest {
    return SecretSvcDecryptValueRequestToJSONTyped(json, false);
}

export function SecretSvcDecryptValueRequestToJSONTyped(value?: SecretSvcDecryptValueRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'value': value['value'],
        'values': value['values'],
    };
}

