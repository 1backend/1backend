/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface UserSvcIsAuthorizedRequest
 */
export interface UserSvcIsAuthorizedRequest {
    /**
     *
     * @type {Array<string>}
     * @memberof UserSvcIsAuthorizedRequest
     */
    contactsGranted?: Array<string>;
    /**
     *
     * @type {Array<string>}
     * @memberof UserSvcIsAuthorizedRequest
     */
    grantedSlugs?: Array<string>;
}
/**
 * Check if a given object implements the UserSvcIsAuthorizedRequest interface.
 */
export declare function instanceOfUserSvcIsAuthorizedRequest(value: object): value is UserSvcIsAuthorizedRequest;
export declare function UserSvcIsAuthorizedRequestFromJSON(json: any): UserSvcIsAuthorizedRequest;
export declare function UserSvcIsAuthorizedRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcIsAuthorizedRequest;
export declare function UserSvcIsAuthorizedRequestToJSON(json: any): UserSvcIsAuthorizedRequest;
export declare function UserSvcIsAuthorizedRequestToJSONTyped(value?: UserSvcIsAuthorizedRequest | null, ignoreDiscriminator?: boolean): any;
