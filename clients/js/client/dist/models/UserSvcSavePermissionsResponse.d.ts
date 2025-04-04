/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
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
 * @interface UserSvcSavePermissionsResponse
 */
export interface UserSvcSavePermissionsResponse {
    /**
     *
     * @type {Array<UserSvcPermission>}
     * @memberof UserSvcSavePermissionsResponse
     */
    permissions?: Array<UserSvcPermission>;
}
/**
 * Check if a given object implements the UserSvcSavePermissionsResponse interface.
 */
export declare function instanceOfUserSvcSavePermissionsResponse(value: object): value is UserSvcSavePermissionsResponse;
export declare function UserSvcSavePermissionsResponseFromJSON(json: any): UserSvcSavePermissionsResponse;
export declare function UserSvcSavePermissionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSavePermissionsResponse;
export declare function UserSvcSavePermissionsResponseToJSON(json: any): UserSvcSavePermissionsResponse;
export declare function UserSvcSavePermissionsResponseToJSONTyped(value?: UserSvcSavePermissionsResponse | null, ignoreDiscriminator?: boolean): any;
