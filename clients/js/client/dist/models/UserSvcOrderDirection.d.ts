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
 */
export declare const UserSvcOrderDirection: {
    readonly OrderDirectionAsc: "asc";
    readonly OrderDirectionDesc: "desc";
};
export type UserSvcOrderDirection = typeof UserSvcOrderDirection[keyof typeof UserSvcOrderDirection];
export declare function instanceOfUserSvcOrderDirection(value: any): boolean;
export declare function UserSvcOrderDirectionFromJSON(json: any): UserSvcOrderDirection;
export declare function UserSvcOrderDirectionFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcOrderDirection;
export declare function UserSvcOrderDirectionToJSON(value?: UserSvcOrderDirection | null): any;
export declare function UserSvcOrderDirectionToJSONTyped(value: any, ignoreDiscriminator: boolean): UserSvcOrderDirection;
