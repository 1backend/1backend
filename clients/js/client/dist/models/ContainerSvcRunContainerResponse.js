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
import { ContainerSvcPortMappingFromJSON, ContainerSvcPortMappingToJSON, } from './ContainerSvcPortMapping';
/**
 * Check if a given object implements the ContainerSvcRunContainerResponse interface.
 */
export function instanceOfContainerSvcRunContainerResponse(value) {
    return true;
}
export function ContainerSvcRunContainerResponseFromJSON(json) {
    return ContainerSvcRunContainerResponseFromJSONTyped(json, false);
}
export function ContainerSvcRunContainerResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'ports': json['ports'] == null ? undefined : (json['ports'].map(ContainerSvcPortMappingFromJSON)),
        'started': json['started'] == null ? undefined : json['started'],
    };
}
export function ContainerSvcRunContainerResponseToJSON(json) {
    return ContainerSvcRunContainerResponseToJSONTyped(json, false);
}
export function ContainerSvcRunContainerResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'ports': value['ports'] == null ? undefined : (value['ports'].map(ContainerSvcPortMappingToJSON)),
        'started': value['started'],
    };
}
