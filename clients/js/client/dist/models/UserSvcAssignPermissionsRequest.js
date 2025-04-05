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
import { UserSvcGrantFromJSON, UserSvcGrantToJSON, } from './UserSvcGrant';
/**
 * Check if a given object implements the UserSvcSaveGrantsRequest interface.
 */
export function instanceOfUserSvcSaveGrantsRequest(value) {
    return true;
}
export function UserSvcSaveGrantsRequestFromJSON(json) {
    return UserSvcSaveGrantsRequestFromJSONTyped(json, false);
}
export function UserSvcSaveGrantsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'permissionLinks': json['permissionLinks'] == null ? undefined : (json['permissionLinks'].map(UserSvcGrantFromJSON)),
    };
}
export function UserSvcSaveGrantsRequestToJSON(json) {
    return UserSvcSaveGrantsRequestToJSONTyped(json, false);
}
export function UserSvcSaveGrantsRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'permissionLinks': value['permissionLinks'] == null ? undefined : (value['permissionLinks'].map(UserSvcGrantToJSON)),
    };
}
