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
 * Check if a given object implements the ModelSvcKeep interface.
 */
export function instanceOfModelSvcKeep(value) {
    return true;
}
export function ModelSvcKeepFromJSON(json) {
    return ModelSvcKeepFromJSONTyped(json, false);
}
export function ModelSvcKeepFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'path': json['path'] == null ? undefined : json['path'],
        'readOnly': json['readOnly'] == null ? undefined : json['readOnly'],
    };
}
export function ModelSvcKeepToJSON(json) {
    return ModelSvcKeepToJSONTyped(json, false);
}
export function ModelSvcKeepToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'path': value['path'],
        'readOnly': value['readOnly'],
    };
}
