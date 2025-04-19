/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { RegistrySvcUsageFromJSON, RegistrySvcUsageToJSON, } from './RegistrySvcUsage';
/**
 * Check if a given object implements the RegistrySvcResourceUsage interface.
 */
export function instanceOfRegistrySvcResourceUsage(value) {
    return true;
}
export function RegistrySvcResourceUsageFromJSON(json) {
    return RegistrySvcResourceUsageFromJSONTyped(json, false);
}
export function RegistrySvcResourceUsageFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'cpu': json['cpu'] == null ? undefined : RegistrySvcUsageFromJSON(json['cpu']),
        'disk': json['disk'] == null ? undefined : RegistrySvcUsageFromJSON(json['disk']),
        'memory': json['memory'] == null ? undefined : RegistrySvcUsageFromJSON(json['memory']),
    };
}
export function RegistrySvcResourceUsageToJSON(json) {
    return RegistrySvcResourceUsageToJSONTyped(json, false);
}
export function RegistrySvcResourceUsageToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'cpu': RegistrySvcUsageToJSON(value['cpu']),
        'disk': RegistrySvcUsageToJSON(value['disk']),
        'memory': RegistrySvcUsageToJSON(value['memory']),
    };
}
