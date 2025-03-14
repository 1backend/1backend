/**
 * OpenOrch
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
 * @interface UserSvcCreateUserRequest
 */
export interface UserSvcCreateUserRequest {
    /**
     *
     * @type {string}
     * @memberof UserSvcCreateUserRequest
     */
    password?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof UserSvcCreateUserRequest
     */
    roleIds?: Array<string>;
    /**
     *
     * @type {UserSvcUser}
     * @memberof UserSvcCreateUserRequest
     */
    user?: UserSvcUser;
}
/**
 * Check if a given object implements the UserSvcCreateUserRequest interface.
 */
export declare function instanceOfUserSvcCreateUserRequest(value: object): value is UserSvcCreateUserRequest;
export declare function UserSvcCreateUserRequestFromJSON(json: any): UserSvcCreateUserRequest;
export declare function UserSvcCreateUserRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcCreateUserRequest;
export declare function UserSvcCreateUserRequestToJSON(json: any): UserSvcCreateUserRequest;
export declare function UserSvcCreateUserRequestToJSONTyped(value?: UserSvcCreateUserRequest | null, ignoreDiscriminator?: boolean): any;
