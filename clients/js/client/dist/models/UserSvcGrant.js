/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcGrant interface.
 */
export function instanceOfUserSvcGrant(value) {
    if (!('permission' in value) || value['permission'] === undefined)
        return false;
    return true;
}
export function UserSvcGrantFromJSON(json) {
    return UserSvcGrantFromJSONTyped(json, false);
}
export function UserSvcGrantFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'id': json['id'] == null ? undefined : json['id'],
        'permission': json['permission'],
        'roles': json['roles'] == null ? undefined : json['roles'],
        'slugs': json['slugs'] == null ? undefined : json['slugs'],
    };
}
export function UserSvcGrantToJSON(json) {
    return UserSvcGrantToJSONTyped(json, false);
}
export function UserSvcGrantToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'id': value['id'],
        'permission': value['permission'],
        'roles': value['roles'],
        'slugs': value['slugs'],
    };
}
