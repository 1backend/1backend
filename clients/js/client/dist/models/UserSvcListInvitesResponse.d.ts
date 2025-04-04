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
import type { UserSvcInvite } from './UserSvcInvite';
/**
 *
 * @export
 * @interface UserSvcListInvitesResponse
 */
export interface UserSvcListInvitesResponse {
    /**
     *
     * @type {Array<UserSvcInvite>}
     * @memberof UserSvcListInvitesResponse
     */
    invites: Array<UserSvcInvite>;
}
/**
 * Check if a given object implements the UserSvcListInvitesResponse interface.
 */
export declare function instanceOfUserSvcListInvitesResponse(value: object): value is UserSvcListInvitesResponse;
export declare function UserSvcListInvitesResponseFromJSON(json: any): UserSvcListInvitesResponse;
export declare function UserSvcListInvitesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcListInvitesResponse;
export declare function UserSvcListInvitesResponseToJSON(json: any): UserSvcListInvitesResponse;
export declare function UserSvcListInvitesResponseToJSONTyped(value?: UserSvcListInvitesResponse | null, ignoreDiscriminator?: boolean): any;
