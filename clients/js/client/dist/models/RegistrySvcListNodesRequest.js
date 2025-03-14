/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the RegistrySvcListNodesRequest interface.
 */
export function instanceOfRegistrySvcListNodesRequest(value) {
    return true;
}
export function RegistrySvcListNodesRequestFromJSON(json) {
    return RegistrySvcListNodesRequestFromJSONTyped(json, false);
}
export function RegistrySvcListNodesRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'ids': json['ids'] == null ? undefined : json['ids'],
    };
}
export function RegistrySvcListNodesRequestToJSON(json) {
    return RegistrySvcListNodesRequestToJSONTyped(json, false);
}
export function RegistrySvcListNodesRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'ids': value['ids'],
    };
}
