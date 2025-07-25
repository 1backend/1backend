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
/**
 * Check if a given object implements the ContainerSvcPortMapping interface.
 */
export function instanceOfContainerSvcPortMapping(value) {
    if (!('host' in value) || value['host'] === undefined)
        return false;
    if (!('internal' in value) || value['internal'] === undefined)
        return false;
    return true;
}
export function ContainerSvcPortMappingFromJSON(json) {
    return ContainerSvcPortMappingFromJSONTyped(json, false);
}
export function ContainerSvcPortMappingFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'host': json['host'],
        'internal': json['internal'],
    };
}
export function ContainerSvcPortMappingToJSON(json) {
    return ContainerSvcPortMappingToJSONTyped(json, false);
}
export function ContainerSvcPortMappingToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'host': value['host'],
        'internal': value['internal'],
    };
}
