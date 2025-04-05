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
 * @interface SecretSvcListSecretsResponse
 */
export interface SecretSvcListSecretsResponse {
    /**
     * 
     * @type {Array<SecretSvcSecret>}
     * @memberof SecretSvcListSecretsResponse
     */
    secrets?: Array<SecretSvcSecret>;
}

/**
 * Check if a given object implements the SecretSvcListSecretsResponse interface.
 */
export function instanceOfSecretSvcListSecretsResponse(value: object): value is SecretSvcListSecretsResponse {
    return true;
}

export function SecretSvcListSecretsResponseFromJSON(json: any): SecretSvcListSecretsResponse {
    return SecretSvcListSecretsResponseFromJSONTyped(json, false);
}

export function SecretSvcListSecretsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcListSecretsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'secrets': json['secrets'] == null ? undefined : ((json['secrets'] as Array<any>).map(SecretSvcSecretFromJSON)),
    };
}

export function SecretSvcListSecretsResponseToJSON(json: any): SecretSvcListSecretsResponse {
    return SecretSvcListSecretsResponseToJSONTyped(json, false);
}

export function SecretSvcListSecretsResponseToJSONTyped(value?: SecretSvcListSecretsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'secrets': value['secrets'] == null ? undefined : ((value['secrets'] as Array<any>).map(SecretSvcSecretToJSON)),
    };
}

