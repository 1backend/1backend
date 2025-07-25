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
 * @interface SecretSvcListSecretsRequest
 */
export interface SecretSvcListSecretsRequest {
    /**
     * 
     * @type {string}
     * @memberof SecretSvcListSecretsRequest
     */
    key?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof SecretSvcListSecretsRequest
     */
    keys?: Array<string>;
}

/**
 * Check if a given object implements the SecretSvcListSecretsRequest interface.
 */
export function instanceOfSecretSvcListSecretsRequest(value: object): value is SecretSvcListSecretsRequest {
    return true;
}

export function SecretSvcListSecretsRequestFromJSON(json: any): SecretSvcListSecretsRequest {
    return SecretSvcListSecretsRequestFromJSONTyped(json, false);
}

export function SecretSvcListSecretsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcListSecretsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'key': json['key'] == null ? undefined : json['key'],
        'keys': json['keys'] == null ? undefined : json['keys'],
    };
}

export function SecretSvcListSecretsRequestToJSON(json: any): SecretSvcListSecretsRequest {
    return SecretSvcListSecretsRequestToJSONTyped(json, false);
}

export function SecretSvcListSecretsRequestToJSONTyped(value?: SecretSvcListSecretsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'key': value['key'],
        'keys': value['keys'],
    };
}

