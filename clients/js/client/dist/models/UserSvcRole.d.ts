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
/**
 *
 * @export
 * @interface UserSvcRole
 */
export interface UserSvcRole {
    /**
     *
     * @type {string}
     * @memberof UserSvcRole
     */
    createdAt?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcRole
     */
    description?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcRole
     */
    id?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcRole
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcRole
     */
    ownerId?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcRole
     */
    updatedAt?: string;
}
/**
 * Check if a given object implements the UserSvcRole interface.
 */
export declare function instanceOfUserSvcRole(value: object): value is UserSvcRole;
export declare function UserSvcRoleFromJSON(json: any): UserSvcRole;
export declare function UserSvcRoleFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcRole;
export declare function UserSvcRoleToJSON(json: any): UserSvcRole;
export declare function UserSvcRoleToJSONTyped(value?: UserSvcRole | null, ignoreDiscriminator?: boolean): any;
