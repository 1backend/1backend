/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the DeploySvcTargetRegion interface.
 */
export function instanceOfDeploySvcTargetRegion(value) {
    return true;
}
export function DeploySvcTargetRegionFromJSON(json) {
    return DeploySvcTargetRegionFromJSONTyped(json, false);
}
export function DeploySvcTargetRegionFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'cluster': json['cluster'] == null ? undefined : json['cluster'],
        'zone': json['zone'] == null ? undefined : json['zone'],
    };
}
export function DeploySvcTargetRegionToJSON(json) {
    return DeploySvcTargetRegionToJSONTyped(json, false);
}
export function DeploySvcTargetRegionToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'cluster': value['cluster'],
        'zone': value['zone'],
    };
}
