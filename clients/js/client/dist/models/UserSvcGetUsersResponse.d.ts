/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { UserSvcUser } from './UserSvcUser';
/**
 *
 * @export
 * @interface UserSvcGetUsersResponse
 */
export interface UserSvcGetUsersResponse {
    /**
     *
     * @type {string}
     * @memberof UserSvcGetUsersResponse
     */
    after?: string;
    /**
     *
     * @type {number}
     * @memberof UserSvcGetUsersResponse
     */
    count?: number;
    /**
     *
     * @type {Array<UserSvcUser>}
     * @memberof UserSvcGetUsersResponse
     */
    users?: Array<UserSvcUser>;
}
/**
 * Check if a given object implements the UserSvcGetUsersResponse interface.
 */
export declare function instanceOfUserSvcGetUsersResponse(value: object): value is UserSvcGetUsersResponse;
export declare function UserSvcGetUsersResponseFromJSON(json: any): UserSvcGetUsersResponse;
export declare function UserSvcGetUsersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcGetUsersResponse;
export declare function UserSvcGetUsersResponseToJSON(json: any): UserSvcGetUsersResponse;
export declare function UserSvcGetUsersResponseToJSONTyped(value?: UserSvcGetUsersResponse | null, ignoreDiscriminator?: boolean): any;
