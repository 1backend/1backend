/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcPermitInputFromJSON, UserSvcPermitInputToJSON, } from './UserSvcPermitInput';
/**
 * Check if a given object implements the UserSvcSavePermitsRequest interface.
 */
export function instanceOfUserSvcSavePermitsRequest(value) {
    return true;
}
export function UserSvcSavePermitsRequestFromJSON(json) {
    return UserSvcSavePermitsRequestFromJSONTyped(json, false);
}
export function UserSvcSavePermitsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'permits': json['permits'] == null ? undefined : (json['permits'].map(UserSvcPermitInputFromJSON)),
    };
}
export function UserSvcSavePermitsRequestToJSON(json) {
    return UserSvcSavePermitsRequestToJSONTyped(json, false);
}
export function UserSvcSavePermitsRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'permits': value['permits'] == null ? undefined : (value['permits'].map(UserSvcPermitInputToJSON)),
    };
}
