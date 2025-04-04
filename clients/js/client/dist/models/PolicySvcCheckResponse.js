/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the PolicySvcCheckResponse interface.
 */
export function instanceOfPolicySvcCheckResponse(value) {
    if (!('allowed' in value) || value['allowed'] === undefined)
        return false;
    return true;
}
export function PolicySvcCheckResponseFromJSON(json) {
    return PolicySvcCheckResponseFromJSONTyped(json, false);
}
export function PolicySvcCheckResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'allowed': json['allowed'],
    };
}
export function PolicySvcCheckResponseToJSON(json) {
    return PolicySvcCheckResponseToJSONTyped(json, false);
}
export function PolicySvcCheckResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'allowed': value['allowed'],
    };
}
