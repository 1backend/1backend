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
 * Check if a given object implements the UserSvcEnroll interface.
 */
export function instanceOfUserSvcEnroll(value) {
    if (!('createdAt' in value) || value['createdAt'] === undefined)
        return false;
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('role' in value) || value['role'] === undefined)
        return false;
    if (!('updatedAt' in value) || value['updatedAt'] === undefined)
        return false;
    return true;
}
export function UserSvcEnrollFromJSON(json) {
    return UserSvcEnrollFromJSONTyped(json, false);
}
export function UserSvcEnrollFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'app': json['app'] == null ? undefined : json['app'],
        'contactId': json['contactId'] == null ? undefined : json['contactId'],
        'createdAt': json['createdAt'],
        'createdBy': json['createdBy'] == null ? undefined : json['createdBy'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'id': json['id'],
        'role': json['role'],
        'updatedAt': json['updatedAt'],
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}
export function UserSvcEnrollToJSON(json) {
    return UserSvcEnrollToJSONTyped(json, false);
}
export function UserSvcEnrollToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'app': value['app'],
        'contactId': value['contactId'],
        'createdAt': value['createdAt'],
        'createdBy': value['createdBy'],
        'deletedAt': value['deletedAt'],
        'id': value['id'],
        'role': value['role'],
        'updatedAt': value['updatedAt'],
        'userId': value['userId'],
    };
}
