/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { UserSvcRole } from './UserSvcRole';
/**
 *
 * @export
 * @interface UserSvcCreateRoleResponse
 */
export interface UserSvcCreateRoleResponse {
    /**
     *
     * @type {UserSvcRole}
     * @memberof UserSvcCreateRoleResponse
     */
    role?: UserSvcRole;
}
/**
 * Check if a given object implements the UserSvcCreateRoleResponse interface.
 */
export declare function instanceOfUserSvcCreateRoleResponse(value: object): value is UserSvcCreateRoleResponse;
export declare function UserSvcCreateRoleResponseFromJSON(json: any): UserSvcCreateRoleResponse;
export declare function UserSvcCreateRoleResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcCreateRoleResponse;
export declare function UserSvcCreateRoleResponseToJSON(json: any): UserSvcCreateRoleResponse;
export declare function UserSvcCreateRoleResponseToJSONTyped(value?: UserSvcCreateRoleResponse | null, ignoreDiscriminator?: boolean): any;
