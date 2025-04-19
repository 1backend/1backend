/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { UserSvcPermitInput } from './UserSvcPermitInput';
import {
    UserSvcPermitInputFromJSON,
    UserSvcPermitInputFromJSONTyped,
    UserSvcPermitInputToJSON,
    UserSvcPermitInputToJSONTyped,
} from './UserSvcPermitInput';

/**
 * 
 * @export
 * @interface UserSvcSavePermitsRequest
 */
export interface UserSvcSavePermitsRequest {
    /**
     * 
     * @type {Array<UserSvcPermitInput>}
     * @memberof UserSvcSavePermitsRequest
     */
    permits?: Array<UserSvcPermitInput>;
}

/**
 * Check if a given object implements the UserSvcSavePermitsRequest interface.
 */
export function instanceOfUserSvcSavePermitsRequest(value: object): value is UserSvcSavePermitsRequest {
    return true;
}

export function UserSvcSavePermitsRequestFromJSON(json: any): UserSvcSavePermitsRequest {
    return UserSvcSavePermitsRequestFromJSONTyped(json, false);
}

export function UserSvcSavePermitsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSavePermitsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'permits': json['permits'] == null ? undefined : ((json['permits'] as Array<any>).map(UserSvcPermitInputFromJSON)),
    };
}

export function UserSvcSavePermitsRequestToJSON(json: any): UserSvcSavePermitsRequest {
    return UserSvcSavePermitsRequestToJSONTyped(json, false);
}

export function UserSvcSavePermitsRequestToJSONTyped(value?: UserSvcSavePermitsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'permits': value['permits'] == null ? undefined : ((value['permits'] as Array<any>).map(UserSvcPermitInputToJSON)),
    };
}

