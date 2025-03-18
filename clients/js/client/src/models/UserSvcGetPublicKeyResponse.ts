/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
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
 * @interface UserSvcGetPublicKeyResponse
 */
export interface UserSvcGetPublicKeyResponse {
    /**
     * 
     * @type {string}
     * @memberof UserSvcGetPublicKeyResponse
     */
    publicKey?: string;
}

/**
 * Check if a given object implements the UserSvcGetPublicKeyResponse interface.
 */
export function instanceOfUserSvcGetPublicKeyResponse(value: object): value is UserSvcGetPublicKeyResponse {
    return true;
}

export function UserSvcGetPublicKeyResponseFromJSON(json: any): UserSvcGetPublicKeyResponse {
    return UserSvcGetPublicKeyResponseFromJSONTyped(json, false);
}

export function UserSvcGetPublicKeyResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcGetPublicKeyResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'publicKey': json['publicKey'] == null ? undefined : json['publicKey'],
    };
}

export function UserSvcGetPublicKeyResponseToJSON(json: any): UserSvcGetPublicKeyResponse {
    return UserSvcGetPublicKeyResponseToJSONTyped(json, false);
}

export function UserSvcGetPublicKeyResponseToJSONTyped(value?: UserSvcGetPublicKeyResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'publicKey': value['publicKey'],
    };
}

