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
 * Check if a given object implements the UserSvcContactInput interface.
 */
export function instanceOfUserSvcContactInput(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('platform' in value) || value['platform'] === undefined)
        return false;
    return true;
}
export function UserSvcContactInputFromJSON(json) {
    return UserSvcContactInputFromJSONTyped(json, false);
}
export function UserSvcContactInputFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'handle': json['handle'] == null ? undefined : json['handle'],
        'id': json['id'],
        'platform': json['platform'],
    };
}
export function UserSvcContactInputToJSON(json) {
    return UserSvcContactInputToJSONTyped(json, false);
}
export function UserSvcContactInputToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'handle': value['handle'],
        'id': value['id'],
        'platform': value['platform'],
    };
}
