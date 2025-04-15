/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcSaveOrganizationRequest interface.
 */
export function instanceOfUserSvcSaveOrganizationRequest(value) {
    if (!('slug' in value) || value['slug'] === undefined)
        return false;
    return true;
}
export function UserSvcSaveOrganizationRequestFromJSON(json) {
    return UserSvcSaveOrganizationRequestFromJSONTyped(json, false);
}
export function UserSvcSaveOrganizationRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'slug': json['slug'],
        'thumbnailFileId': json['thumbnailFileId'] == null ? undefined : json['thumbnailFileId'],
    };
}
export function UserSvcSaveOrganizationRequestToJSON(json) {
    return UserSvcSaveOrganizationRequestToJSONTyped(json, false);
}
export function UserSvcSaveOrganizationRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'id': value['id'],
        'name': value['name'],
        'slug': value['slug'],
        'thumbnailFileId': value['thumbnailFileId'],
    };
}
