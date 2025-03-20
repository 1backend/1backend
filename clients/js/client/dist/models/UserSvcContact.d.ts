/**
 * 1Backend
 * AI-native microservices platform.
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
 * @interface UserSvcContact
 */
export interface UserSvcContact {
    /**
     *
     * @type {string}
     * @memberof UserSvcContact
     */
    createdAt?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcContact
     */
    deletedAt?: string;
    /**
     * The unique identifier, which can be a URL.
     *
     * Example values: "joe12" (1backend username), "twitter.com/thejoe" (twitter url), "joe@joesdomain.com" (email)
     * @type {string}
     * @memberof UserSvcContact
     */
    id?: string;
    /**
     * If this is the primary contact method
     * @type {boolean}
     * @memberof UserSvcContact
     */
    isPrimary?: boolean;
    /**
     * Platform of the contact (e.g., "email", "phone", "twitter")
     * @type {string}
     * @memberof UserSvcContact
     */
    platform?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcContact
     */
    updatedAt?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcContact
     */
    userId?: string;
    /**
     * Value is the platform local unique identifier.
     * Ie. while the `id` of a Twitter contact is `twitter.com/thejoe`, the value will be only `thejoe`.
     * For email and phones the `id` and the `value` will be the same.
     * This field mostly exists for display purposes.
     *
     * Example values: "joe12" (1backend username), "thejoe" (twitter username), "joe@joesdomain.com" (email)
     * @type {string}
     * @memberof UserSvcContact
     */
    value?: string;
    /**
     * Whether the contact is verified
     * @type {boolean}
     * @memberof UserSvcContact
     */
    verified?: boolean;
}
/**
 * Check if a given object implements the UserSvcContact interface.
 */
export declare function instanceOfUserSvcContact(value: object): value is UserSvcContact;
export declare function UserSvcContactFromJSON(json: any): UserSvcContact;
export declare function UserSvcContactFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcContact;
export declare function UserSvcContactToJSON(json: any): UserSvcContact;
export declare function UserSvcContactToJSONTyped(value?: UserSvcContact | null, ignoreDiscriminator?: boolean): any;
