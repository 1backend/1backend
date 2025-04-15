/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface UserSvcEnrollInput
 */
export interface UserSvcEnrollInput {
    /**
     * ContactId is the the recipient of the enroll.
     * If the user is already registered, the role is assigned immediately;
     * otherwise, it is applied upon registration.
     * @type {string}
     * @memberof UserSvcEnrollInput
     */
    contactId?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcEnrollInput
     */
    id?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcEnrollInput
     */
    role: string;
    /**
     * UserId is the recipient of the enroll.
     * If the user is already registered, the role is assigned immediately;
     * otherwise, it is applied upon registration.
     * @type {string}
     * @memberof UserSvcEnrollInput
     */
    userId?: string;
}
/**
 * Check if a given object implements the UserSvcEnrollInput interface.
 */
export declare function instanceOfUserSvcEnrollInput(value: object): value is UserSvcEnrollInput;
export declare function UserSvcEnrollInputFromJSON(json: any): UserSvcEnrollInput;
export declare function UserSvcEnrollInputFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcEnrollInput;
export declare function UserSvcEnrollInputToJSON(json: any): UserSvcEnrollInput;
export declare function UserSvcEnrollInputToJSONTyped(value?: UserSvcEnrollInput | null, ignoreDiscriminator?: boolean): any;
