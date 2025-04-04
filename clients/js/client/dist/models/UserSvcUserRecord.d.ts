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
 * @interface UserSvcUserRecord
 */
export interface UserSvcUserRecord {
    /**
     *
     * @type {Array<string>}
     * @memberof UserSvcUserRecord
     */
    contactIds?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof UserSvcUserRecord
     */
    createdAt?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcUserRecord
     */
    id: string;
    /**
     * Full name of the user.
     * @type {string}
     * @memberof UserSvcUserRecord
     */
    name?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof UserSvcUserRecord
     */
    roleIds?: Array<string>;
    /**
     * URL-friendly unique (inside the 1Backend platform) identifier for the `user`.
     * @type {string}
     * @memberof UserSvcUserRecord
     */
    slug: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcUserRecord
     */
    updatedAt?: string;
}
/**
 * Check if a given object implements the UserSvcUserRecord interface.
 */
export declare function instanceOfUserSvcUserRecord(value: object): value is UserSvcUserRecord;
export declare function UserSvcUserRecordFromJSON(json: any): UserSvcUserRecord;
export declare function UserSvcUserRecordFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcUserRecord;
export declare function UserSvcUserRecordToJSON(json: any): UserSvcUserRecord;
export declare function UserSvcUserRecordToJSONTyped(value?: UserSvcUserRecord | null, ignoreDiscriminator?: boolean): any;
