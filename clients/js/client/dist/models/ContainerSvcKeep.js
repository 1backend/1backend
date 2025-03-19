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
 * Check if a given object implements the ContainerSvcKeep interface.
 */
export function instanceOfContainerSvcKeep(value) {
    if (!('path' in value) || value['path'] === undefined)
        return false;
    return true;
}
export function ContainerSvcKeepFromJSON(json) {
    return ContainerSvcKeepFromJSONTyped(json, false);
}
export function ContainerSvcKeepFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'path': json['path'],
        'readOnly': json['readOnly'] == null ? undefined : json['readOnly'],
    };
}
export function ContainerSvcKeepToJSON(json) {
    return ContainerSvcKeepToJSONTyped(json, false);
}
export function ContainerSvcKeepToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'path': value['path'],
        'readOnly': value['readOnly'],
    };
}
