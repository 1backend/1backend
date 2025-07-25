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
 * Check if a given object implements the UserSvcSaveSelfRequest interface.
 */
export function instanceOfUserSvcSaveSelfRequest(value) {
    return true;
}
export function UserSvcSaveSelfRequestFromJSON(json) {
    return UserSvcSaveSelfRequestFromJSONTyped(json, false);
}
export function UserSvcSaveSelfRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'labels': json['labels'] == null ? undefined : json['labels'],
        'name': json['name'] == null ? undefined : json['name'],
        'thumbnailFileId': json['thumbnailFileId'] == null ? undefined : json['thumbnailFileId'],
    };
}
export function UserSvcSaveSelfRequestToJSON(json) {
    return UserSvcSaveSelfRequestToJSONTyped(json, false);
}
export function UserSvcSaveSelfRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'labels': value['labels'],
        'name': value['name'],
        'thumbnailFileId': value['thumbnailFileId'],
    };
}
