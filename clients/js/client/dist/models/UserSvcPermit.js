/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcPermit interface.
 */
export function instanceOfUserSvcPermit(value) {
    if (!('createdAt' in value) || value['createdAt'] === undefined)
        return false;
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('permission' in value) || value['permission'] === undefined)
        return false;
    if (!('updatedAt' in value) || value['updatedAt'] === undefined)
        return false;
    return true;
}
export function UserSvcPermitFromJSON(json) {
    return UserSvcPermitFromJSONTyped(json, false);
}
export function UserSvcPermitFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'createdAt': json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'id': json['id'],
        'permission': json['permission'],
        'roles': json['roles'] == null ? undefined : json['roles'],
        'slugs': json['slugs'] == null ? undefined : json['slugs'],
        'updatedAt': json['updatedAt'],
    };
}
export function UserSvcPermitToJSON(json) {
    return UserSvcPermitToJSONTyped(json, false);
}
export function UserSvcPermitToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'createdAt': value['createdAt'],
        'deletedAt': value['deletedAt'],
        'id': value['id'],
        'permission': value['permission'],
        'roles': value['roles'],
        'slugs': value['slugs'],
        'updatedAt': value['updatedAt'],
    };
}
