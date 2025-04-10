/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { UserSvcOrganization } from './UserSvcOrganization';
import type { UserSvcUser } from './UserSvcUser';
/**
 *
 * @export
 * @interface UserSvcReadUserByTokenResponse
 */
export interface UserSvcReadUserByTokenResponse {
    /**
     * Active organization of the caller user, if it has any.
     * @type {string}
     * @memberof UserSvcReadUserByTokenResponse
     */
    activeOrganizationId?: string;
    /**
     * Organizations of the caller user.
     * @type {Array<UserSvcOrganization>}
     * @memberof UserSvcReadUserByTokenResponse
     */
    organizations?: Array<UserSvcOrganization>;
    /**
     * Roles the token has that made this request.
     * @type {Array<string>}
     * @memberof UserSvcReadUserByTokenResponse
     */
    roles?: Array<string>;
    /**
     * The user who made the request.
     * @type {UserSvcUser}
     * @memberof UserSvcReadUserByTokenResponse
     */
    user: UserSvcUser;
}
/**
 * Check if a given object implements the UserSvcReadUserByTokenResponse interface.
 */
export declare function instanceOfUserSvcReadUserByTokenResponse(value: object): value is UserSvcReadUserByTokenResponse;
export declare function UserSvcReadUserByTokenResponseFromJSON(json: any): UserSvcReadUserByTokenResponse;
export declare function UserSvcReadUserByTokenResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcReadUserByTokenResponse;
export declare function UserSvcReadUserByTokenResponseToJSON(json: any): UserSvcReadUserByTokenResponse;
export declare function UserSvcReadUserByTokenResponseToJSONTyped(value?: UserSvcReadUserByTokenResponse | null, ignoreDiscriminator?: boolean): any;
