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
import type { UserSvcGrant } from './UserSvcGrant';
/**
 *
 * @export
 * @interface UserSvcListGrantsResponse
 */
export interface UserSvcListGrantsResponse {
    /**
     *
     * @type {Array<UserSvcGrant>}
     * @memberof UserSvcListGrantsResponse
     */
    grants?: Array<UserSvcGrant>;
}
/**
 * Check if a given object implements the UserSvcListGrantsResponse interface.
 */
export declare function instanceOfUserSvcListGrantsResponse(value: object): value is UserSvcListGrantsResponse;
export declare function UserSvcListGrantsResponseFromJSON(json: any): UserSvcListGrantsResponse;
export declare function UserSvcListGrantsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcListGrantsResponse;
export declare function UserSvcListGrantsResponseToJSON(json: any): UserSvcListGrantsResponse;
export declare function UserSvcListGrantsResponseToJSONTyped(value?: UserSvcListGrantsResponse | null, ignoreDiscriminator?: boolean): any;
