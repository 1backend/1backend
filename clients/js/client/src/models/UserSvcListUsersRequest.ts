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
import type { UserSvcOrderDirection } from './UserSvcOrderDirection';
import {
    UserSvcOrderDirectionFromJSON,
    UserSvcOrderDirectionFromJSONTyped,
    UserSvcOrderDirectionToJSON,
    UserSvcOrderDirectionToJSONTyped,
} from './UserSvcOrderDirection';
import type { UserSvcListUsersOrderBy } from './UserSvcListUsersOrderBy';
import {
    UserSvcListUsersOrderByFromJSON,
    UserSvcListUsersOrderByFromJSONTyped,
    UserSvcListUsersOrderByToJSON,
    UserSvcListUsersOrderByToJSONTyped,
} from './UserSvcListUsersOrderBy';

/**
 * 
 * @export
 * @interface UserSvcListUsersRequest
 */
export interface UserSvcListUsersRequest {
    /**
     * AfterTime is a time in RFC3339 format.
     * It is used to paginate the results when the `orderByField` is set to `createdAt` or `updatedAt`.
     * The results will be returned after this time.
     * @type {string}
     * @memberof UserSvcListUsersRequest
     */
    afterTime?: string;
    /**
     * ContactId is the id of the contact the user is associated with.
     * Will return a user list with one element if set.
     * @type {string}
     * @memberof UserSvcListUsersRequest
     */
    contactId?: string;
    /**
     * Count is a flag that indicates if the count of the users should be returned.
     * @type {boolean}
     * @memberof UserSvcListUsersRequest
     */
    count?: boolean;
    /**
     * Ids of the users to list.
     * @type {Array<string>}
     * @memberof UserSvcListUsersRequest
     */
    ids?: Array<string>;
    /**
     * 
     * @type {number}
     * @memberof UserSvcListUsersRequest
     */
    limit?: number;
    /**
     * 
     * @type {UserSvcOrderDirection}
     * @memberof UserSvcListUsersRequest
     */
    order?: UserSvcOrderDirection;
    /**
     * 
     * @type {UserSvcListUsersOrderBy}
     * @memberof UserSvcListUsersRequest
     */
    orderBy?: UserSvcListUsersOrderBy;
    /**
     * Search term used to find users. Returns users whose slug, username, or contact ID exactly matches the term.
     * @type {string}
     * @memberof UserSvcListUsersRequest
     */
    search?: string;
}



/**
 * Check if a given object implements the UserSvcListUsersRequest interface.
 */
export function instanceOfUserSvcListUsersRequest(value: object): value is UserSvcListUsersRequest {
    return true;
}

export function UserSvcListUsersRequestFromJSON(json: any): UserSvcListUsersRequest {
    return UserSvcListUsersRequestFromJSONTyped(json, false);
}

export function UserSvcListUsersRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcListUsersRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'afterTime': json['afterTime'] == null ? undefined : json['afterTime'],
        'contactId': json['contactId'] == null ? undefined : json['contactId'],
        'count': json['count'] == null ? undefined : json['count'],
        'ids': json['ids'] == null ? undefined : json['ids'],
        'limit': json['limit'] == null ? undefined : json['limit'],
        'order': json['order'] == null ? undefined : UserSvcOrderDirectionFromJSON(json['order']),
        'orderBy': json['orderBy'] == null ? undefined : UserSvcListUsersOrderByFromJSON(json['orderBy']),
        'search': json['search'] == null ? undefined : json['search'],
    };
}

export function UserSvcListUsersRequestToJSON(json: any): UserSvcListUsersRequest {
    return UserSvcListUsersRequestToJSONTyped(json, false);
}

export function UserSvcListUsersRequestToJSONTyped(value?: UserSvcListUsersRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'afterTime': value['afterTime'],
        'contactId': value['contactId'],
        'count': value['count'],
        'ids': value['ids'],
        'limit': value['limit'],
        'order': UserSvcOrderDirectionToJSON(value['order']),
        'orderBy': UserSvcListUsersOrderByToJSON(value['orderBy']),
        'search': value['search'],
    };
}

