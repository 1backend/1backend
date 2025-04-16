/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface UserSvcSaveSelfRequest
 */
export interface UserSvcSaveSelfRequest {
    /**
     *
     * @type {{ [key: string]: string; }}
     * @memberof UserSvcSaveSelfRequest
     */
    labels?: {
        [key: string]: string;
    };
    /**
     *
     * @type {string}
     * @memberof UserSvcSaveSelfRequest
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcSaveSelfRequest
     */
    thumbnailFileId?: string;
}
/**
 * Check if a given object implements the UserSvcSaveSelfRequest interface.
 */
export declare function instanceOfUserSvcSaveSelfRequest(value: object): value is UserSvcSaveSelfRequest;
export declare function UserSvcSaveSelfRequestFromJSON(json: any): UserSvcSaveSelfRequest;
export declare function UserSvcSaveSelfRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveSelfRequest;
export declare function UserSvcSaveSelfRequestToJSON(json: any): UserSvcSaveSelfRequest;
export declare function UserSvcSaveSelfRequestToJSONTyped(value?: UserSvcSaveSelfRequest | null, ignoreDiscriminator?: boolean): any;
