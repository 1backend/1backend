/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ModelSvcEnvVar interface.
 */
export function instanceOfModelSvcEnvVar(value) {
    return true;
}
export function ModelSvcEnvVarFromJSON(json) {
    return ModelSvcEnvVarFromJSONTyped(json, false);
}
export function ModelSvcEnvVarFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'key': json['key'] == null ? undefined : json['key'],
        'value': json['value'] == null ? undefined : json['value'],
    };
}
export function ModelSvcEnvVarToJSON(json) {
    return ModelSvcEnvVarToJSONTyped(json, false);
}
export function ModelSvcEnvVarToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'key': value['key'],
        'value': value['value'],
    };
}
