/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface UserSvcErrorResponse
 */
export interface UserSvcErrorResponse {
    /**
     *
     * @type {string}
     * @memberof UserSvcErrorResponse
     */
    error?: string;
}
/**
 * Check if a given object implements the UserSvcErrorResponse interface.
 */
export declare function instanceOfUserSvcErrorResponse(value: object): value is UserSvcErrorResponse;
export declare function UserSvcErrorResponseFromJSON(json: any): UserSvcErrorResponse;
export declare function UserSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcErrorResponse;
export declare function UserSvcErrorResponseToJSON(json: any): UserSvcErrorResponse;
export declare function UserSvcErrorResponseToJSONTyped(value?: UserSvcErrorResponse | null, ignoreDiscriminator?: boolean): any;
