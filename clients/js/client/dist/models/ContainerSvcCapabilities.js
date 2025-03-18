/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ContainerSvcCapabilities interface.
 */
export function instanceOfContainerSvcCapabilities(value) {
    return true;
}
export function ContainerSvcCapabilitiesFromJSON(json) {
    return ContainerSvcCapabilitiesFromJSONTyped(json, false);
}
export function ContainerSvcCapabilitiesFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'gpuEnabled': json['gpuEnabled'] == null ? undefined : json['gpuEnabled'],
    };
}
export function ContainerSvcCapabilitiesToJSON(json) {
    return ContainerSvcCapabilitiesToJSONTyped(json, false);
}
export function ContainerSvcCapabilitiesToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'gpuEnabled': value['gpuEnabled'],
    };
}
