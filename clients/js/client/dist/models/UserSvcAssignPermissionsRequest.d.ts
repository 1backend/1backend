/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { UserSvcPermissionLink } from './UserSvcPermissionLink';
/**
 *
 * @export
 * @interface UserSvcAssignPermissionsRequest
 */
export interface UserSvcAssignPermissionsRequest {
    /**
     *
     * @type {Array<UserSvcPermissionLink>}
     * @memberof UserSvcAssignPermissionsRequest
     */
    permissionLinks?: Array<UserSvcPermissionLink>;
}
/**
 * Check if a given object implements the UserSvcAssignPermissionsRequest interface.
 */
export declare function instanceOfUserSvcAssignPermissionsRequest(value: object): value is UserSvcAssignPermissionsRequest;
export declare function UserSvcAssignPermissionsRequestFromJSON(json: any): UserSvcAssignPermissionsRequest;
export declare function UserSvcAssignPermissionsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcAssignPermissionsRequest;
export declare function UserSvcAssignPermissionsRequestToJSON(json: any): UserSvcAssignPermissionsRequest;
export declare function UserSvcAssignPermissionsRequestToJSONTyped(value?: UserSvcAssignPermissionsRequest | null, ignoreDiscriminator?: boolean): any;
