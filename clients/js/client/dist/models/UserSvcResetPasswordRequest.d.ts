/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface UserSvcResetPasswordRequest
 */
export interface UserSvcResetPasswordRequest {
    /**
     *
     * @type {string}
     * @memberof UserSvcResetPasswordRequest
     */
    newPassword: string;
}
/**
 * Check if a given object implements the UserSvcResetPasswordRequest interface.
 */
export declare function instanceOfUserSvcResetPasswordRequest(value: object): value is UserSvcResetPasswordRequest;
export declare function UserSvcResetPasswordRequestFromJSON(json: any): UserSvcResetPasswordRequest;
export declare function UserSvcResetPasswordRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcResetPasswordRequest;
export declare function UserSvcResetPasswordRequestToJSON(json: any): UserSvcResetPasswordRequest;
export declare function UserSvcResetPasswordRequestToJSONTyped(value?: UserSvcResetPasswordRequest | null, ignoreDiscriminator?: boolean): any;
