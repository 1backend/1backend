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
import { UserSvcPermitFromJSON, UserSvcPermitToJSON, } from './UserSvcPermit';
/**
 * Check if a given object implements the UserSvcListPermitsResponse interface.
 */
export function instanceOfUserSvcListPermitsResponse(value) {
    return true;
}
export function UserSvcListPermitsResponseFromJSON(json) {
    return UserSvcListPermitsResponseFromJSONTyped(json, false);
}
export function UserSvcListPermitsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'permits': json['permits'] == null ? undefined : (json['permits'].map(UserSvcPermitFromJSON)),
    };
}
export function UserSvcListPermitsResponseToJSON(json) {
    return UserSvcListPermitsResponseToJSONTyped(json, false);
}
export function UserSvcListPermitsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'permits': value['permits'] == null ? undefined : (value['permits'].map(UserSvcPermitToJSON)),
    };
}
