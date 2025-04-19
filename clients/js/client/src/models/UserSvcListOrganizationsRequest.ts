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
/**
 * 
 * @export
 * @interface UserSvcListOrganizationsRequest
 */
export interface UserSvcListOrganizationsRequest {
    /**
     * Organizations by default come back ordered
     * desc by `createdAt` field.
     * @type {string}
     * @memberof UserSvcListOrganizationsRequest
     */
    afterTime?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof UserSvcListOrganizationsRequest
     */
    ids?: Array<string>;
    /**
     * 
     * @type {number}
     * @memberof UserSvcListOrganizationsRequest
     */
    limit?: number;
}

/**
 * Check if a given object implements the UserSvcListOrganizationsRequest interface.
 */
export function instanceOfUserSvcListOrganizationsRequest(value: object): value is UserSvcListOrganizationsRequest {
    return true;
}

export function UserSvcListOrganizationsRequestFromJSON(json: any): UserSvcListOrganizationsRequest {
    return UserSvcListOrganizationsRequestFromJSONTyped(json, false);
}

export function UserSvcListOrganizationsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcListOrganizationsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'afterTime': json['afterTime'] == null ? undefined : json['afterTime'],
        'ids': json['ids'] == null ? undefined : json['ids'],
        'limit': json['limit'] == null ? undefined : json['limit'],
    };
}

export function UserSvcListOrganizationsRequestToJSON(json: any): UserSvcListOrganizationsRequest {
    return UserSvcListOrganizationsRequestToJSONTyped(json, false);
}

export function UserSvcListOrganizationsRequestToJSONTyped(value?: UserSvcListOrganizationsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'afterTime': value['afterTime'],
        'ids': value['ids'],
        'limit': value['limit'],
    };
}

