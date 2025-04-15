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
 * @interface UserSvcListEnrollsRequest
 */
export interface UserSvcListEnrollsRequest {
    /**
     *
     * @type {string}
     * @memberof UserSvcListEnrollsRequest
     */
    contactId?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcListEnrollsRequest
     */
    role?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcListEnrollsRequest
     */
    userId?: string;
}
/**
 * Check if a given object implements the UserSvcListEnrollsRequest interface.
 */
export declare function instanceOfUserSvcListEnrollsRequest(value: object): value is UserSvcListEnrollsRequest;
export declare function UserSvcListEnrollsRequestFromJSON(json: any): UserSvcListEnrollsRequest;
export declare function UserSvcListEnrollsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcListEnrollsRequest;
export declare function UserSvcListEnrollsRequestToJSON(json: any): UserSvcListEnrollsRequest;
export declare function UserSvcListEnrollsRequestToJSONTyped(value?: UserSvcListEnrollsRequest | null, ignoreDiscriminator?: boolean): any;
