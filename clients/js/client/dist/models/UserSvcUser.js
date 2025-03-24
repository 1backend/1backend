/* tslint:disable */
/* eslint-disable */
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
import { UserSvcContactFromJSON, UserSvcContactToJSON, } from './UserSvcContact';
/**
 * Check if a given object implements the UserSvcUser interface.
 */
export function instanceOfUserSvcUser(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('slug' in value) || value['slug'] === undefined)
        return false;
    return true;
}
export function UserSvcUserFromJSON(json) {
    return UserSvcUserFromJSONTyped(json, false);
}
export function UserSvcUserFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'contacts': json['contacts'] == null ? undefined : (json['contacts'].map(UserSvcContactFromJSON)),
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'id': json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'passwordHash': json['passwordHash'] == null ? undefined : json['passwordHash'],
        'slug': json['slug'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
    };
}
export function UserSvcUserToJSON(json) {
    return UserSvcUserToJSONTyped(json, false);
}
export function UserSvcUserToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'contacts': value['contacts'] == null ? undefined : (value['contacts'].map(UserSvcContactToJSON)),
        'createdAt': value['createdAt'],
        'deletedAt': value['deletedAt'],
        'id': value['id'],
        'name': value['name'],
        'passwordHash': value['passwordHash'],
        'slug': value['slug'],
        'updatedAt': value['updatedAt'],
    };
}
