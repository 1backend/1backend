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
 * @interface UserSvcUser
 */
export interface UserSvcUser {
    /**
     *
     * @type {string}
     * @memberof UserSvcUser
     */
    createdAt?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcUser
     */
    deletedAt?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcUser
     */
    id: string;
    /**
     * Full name of the user.
     * @type {string}
     * @memberof UserSvcUser
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcUser
     */
    passwordHash?: string;
    /**
     * URL-friendly unique (inside the 1Backend platform) identifier for the `user`.
     * @type {string}
     * @memberof UserSvcUser
     */
    slug: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcUser
     */
    thumbnailFileId?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcUser
     */
    updatedAt?: string;
}
/**
 * Check if a given object implements the UserSvcUser interface.
 */
export declare function instanceOfUserSvcUser(value: object): value is UserSvcUser;
export declare function UserSvcUserFromJSON(json: any): UserSvcUser;
export declare function UserSvcUserFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcUser;
export declare function UserSvcUserToJSON(json: any): UserSvcUser;
export declare function UserSvcUserToJSONTyped(value?: UserSvcUser | null, ignoreDiscriminator?: boolean): any;
