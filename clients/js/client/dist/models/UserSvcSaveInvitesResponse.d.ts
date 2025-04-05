/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { UserSvcInvite } from './UserSvcInvite';
/**
 *
 * @export
 * @interface UserSvcSaveInvitesResponse
 */
export interface UserSvcSaveInvitesResponse {
    /**
     *
     * @type {Array<UserSvcInvite>}
     * @memberof UserSvcSaveInvitesResponse
     */
    invites: Array<UserSvcInvite>;
}
/**
 * Check if a given object implements the UserSvcSaveInvitesResponse interface.
 */
export declare function instanceOfUserSvcSaveInvitesResponse(value: object): value is UserSvcSaveInvitesResponse;
export declare function UserSvcSaveInvitesResponseFromJSON(json: any): UserSvcSaveInvitesResponse;
export declare function UserSvcSaveInvitesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveInvitesResponse;
export declare function UserSvcSaveInvitesResponseToJSON(json: any): UserSvcSaveInvitesResponse;
export declare function UserSvcSaveInvitesResponseToJSONTyped(value?: UserSvcSaveInvitesResponse | null, ignoreDiscriminator?: boolean): any;
