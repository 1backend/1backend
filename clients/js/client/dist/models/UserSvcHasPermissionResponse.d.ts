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
import type { UserSvcUser } from './UserSvcUser';
/**
 *
 * @export
 * @interface UserSvcHasPermissionResponse
 */
export interface UserSvcHasPermissionResponse {
    /**
     *
     * @type {string}
     * @memberof UserSvcHasPermissionResponse
     */
    app?: string;
    /**
     *
     * @type {boolean}
     * @memberof UserSvcHasPermissionResponse
     */
    authorized: boolean;
    /**
     *
     * @type {string}
     * @memberof UserSvcHasPermissionResponse
     */
    until: string;
    /**
     *
     * @type {UserSvcUser}
     * @memberof UserSvcHasPermissionResponse
     */
    user: UserSvcUser;
}
/**
 * Check if a given object implements the UserSvcHasPermissionResponse interface.
 */
export declare function instanceOfUserSvcHasPermissionResponse(value: object): value is UserSvcHasPermissionResponse;
export declare function UserSvcHasPermissionResponseFromJSON(json: any): UserSvcHasPermissionResponse;
export declare function UserSvcHasPermissionResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcHasPermissionResponse;
export declare function UserSvcHasPermissionResponseToJSON(json: any): UserSvcHasPermissionResponse;
export declare function UserSvcHasPermissionResponseToJSONTyped(value?: UserSvcHasPermissionResponse | null, ignoreDiscriminator?: boolean): any;
