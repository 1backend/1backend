/* tslint:disable */
/* eslint-disable */
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
export const UserSvcListUsersOrderBy = {
    ListUsersOrderByCreatedAt: 'createdAt',
    ListUsersOrderByUpdatedAt: 'updatedAt'
} as const;
export type UserSvcListUsersOrderBy = typeof UserSvcListUsersOrderBy[keyof typeof UserSvcListUsersOrderBy];


export function instanceOfUserSvcListUsersOrderBy(value: any): boolean {
    for (const key in UserSvcListUsersOrderBy) {
        if (Object.prototype.hasOwnProperty.call(UserSvcListUsersOrderBy, key)) {
            if (UserSvcListUsersOrderBy[key as keyof typeof UserSvcListUsersOrderBy] === value) {
                return true;
            }
        }
    }
    return false;
}

export function UserSvcListUsersOrderByFromJSON(json: any): UserSvcListUsersOrderBy {
    return UserSvcListUsersOrderByFromJSONTyped(json, false);
}

export function UserSvcListUsersOrderByFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcListUsersOrderBy {
    return json as UserSvcListUsersOrderBy;
}

export function UserSvcListUsersOrderByToJSON(value?: UserSvcListUsersOrderBy | null): any {
    return value as any;
}

export function UserSvcListUsersOrderByToJSONTyped(value: any, ignoreDiscriminator: boolean): UserSvcListUsersOrderBy {
    return value as UserSvcListUsersOrderBy;
}

