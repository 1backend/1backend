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
/**
 * Check if a given object implements the DeploySvcAutoScalingConfig interface.
 */
export function instanceOfDeploySvcAutoScalingConfig(value) {
    return true;
}
export function DeploySvcAutoScalingConfigFromJSON(json) {
    return DeploySvcAutoScalingConfigFromJSONTyped(json, false);
}
export function DeploySvcAutoScalingConfigFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'cpuThreshold': json['cpuThreshold'] == null ? undefined : json['cpuThreshold'],
        'maxReplicas': json['maxReplicas'] == null ? undefined : json['maxReplicas'],
        'minReplicas': json['minReplicas'] == null ? undefined : json['minReplicas'],
    };
}
export function DeploySvcAutoScalingConfigToJSON(json) {
    return DeploySvcAutoScalingConfigToJSONTyped(json, false);
}
export function DeploySvcAutoScalingConfigToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'cpuThreshold': value['cpuThreshold'],
        'maxReplicas': value['maxReplicas'],
        'minReplicas': value['minReplicas'],
    };
}
