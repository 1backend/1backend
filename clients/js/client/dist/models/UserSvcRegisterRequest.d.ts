/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { UserSvcContact } from './UserSvcContact';
/**
 *
 * @export
 * @interface UserSvcRegisterRequest
 */
export interface UserSvcRegisterRequest {
    /**
     *
     * @type {UserSvcContact}
     * @memberof UserSvcRegisterRequest
     */
    contact?: UserSvcContact;
    /**
     *
     * @type {string}
     * @memberof UserSvcRegisterRequest
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcRegisterRequest
     */
    password?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcRegisterRequest
     */
    slug?: string;
}
/**
 * Check if a given object implements the UserSvcRegisterRequest interface.
 */
export declare function instanceOfUserSvcRegisterRequest(value: object): value is UserSvcRegisterRequest;
export declare function UserSvcRegisterRequestFromJSON(json: any): UserSvcRegisterRequest;
export declare function UserSvcRegisterRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcRegisterRequest;
export declare function UserSvcRegisterRequestToJSON(json: any): UserSvcRegisterRequest;
export declare function UserSvcRegisterRequestToJSONTyped(value?: UserSvcRegisterRequest | null, ignoreDiscriminator?: boolean): any;
