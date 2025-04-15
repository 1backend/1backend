/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { UserSvcOrganization } from './UserSvcOrganization';
import {
    UserSvcOrganizationFromJSON,
    UserSvcOrganizationFromJSONTyped,
    UserSvcOrganizationToJSON,
    UserSvcOrganizationToJSONTyped,
} from './UserSvcOrganization';

/**
 * 
 * @export
 * @interface UserSvcListOrganizationsResponse
 */
export interface UserSvcListOrganizationsResponse {
    /**
     * 
     * @type {Array<UserSvcOrganization>}
     * @memberof UserSvcListOrganizationsResponse
     */
    organizations?: Array<UserSvcOrganization>;
}

/**
 * Check if a given object implements the UserSvcListOrganizationsResponse interface.
 */
export function instanceOfUserSvcListOrganizationsResponse(value: object): value is UserSvcListOrganizationsResponse {
    return true;
}

export function UserSvcListOrganizationsResponseFromJSON(json: any): UserSvcListOrganizationsResponse {
    return UserSvcListOrganizationsResponseFromJSONTyped(json, false);
}

export function UserSvcListOrganizationsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcListOrganizationsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'organizations': json['organizations'] == null ? undefined : ((json['organizations'] as Array<any>).map(UserSvcOrganizationFromJSON)),
    };
}

export function UserSvcListOrganizationsResponseToJSON(json: any): UserSvcListOrganizationsResponse {
    return UserSvcListOrganizationsResponseToJSONTyped(json, false);
}

export function UserSvcListOrganizationsResponseToJSONTyped(value?: UserSvcListOrganizationsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'organizations': value['organizations'] == null ? undefined : ((value['organizations'] as Array<any>).map(UserSvcOrganizationToJSON)),
    };
}

