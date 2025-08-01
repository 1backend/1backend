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
import type { UserSvcEnrollInput } from './UserSvcEnrollInput';
/**
 *
 * @export
 * @interface UserSvcSaveEnrollsRequest
 */
export interface UserSvcSaveEnrollsRequest {
    /**
     *
     * @type {Array<UserSvcEnrollInput>}
     * @memberof UserSvcSaveEnrollsRequest
     */
    enrolls: Array<UserSvcEnrollInput>;
}
/**
 * Check if a given object implements the UserSvcSaveEnrollsRequest interface.
 */
export declare function instanceOfUserSvcSaveEnrollsRequest(value: object): value is UserSvcSaveEnrollsRequest;
export declare function UserSvcSaveEnrollsRequestFromJSON(json: any): UserSvcSaveEnrollsRequest;
export declare function UserSvcSaveEnrollsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveEnrollsRequest;
export declare function UserSvcSaveEnrollsRequestToJSON(json: any): UserSvcSaveEnrollsRequest;
export declare function UserSvcSaveEnrollsRequestToJSONTyped(value?: UserSvcSaveEnrollsRequest | null, ignoreDiscriminator?: boolean): any;
