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
import { RegistrySvcNodeFromJSON, RegistrySvcNodeToJSON, } from './RegistrySvcNode';
/**
 * Check if a given object implements the RegistrySvcNodeSelfResponse interface.
 */
export function instanceOfRegistrySvcNodeSelfResponse(value) {
    if (!('node' in value) || value['node'] === undefined)
        return false;
    return true;
}
export function RegistrySvcNodeSelfResponseFromJSON(json) {
    return RegistrySvcNodeSelfResponseFromJSONTyped(json, false);
}
export function RegistrySvcNodeSelfResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'node': RegistrySvcNodeFromJSON(json['node']),
    };
}
export function RegistrySvcNodeSelfResponseToJSON(json) {
    return RegistrySvcNodeSelfResponseToJSONTyped(json, false);
}
export function RegistrySvcNodeSelfResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'node': RegistrySvcNodeToJSON(value['node']),
    };
}
