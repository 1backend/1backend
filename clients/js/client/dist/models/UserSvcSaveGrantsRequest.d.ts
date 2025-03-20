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
import type { UserSvcGrant } from './UserSvcGrant';
/**
 *
 * @export
 * @interface UserSvcSaveGrantsRequest
 */
export interface UserSvcSaveGrantsRequest {
    /**
     *
     * @type {Array<UserSvcGrant>}
     * @memberof UserSvcSaveGrantsRequest
     */
    grants?: Array<UserSvcGrant>;
}
/**
 * Check if a given object implements the UserSvcSaveGrantsRequest interface.
 */
export declare function instanceOfUserSvcSaveGrantsRequest(value: object): value is UserSvcSaveGrantsRequest;
export declare function UserSvcSaveGrantsRequestFromJSON(json: any): UserSvcSaveGrantsRequest;
export declare function UserSvcSaveGrantsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveGrantsRequest;
export declare function UserSvcSaveGrantsRequestToJSON(json: any): UserSvcSaveGrantsRequest;
export declare function UserSvcSaveGrantsRequestToJSONTyped(value?: UserSvcSaveGrantsRequest | null, ignoreDiscriminator?: boolean): any;
