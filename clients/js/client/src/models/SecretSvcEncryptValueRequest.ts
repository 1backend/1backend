/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
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
 * @interface SecretSvcEncryptValueRequest
 */
export interface SecretSvcEncryptValueRequest {
    /**
     * 
     * @type {string}
     * @memberof SecretSvcEncryptValueRequest
     */
    value?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof SecretSvcEncryptValueRequest
     */
    values?: Array<string>;
}

/**
 * Check if a given object implements the SecretSvcEncryptValueRequest interface.
 */
export function instanceOfSecretSvcEncryptValueRequest(value: object): value is SecretSvcEncryptValueRequest {
    return true;
}

export function SecretSvcEncryptValueRequestFromJSON(json: any): SecretSvcEncryptValueRequest {
    return SecretSvcEncryptValueRequestFromJSONTyped(json, false);
}

export function SecretSvcEncryptValueRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcEncryptValueRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'value': json['value'] == null ? undefined : json['value'],
        'values': json['values'] == null ? undefined : json['values'],
    };
}

export function SecretSvcEncryptValueRequestToJSON(json: any): SecretSvcEncryptValueRequest {
    return SecretSvcEncryptValueRequestToJSONTyped(json, false);
}

export function SecretSvcEncryptValueRequestToJSONTyped(value?: SecretSvcEncryptValueRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'value': value['value'],
        'values': value['values'],
    };
}

