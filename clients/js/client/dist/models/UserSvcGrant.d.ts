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
/**
 *
 * @export
 * @interface UserSvcGrant
 */
export interface UserSvcGrant {
    /**
     *
     * @type {string}
     * @memberof UserSvcGrant
     */
    id?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcGrant
     */
    permissionId?: string;
    /**
     * Slugs who are granted the PermissionId
     * @type {Array<string>}
     * @memberof UserSvcGrant
     */
    slugs?: Array<string>;
}
/**
 * Check if a given object implements the UserSvcGrant interface.
 */
export declare function instanceOfUserSvcGrant(value: object): value is UserSvcGrant;
export declare function UserSvcGrantFromJSON(json: any): UserSvcGrant;
export declare function UserSvcGrantFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcGrant;
export declare function UserSvcGrantToJSON(json: any): UserSvcGrant;
export declare function UserSvcGrantToJSONTyped(value?: UserSvcGrant | null, ignoreDiscriminator?: boolean): any;
