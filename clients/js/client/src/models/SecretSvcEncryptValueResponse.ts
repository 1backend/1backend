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
/**
 * 
 * @export
 * @interface SecretSvcEncryptValueResponse
 */
export interface SecretSvcEncryptValueResponse {
    /**
     * 
     * @type {string}
     * @memberof SecretSvcEncryptValueResponse
     */
    value?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof SecretSvcEncryptValueResponse
     */
    values?: Array<string>;
}

/**
 * Check if a given object implements the SecretSvcEncryptValueResponse interface.
 */
export function instanceOfSecretSvcEncryptValueResponse(value: object): value is SecretSvcEncryptValueResponse {
    return true;
}

export function SecretSvcEncryptValueResponseFromJSON(json: any): SecretSvcEncryptValueResponse {
    return SecretSvcEncryptValueResponseFromJSONTyped(json, false);
}

export function SecretSvcEncryptValueResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcEncryptValueResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'value': json['value'] == null ? undefined : json['value'],
        'values': json['values'] == null ? undefined : json['values'],
    };
}

export function SecretSvcEncryptValueResponseToJSON(json: any): SecretSvcEncryptValueResponse {
    return SecretSvcEncryptValueResponseToJSONTyped(json, false);
}

export function SecretSvcEncryptValueResponseToJSONTyped(value?: SecretSvcEncryptValueResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'value': value['value'],
        'values': value['values'],
    };
}

