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
import { PolicySvcRateLimitParametersFromJSON, PolicySvcRateLimitParametersToJSON, } from './PolicySvcRateLimitParameters';
import { PolicySvcBlocklistParametersFromJSON, PolicySvcBlocklistParametersToJSON, } from './PolicySvcBlocklistParameters';
/**
 * Check if a given object implements the PolicySvcParameters interface.
 */
export function instanceOfPolicySvcParameters(value) {
    return true;
}
export function PolicySvcParametersFromJSON(json) {
    return PolicySvcParametersFromJSONTyped(json, false);
}
export function PolicySvcParametersFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'blocklist': json['blocklist'] == null ? undefined : PolicySvcBlocklistParametersFromJSON(json['blocklist']),
        'rateLimit': json['rateLimit'] == null ? undefined : PolicySvcRateLimitParametersFromJSON(json['rateLimit']),
    };
}
export function PolicySvcParametersToJSON(json) {
    return PolicySvcParametersToJSONTyped(json, false);
}
export function PolicySvcParametersToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'blocklist': PolicySvcBlocklistParametersToJSON(value['blocklist']),
        'rateLimit': PolicySvcRateLimitParametersToJSON(value['rateLimit']),
    };
}
