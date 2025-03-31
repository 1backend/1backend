/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcUserRecord interface.
 */
export function instanceOfUserSvcUserRecord(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('slug' in value) || value['slug'] === undefined)
        return false;
    return true;
}
export function UserSvcUserRecordFromJSON(json) {
    return UserSvcUserRecordFromJSONTyped(json, false);
}
export function UserSvcUserRecordFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'contactIds': json['contactIds'] == null ? undefined : json['contactIds'],
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'id': json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'roleIds': json['roleIds'] == null ? undefined : json['roleIds'],
        'slug': json['slug'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
    };
}
export function UserSvcUserRecordToJSON(json) {
    return UserSvcUserRecordToJSONTyped(json, false);
}
export function UserSvcUserRecordToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'contactIds': value['contactIds'],
        'createdAt': value['createdAt'],
        'id': value['id'],
        'name': value['name'],
        'roleIds': value['roleIds'],
        'slug': value['slug'],
        'updatedAt': value['updatedAt'],
    };
}
