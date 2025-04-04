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
import type { UserSvcGrant } from './UserSvcGrant';
import {
    UserSvcGrantFromJSON,
    UserSvcGrantFromJSONTyped,
    UserSvcGrantToJSON,
    UserSvcGrantToJSONTyped,
} from './UserSvcGrant';

/**
 * 
 * @export
 * @interface UserSvcSaveGrantsRequest
 */
export interface UserSvcSaveGrantsRequest {
    /**
     * 
     * @type {Array<UserSvcGrant>}
     * @memberof UserSvcSaveGrantsRequest
     */
    grants?: Array<UserSvcGrant>;
}

/**
 * Check if a given object implements the UserSvcSaveGrantsRequest interface.
 */
export function instanceOfUserSvcSaveGrantsRequest(value: object): value is UserSvcSaveGrantsRequest {
    return true;
}

export function UserSvcSaveGrantsRequestFromJSON(json: any): UserSvcSaveGrantsRequest {
    return UserSvcSaveGrantsRequestFromJSONTyped(json, false);
}

export function UserSvcSaveGrantsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveGrantsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'grants': json['grants'] == null ? undefined : ((json['grants'] as Array<any>).map(UserSvcGrantFromJSON)),
    };
}

export function UserSvcSaveGrantsRequestToJSON(json: any): UserSvcSaveGrantsRequest {
    return UserSvcSaveGrantsRequestToJSONTyped(json, false);
}

export function UserSvcSaveGrantsRequestToJSONTyped(value?: UserSvcSaveGrantsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'grants': value['grants'] == null ? undefined : ((value['grants'] as Array<any>).map(UserSvcGrantToJSON)),
    };
}

