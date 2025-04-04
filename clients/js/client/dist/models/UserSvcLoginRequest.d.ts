/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface UserSvcLoginRequest
 */
export interface UserSvcLoginRequest {
    /**
     *
     * @type {string}
     * @memberof UserSvcLoginRequest
     */
    contact?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcLoginRequest
     */
    password?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcLoginRequest
     */
    slug?: string;
}
/**
 * Check if a given object implements the UserSvcLoginRequest interface.
 */
export declare function instanceOfUserSvcLoginRequest(value: object): value is UserSvcLoginRequest;
export declare function UserSvcLoginRequestFromJSON(json: any): UserSvcLoginRequest;
export declare function UserSvcLoginRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcLoginRequest;
export declare function UserSvcLoginRequestToJSON(json: any): UserSvcLoginRequest;
export declare function UserSvcLoginRequestToJSONTyped(value?: UserSvcLoginRequest | null, ignoreDiscriminator?: boolean): any;
