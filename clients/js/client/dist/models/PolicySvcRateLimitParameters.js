/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { PolicySvcScopeFromJSON, PolicySvcScopeToJSON, } from './PolicySvcScope';
import { PolicySvcEntityFromJSON, PolicySvcEntityToJSON, } from './PolicySvcEntity';
/**
 * Check if a given object implements the PolicySvcRateLimitParameters interface.
 */
export function instanceOfPolicySvcRateLimitParameters(value) {
    return true;
}
export function PolicySvcRateLimitParametersFromJSON(json) {
    return PolicySvcRateLimitParametersFromJSONTyped(json, false);
}
export function PolicySvcRateLimitParametersFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'entity': json['entity'] == null ? undefined : PolicySvcEntityFromJSON(json['entity']),
        'maxRequests': json['maxRequests'] == null ? undefined : json['maxRequests'],
        'scope': json['scope'] == null ? undefined : PolicySvcScopeFromJSON(json['scope']),
        'timeWindow': json['timeWindow'] == null ? undefined : json['timeWindow'],
    };
}
export function PolicySvcRateLimitParametersToJSON(json) {
    return PolicySvcRateLimitParametersToJSONTyped(json, false);
}
export function PolicySvcRateLimitParametersToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'entity': PolicySvcEntityToJSON(value['entity']),
        'maxRequests': value['maxRequests'],
        'scope': PolicySvcScopeToJSON(value['scope']),
        'timeWindow': value['timeWindow'],
    };
}
