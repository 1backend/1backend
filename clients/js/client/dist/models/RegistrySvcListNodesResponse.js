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
import { RegistrySvcNodeFromJSON, RegistrySvcNodeToJSON, } from './RegistrySvcNode';
/**
 * Check if a given object implements the RegistrySvcListNodesResponse interface.
 */
export function instanceOfRegistrySvcListNodesResponse(value) {
    return true;
}
export function RegistrySvcListNodesResponseFromJSON(json) {
    return RegistrySvcListNodesResponseFromJSONTyped(json, false);
}
export function RegistrySvcListNodesResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'nodes': json['nodes'] == null ? undefined : (json['nodes'].map(RegistrySvcNodeFromJSON)),
    };
}
export function RegistrySvcListNodesResponseToJSON(json) {
    return RegistrySvcListNodesResponseToJSONTyped(json, false);
}
export function RegistrySvcListNodesResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'nodes': value['nodes'] == null ? undefined : (value['nodes'].map(RegistrySvcNodeToJSON)),
    };
}
