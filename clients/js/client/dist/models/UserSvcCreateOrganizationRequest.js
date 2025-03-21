/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcCreateOrganizationRequest interface.
 */
export function instanceOfUserSvcCreateOrganizationRequest(value) {
    return true;
}
export function UserSvcCreateOrganizationRequestFromJSON(json) {
    return UserSvcCreateOrganizationRequestFromJSONTyped(json, false);
}
export function UserSvcCreateOrganizationRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'slug': json['slug'] == null ? undefined : json['slug'],
    };
}
export function UserSvcCreateOrganizationRequestToJSON(json) {
    return UserSvcCreateOrganizationRequestToJSONTyped(json, false);
}
export function UserSvcCreateOrganizationRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'id': value['id'],
        'name': value['name'],
        'slug': value['slug'],
    };
}
