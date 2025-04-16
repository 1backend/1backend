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
import type { SecretSvcSecret } from './SecretSvcSecret';
import {
    SecretSvcSecretFromJSON,
    SecretSvcSecretFromJSONTyped,
    SecretSvcSecretToJSON,
    SecretSvcSecretToJSONTyped,
} from './SecretSvcSecret';

/**
 * 
 * @export
 * @interface SecretSvcSaveSecretsRequest
 */
export interface SecretSvcSaveSecretsRequest {
    /**
     * 
     * @type {Array<SecretSvcSecret>}
     * @memberof SecretSvcSaveSecretsRequest
     */
    secrets?: Array<SecretSvcSecret>;
}

/**
 * Check if a given object implements the SecretSvcSaveSecretsRequest interface.
 */
export function instanceOfSecretSvcSaveSecretsRequest(value: object): value is SecretSvcSaveSecretsRequest {
    return true;
}

export function SecretSvcSaveSecretsRequestFromJSON(json: any): SecretSvcSaveSecretsRequest {
    return SecretSvcSaveSecretsRequestFromJSONTyped(json, false);
}

export function SecretSvcSaveSecretsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcSaveSecretsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'secrets': json['secrets'] == null ? undefined : ((json['secrets'] as Array<any>).map(SecretSvcSecretFromJSON)),
    };
}

export function SecretSvcSaveSecretsRequestToJSON(json: any): SecretSvcSaveSecretsRequest {
    return SecretSvcSaveSecretsRequestToJSONTyped(json, false);
}

export function SecretSvcSaveSecretsRequestToJSONTyped(value?: SecretSvcSaveSecretsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'secrets': value['secrets'] == null ? undefined : ((value['secrets'] as Array<any>).map(SecretSvcSecretToJSON)),
    };
}

