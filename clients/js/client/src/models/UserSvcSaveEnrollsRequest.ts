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
import type { UserSvcEnrollInput } from './UserSvcEnrollInput';
import {
    UserSvcEnrollInputFromJSON,
    UserSvcEnrollInputFromJSONTyped,
    UserSvcEnrollInputToJSON,
    UserSvcEnrollInputToJSONTyped,
} from './UserSvcEnrollInput';

/**
 * 
 * @export
 * @interface UserSvcSaveEnrollsRequest
 */
export interface UserSvcSaveEnrollsRequest {
    /**
     * 
     * @type {Array<UserSvcEnrollInput>}
     * @memberof UserSvcSaveEnrollsRequest
     */
    enrolls: Array<UserSvcEnrollInput>;
}

/**
 * Check if a given object implements the UserSvcSaveEnrollsRequest interface.
 */
export function instanceOfUserSvcSaveEnrollsRequest(value: object): value is UserSvcSaveEnrollsRequest {
    if (!('enrolls' in value) || value['enrolls'] === undefined) return false;
    return true;
}

export function UserSvcSaveEnrollsRequestFromJSON(json: any): UserSvcSaveEnrollsRequest {
    return UserSvcSaveEnrollsRequestFromJSONTyped(json, false);
}

export function UserSvcSaveEnrollsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveEnrollsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'enrolls': ((json['enrolls'] as Array<any>).map(UserSvcEnrollInputFromJSON)),
    };
}

export function UserSvcSaveEnrollsRequestToJSON(json: any): UserSvcSaveEnrollsRequest {
    return UserSvcSaveEnrollsRequestToJSONTyped(json, false);
}

export function UserSvcSaveEnrollsRequestToJSONTyped(value?: UserSvcSaveEnrollsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'enrolls': ((value['enrolls'] as Array<any>).map(UserSvcEnrollInputToJSON)),
    };
}

