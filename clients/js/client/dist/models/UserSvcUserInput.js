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
 * Check if a given object implements the UserSvcUserInput interface.
 */
export function instanceOfUserSvcUserInput(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('slug' in value) || value['slug'] === undefined)
        return false;
    return true;
}
export function UserSvcUserInputFromJSON(json) {
    return UserSvcUserInputFromJSONTyped(json, false);
}
export function UserSvcUserInputFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'id': json['id'],
        'labels': json['labels'] == null ? undefined : json['labels'],
        'name': json['name'] == null ? undefined : json['name'],
        'slug': json['slug'],
        'thumbnailFileId': json['thumbnailFileId'] == null ? undefined : json['thumbnailFileId'],
    };
}
export function UserSvcUserInputToJSON(json) {
    return UserSvcUserInputToJSONTyped(json, false);
}
export function UserSvcUserInputToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'id': value['id'],
        'labels': value['labels'],
        'name': value['name'],
        'slug': value['slug'],
        'thumbnailFileId': value['thumbnailFileId'],
    };
}
