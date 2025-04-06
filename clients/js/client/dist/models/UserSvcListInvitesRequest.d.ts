/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface UserSvcListInvitesRequest
 */
export interface UserSvcListInvitesRequest {
    /**
     *
     * @type {string}
     * @memberof UserSvcListInvitesRequest
     */
    contactId?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcListInvitesRequest
     */
    role?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcListInvitesRequest
     */
    userId?: string;
}
/**
 * Check if a given object implements the UserSvcListInvitesRequest interface.
 */
export declare function instanceOfUserSvcListInvitesRequest(value: object): value is UserSvcListInvitesRequest;
export declare function UserSvcListInvitesRequestFromJSON(json: any): UserSvcListInvitesRequest;
export declare function UserSvcListInvitesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcListInvitesRequest;
export declare function UserSvcListInvitesRequestToJSON(json: any): UserSvcListInvitesRequest;
export declare function UserSvcListInvitesRequestToJSONTyped(value?: UserSvcListInvitesRequest | null, ignoreDiscriminator?: boolean): any;
