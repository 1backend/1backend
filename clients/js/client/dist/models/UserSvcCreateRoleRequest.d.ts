/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface UserSvcCreateRoleRequest
 */
export interface UserSvcCreateRoleRequest {
    /**
     *
     * @type {string}
     * @memberof UserSvcCreateRoleRequest
     */
    description?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcCreateRoleRequest
     */
    id: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcCreateRoleRequest
     */
    name?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof UserSvcCreateRoleRequest
     */
    permissionIds?: Array<string>;
}
/**
 * Check if a given object implements the UserSvcCreateRoleRequest interface.
 */
export declare function instanceOfUserSvcCreateRoleRequest(value: object): value is UserSvcCreateRoleRequest;
export declare function UserSvcCreateRoleRequestFromJSON(json: any): UserSvcCreateRoleRequest;
export declare function UserSvcCreateRoleRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcCreateRoleRequest;
export declare function UserSvcCreateRoleRequestToJSON(json: any): UserSvcCreateRoleRequest;
export declare function UserSvcCreateRoleRequestToJSONTyped(value?: UserSvcCreateRoleRequest | null, ignoreDiscriminator?: boolean): any;
