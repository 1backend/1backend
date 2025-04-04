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
 * @interface UserSvcAuthToken
 */
export interface UserSvcAuthToken {
    /**
     * Active tokens contain the most up-to-date information.
     * When a user's role changes—due to role assignment, organization
     * creation/assignment, etc.—all existing tokens are marked inactive.
     * Active tokens are reused during login, while inactive tokens
     * are retained for historical reference.
     * @type {boolean}
     * @memberof UserSvcAuthToken
     */
    active?: boolean;
    /**
     *
     * @type {string}
     * @memberof UserSvcAuthToken
     */
    createdAt?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcAuthToken
     */
    deletedAt?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcAuthToken
     */
    id?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcAuthToken
     */
    token: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcAuthToken
     */
    updatedAt?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcAuthToken
     */
    userId: string;
}
/**
 * Check if a given object implements the UserSvcAuthToken interface.
 */
export declare function instanceOfUserSvcAuthToken(value: object): value is UserSvcAuthToken;
export declare function UserSvcAuthTokenFromJSON(json: any): UserSvcAuthToken;
export declare function UserSvcAuthTokenFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcAuthToken;
export declare function UserSvcAuthTokenToJSON(json: any): UserSvcAuthToken;
export declare function UserSvcAuthTokenToJSONTyped(value?: UserSvcAuthToken | null, ignoreDiscriminator?: boolean): any;
