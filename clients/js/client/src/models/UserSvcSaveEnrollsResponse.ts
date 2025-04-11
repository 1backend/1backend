/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { UserSvcEnroll } from './UserSvcEnroll';
import {
    UserSvcEnrollFromJSON,
    UserSvcEnrollFromJSONTyped,
    UserSvcEnrollToJSON,
    UserSvcEnrollToJSONTyped,
} from './UserSvcEnroll';

/**
 * 
 * @export
 * @interface UserSvcSaveEnrollsResponse
 */
export interface UserSvcSaveEnrollsResponse {
    /**
     * 
     * @type {Array<UserSvcEnroll>}
     * @memberof UserSvcSaveEnrollsResponse
     */
    enrolls: Array<UserSvcEnroll>;
}

/**
 * Check if a given object implements the UserSvcSaveEnrollsResponse interface.
 */
export function instanceOfUserSvcSaveEnrollsResponse(value: object): value is UserSvcSaveEnrollsResponse {
    if (!('enrolls' in value) || value['enrolls'] === undefined) return false;
    return true;
}

export function UserSvcSaveEnrollsResponseFromJSON(json: any): UserSvcSaveEnrollsResponse {
    return UserSvcSaveEnrollsResponseFromJSONTyped(json, false);
}

export function UserSvcSaveEnrollsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveEnrollsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'enrolls': ((json['enrolls'] as Array<any>).map(UserSvcEnrollFromJSON)),
    };
}

export function UserSvcSaveEnrollsResponseToJSON(json: any): UserSvcSaveEnrollsResponse {
    return UserSvcSaveEnrollsResponseToJSONTyped(json, false);
}

export function UserSvcSaveEnrollsResponseToJSONTyped(value?: UserSvcSaveEnrollsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'enrolls': ((value['enrolls'] as Array<any>).map(UserSvcEnrollToJSON)),
    };
}

