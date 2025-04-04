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
 * @interface UserSvcGetPublicKeyResponse
 */
export interface UserSvcGetPublicKeyResponse {
    /**
     *
     * @type {string}
     * @memberof UserSvcGetPublicKeyResponse
     */
    publicKey: string;
}
/**
 * Check if a given object implements the UserSvcGetPublicKeyResponse interface.
 */
export declare function instanceOfUserSvcGetPublicKeyResponse(value: object): value is UserSvcGetPublicKeyResponse;
export declare function UserSvcGetPublicKeyResponseFromJSON(json: any): UserSvcGetPublicKeyResponse;
export declare function UserSvcGetPublicKeyResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcGetPublicKeyResponse;
export declare function UserSvcGetPublicKeyResponseToJSON(json: any): UserSvcGetPublicKeyResponse;
export declare function UserSvcGetPublicKeyResponseToJSONTyped(value?: UserSvcGetPublicKeyResponse | null, ignoreDiscriminator?: boolean): any;
