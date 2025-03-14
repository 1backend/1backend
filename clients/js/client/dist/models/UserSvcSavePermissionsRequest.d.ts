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
import type { UserSvcPermission } from './UserSvcPermission';
/**
 *
 * @export
 * @interface UserSvcSavePermissionsRequest
 */
export interface UserSvcSavePermissionsRequest {
    /**
     *
     * @type {Array<UserSvcPermission>}
     * @memberof UserSvcSavePermissionsRequest
     */
    permissions?: Array<UserSvcPermission>;
}
/**
 * Check if a given object implements the UserSvcSavePermissionsRequest interface.
 */
export declare function instanceOfUserSvcSavePermissionsRequest(value: object): value is UserSvcSavePermissionsRequest;
export declare function UserSvcSavePermissionsRequestFromJSON(json: any): UserSvcSavePermissionsRequest;
export declare function UserSvcSavePermissionsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSavePermissionsRequest;
export declare function UserSvcSavePermissionsRequestToJSON(json: any): UserSvcSavePermissionsRequest;
export declare function UserSvcSavePermissionsRequestToJSONTyped(value?: UserSvcSavePermissionsRequest | null, ignoreDiscriminator?: boolean): any;
