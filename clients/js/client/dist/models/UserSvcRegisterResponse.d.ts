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
import type { UserSvcAuthToken } from './UserSvcAuthToken';
/**
 *
 * @export
 * @interface UserSvcRegisterResponse
 */
export interface UserSvcRegisterResponse {
    /**
     *
     * @type {UserSvcAuthToken}
     * @memberof UserSvcRegisterResponse
     */
    token?: UserSvcAuthToken;
}
/**
 * Check if a given object implements the UserSvcRegisterResponse interface.
 */
export declare function instanceOfUserSvcRegisterResponse(value: object): value is UserSvcRegisterResponse;
export declare function UserSvcRegisterResponseFromJSON(json: any): UserSvcRegisterResponse;
export declare function UserSvcRegisterResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcRegisterResponse;
export declare function UserSvcRegisterResponseToJSON(json: any): UserSvcRegisterResponse;
export declare function UserSvcRegisterResponseToJSONTyped(value?: UserSvcRegisterResponse | null, ignoreDiscriminator?: boolean): any;
