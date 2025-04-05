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
 * Check if a given object implements the UserSvcInvite interface.
 */
export function instanceOfUserSvcInvite(value) {
    if (!('contactId' in value) || value['contactId'] === undefined)
        return false;
    if (!('createdAt' in value) || value['createdAt'] === undefined)
        return false;
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('ownerIds' in value) || value['ownerIds'] === undefined)
        return false;
    if (!('roleId' in value) || value['roleId'] === undefined)
        return false;
    return true;
}
export function UserSvcInviteFromJSON(json) {
    return UserSvcInviteFromJSONTyped(json, false);
}
export function UserSvcInviteFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'appliedAt': json['appliedAt'] == null ? undefined : json['appliedAt'],
        'contactId': json['contactId'],
        'createdAt': json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'expiresAt': json['expiresAt'] == null ? undefined : json['expiresAt'],
        'id': json['id'],
        'ownerIds': json['ownerIds'],
        'roleId': json['roleId'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
    };
}
export function UserSvcInviteToJSON(json) {
    return UserSvcInviteToJSONTyped(json, false);
}
export function UserSvcInviteToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'appliedAt': value['appliedAt'],
        'contactId': value['contactId'],
        'createdAt': value['createdAt'],
        'deletedAt': value['deletedAt'],
        'expiresAt': value['expiresAt'],
        'id': value['id'],
        'ownerIds': value['ownerIds'],
        'roleId': value['roleId'],
        'updatedAt': value['updatedAt'],
    };
}
