/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcOrganization interface.
 */
export function instanceOfUserSvcOrganization(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('name' in value) || value['name'] === undefined)
        return false;
    if (!('slug' in value) || value['slug'] === undefined)
        return false;
    return true;
}
export function UserSvcOrganizationFromJSON(json) {
    return UserSvcOrganizationFromJSONTyped(json, false);
}
export function UserSvcOrganizationFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'id': json['id'],
        'name': json['name'],
        'slug': json['slug'],
        'thumbnailFileId': json['thumbnailFileId'] == null ? undefined : json['thumbnailFileId'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
    };
}
export function UserSvcOrganizationToJSON(json) {
    return UserSvcOrganizationToJSONTyped(json, false);
}
export function UserSvcOrganizationToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'createdAt': value['createdAt'],
        'deletedAt': value['deletedAt'],
        'id': value['id'],
        'name': value['name'],
        'slug': value['slug'],
        'thumbnailFileId': value['thumbnailFileId'],
        'updatedAt': value['updatedAt'],
    };
}
