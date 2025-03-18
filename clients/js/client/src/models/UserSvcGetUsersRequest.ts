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
import type { DatastoreQuery } from './DatastoreQuery';
import {
    DatastoreQueryFromJSON,
    DatastoreQueryFromJSONTyped,
    DatastoreQueryToJSON,
    DatastoreQueryToJSONTyped,
} from './DatastoreQuery';

/**
 * 
 * @export
 * @interface UserSvcGetUsersRequest
 */
export interface UserSvcGetUsersRequest {
    /**
     * 
     * @type {DatastoreQuery}
     * @memberof UserSvcGetUsersRequest
     */
    query?: DatastoreQuery;
}

/**
 * Check if a given object implements the UserSvcGetUsersRequest interface.
 */
export function instanceOfUserSvcGetUsersRequest(value: object): value is UserSvcGetUsersRequest {
    return true;
}

export function UserSvcGetUsersRequestFromJSON(json: any): UserSvcGetUsersRequest {
    return UserSvcGetUsersRequestFromJSONTyped(json, false);
}

export function UserSvcGetUsersRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcGetUsersRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'query': json['query'] == null ? undefined : DatastoreQueryFromJSON(json['query']),
    };
}

export function UserSvcGetUsersRequestToJSON(json: any): UserSvcGetUsersRequest {
    return UserSvcGetUsersRequestToJSONTyped(json, false);
}

export function UserSvcGetUsersRequestToJSONTyped(value?: UserSvcGetUsersRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'query': DatastoreQueryToJSON(value['query']),
    };
}

