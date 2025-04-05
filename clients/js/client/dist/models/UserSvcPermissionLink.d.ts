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
/**
 *
 * @export
 * @interface UserSvcPermissionLink
 */
export interface UserSvcPermissionLink {
    /**
     *
     * @type {string}
     * @memberof UserSvcPermissionLink
     */
    permission: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcPermissionLink
     */
    role: string;
}
/**
 * Check if a given object implements the UserSvcPermissionLink interface.
 */
export declare function instanceOfUserSvcPermissionLink(value: object): value is UserSvcPermissionLink;
export declare function UserSvcPermissionLinkFromJSON(json: any): UserSvcPermissionLink;
export declare function UserSvcPermissionLinkFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcPermissionLink;
export declare function UserSvcPermissionLinkToJSON(json: any): UserSvcPermissionLink;
export declare function UserSvcPermissionLinkToJSONTyped(value?: UserSvcPermissionLink | null, ignoreDiscriminator?: boolean): any;
