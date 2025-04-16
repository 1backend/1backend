/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
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
 * @interface SecretSvcIsSecureResponse
 */
export interface SecretSvcIsSecureResponse {
    /**
     * 
     * @type {boolean}
     * @memberof SecretSvcIsSecureResponse
     */
    isSecure: boolean;
}

/**
 * Check if a given object implements the SecretSvcIsSecureResponse interface.
 */
export function instanceOfSecretSvcIsSecureResponse(value: object): value is SecretSvcIsSecureResponse {
    if (!('isSecure' in value) || value['isSecure'] === undefined) return false;
    return true;
}

export function SecretSvcIsSecureResponseFromJSON(json: any): SecretSvcIsSecureResponse {
    return SecretSvcIsSecureResponseFromJSONTyped(json, false);
}

export function SecretSvcIsSecureResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcIsSecureResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'isSecure': json['isSecure'],
    };
}

export function SecretSvcIsSecureResponseToJSON(json: any): SecretSvcIsSecureResponse {
    return SecretSvcIsSecureResponseToJSONTyped(json, false);
}

export function SecretSvcIsSecureResponseToJSONTyped(value?: SecretSvcIsSecureResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'isSecure': value['isSecure'],
    };
}

