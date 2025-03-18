/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
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
 * @interface UserSvcSaveProfileRequest
 */
export interface UserSvcSaveProfileRequest {
    /**
     *
     * @type {string}
     * @memberof UserSvcSaveProfileRequest
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcSaveProfileRequest
     */
    slug?: string;
}
/**
 * Check if a given object implements the UserSvcSaveProfileRequest interface.
 */
export declare function instanceOfUserSvcSaveProfileRequest(value: object): value is UserSvcSaveProfileRequest;
export declare function UserSvcSaveProfileRequestFromJSON(json: any): UserSvcSaveProfileRequest;
export declare function UserSvcSaveProfileRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveProfileRequest;
export declare function UserSvcSaveProfileRequestToJSON(json: any): UserSvcSaveProfileRequest;
export declare function UserSvcSaveProfileRequestToJSONTyped(value?: UserSvcSaveProfileRequest | null, ignoreDiscriminator?: boolean): any;
