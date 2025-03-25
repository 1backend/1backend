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
import type { UserSvcAuthToken } from './UserSvcAuthToken';
/**
 *
 * @export
 * @interface UserSvcLoginResponse
 */
export interface UserSvcLoginResponse {
    /**
     *
     * @type {UserSvcAuthToken}
     * @memberof UserSvcLoginResponse
     */
    token?: UserSvcAuthToken;
}
/**
 * Check if a given object implements the UserSvcLoginResponse interface.
 */
export declare function instanceOfUserSvcLoginResponse(value: object): value is UserSvcLoginResponse;
export declare function UserSvcLoginResponseFromJSON(json: any): UserSvcLoginResponse;
export declare function UserSvcLoginResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcLoginResponse;
export declare function UserSvcLoginResponseToJSON(json: any): UserSvcLoginResponse;
export declare function UserSvcLoginResponseToJSONTyped(value?: UserSvcLoginResponse | null, ignoreDiscriminator?: boolean): any;
