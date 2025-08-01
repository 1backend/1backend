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
import { UserSvcAuthTokenFromJSON, UserSvcAuthTokenToJSON, } from './UserSvcAuthToken';
import { UserSvcOrganizationFromJSON, UserSvcOrganizationToJSON, } from './UserSvcOrganization';
/**
 * Check if a given object implements the UserSvcSaveOrganizationResponse interface.
 */
export function instanceOfUserSvcSaveOrganizationResponse(value) {
    if (!('organization' in value) || value['organization'] === undefined)
        return false;
    if (!('token' in value) || value['token'] === undefined)
        return false;
    return true;
}
export function UserSvcSaveOrganizationResponseFromJSON(json) {
    return UserSvcSaveOrganizationResponseFromJSONTyped(json, false);
}
export function UserSvcSaveOrganizationResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'organization': UserSvcOrganizationFromJSON(json['organization']),
        'token': UserSvcAuthTokenFromJSON(json['token']),
    };
}
export function UserSvcSaveOrganizationResponseToJSON(json) {
    return UserSvcSaveOrganizationResponseToJSONTyped(json, false);
}
export function UserSvcSaveOrganizationResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'organization': UserSvcOrganizationToJSON(value['organization']),
        'token': UserSvcAuthTokenToJSON(value['token']),
    };
}
