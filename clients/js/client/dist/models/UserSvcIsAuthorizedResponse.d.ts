/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
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
 * @interface UserSvcIsAuthorizedResponse
 */
export interface UserSvcIsAuthorizedResponse {
    /**
     *
     * @type {boolean}
     * @memberof UserSvcIsAuthorizedResponse
     */
    authorized?: boolean;
    /**
     *
     * @type {UserSvcUser}
     * @memberof UserSvcIsAuthorizedResponse
     */
    user?: UserSvcUser;
}
/**
 * Check if a given object implements the UserSvcIsAuthorizedResponse interface.
 */
export declare function instanceOfUserSvcIsAuthorizedResponse(value: object): value is UserSvcIsAuthorizedResponse;
export declare function UserSvcIsAuthorizedResponseFromJSON(json: any): UserSvcIsAuthorizedResponse;
export declare function UserSvcIsAuthorizedResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcIsAuthorizedResponse;
export declare function UserSvcIsAuthorizedResponseToJSON(json: any): UserSvcIsAuthorizedResponse;
export declare function UserSvcIsAuthorizedResponseToJSONTyped(value?: UserSvcIsAuthorizedResponse | null, ignoreDiscriminator?: boolean): any;
