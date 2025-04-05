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
 * Check if a given object implements the DeploySvcResourceLimits interface.
 */
export function instanceOfDeploySvcResourceLimits(value) {
    return true;
}
export function DeploySvcResourceLimitsFromJSON(json) {
    return DeploySvcResourceLimitsFromJSONTyped(json, false);
}
export function DeploySvcResourceLimitsFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'cpu': json['cpu'] == null ? undefined : json['cpu'],
        'memory': json['memory'] == null ? undefined : json['memory'],
        'vram': json['vram'] == null ? undefined : json['vram'],
    };
}
export function DeploySvcResourceLimitsToJSON(json) {
    return DeploySvcResourceLimitsToJSONTyped(json, false);
}
export function DeploySvcResourceLimitsToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'cpu': value['cpu'],
        'memory': value['memory'],
        'vram': value['vram'],
    };
}
