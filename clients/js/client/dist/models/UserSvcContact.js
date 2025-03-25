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
 * Check if a given object implements the UserSvcContact interface.
 */
export function instanceOfUserSvcContact(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('userId' in value) || value['userId'] === undefined)
        return false;
    return true;
}
export function UserSvcContactFromJSON(json) {
    return UserSvcContactFromJSONTyped(json, false);
}
export function UserSvcContactFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'id': json['id'],
        'isPrimary': json['isPrimary'] == null ? undefined : json['isPrimary'],
        'platform': json['platform'] == null ? undefined : json['platform'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
        'userId': json['userId'],
        'value': json['value'] == null ? undefined : json['value'],
        'verified': json['verified'] == null ? undefined : json['verified'],
    };
}
export function UserSvcContactToJSON(json) {
    return UserSvcContactToJSONTyped(json, false);
}
export function UserSvcContactToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'createdAt': value['createdAt'],
        'deletedAt': value['deletedAt'],
        'id': value['id'],
        'isPrimary': value['isPrimary'],
        'platform': value['platform'],
        'updatedAt': value['updatedAt'],
        'userId': value['userId'],
        'value': value['value'],
        'verified': value['verified'],
    };
}
