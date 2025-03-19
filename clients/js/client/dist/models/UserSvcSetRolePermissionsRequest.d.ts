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
/**
 *
 * @export
 * @interface UserSvcSetRolePermissionsRequest
 */
export interface UserSvcSetRolePermissionsRequest {
    /**
     *
     * @type {Array<string>}
     * @memberof UserSvcSetRolePermissionsRequest
     */
    permissionIds?: Array<string>;
}
/**
 * Check if a given object implements the UserSvcSetRolePermissionsRequest interface.
 */
export declare function instanceOfUserSvcSetRolePermissionsRequest(value: object): value is UserSvcSetRolePermissionsRequest;
export declare function UserSvcSetRolePermissionsRequestFromJSON(json: any): UserSvcSetRolePermissionsRequest;
export declare function UserSvcSetRolePermissionsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSetRolePermissionsRequest;
export declare function UserSvcSetRolePermissionsRequestToJSON(json: any): UserSvcSetRolePermissionsRequest;
export declare function UserSvcSetRolePermissionsRequestToJSONTyped(value?: UserSvcSetRolePermissionsRequest | null, ignoreDiscriminator?: boolean): any;
