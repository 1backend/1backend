/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ConfigSvcConfig interface.
 */
export function instanceOfConfigSvcConfig(value) {
    if (!('data' in value) || value['data'] === undefined)
        return false;
    return true;
}
export function ConfigSvcConfigFromJSON(json) {
    return ConfigSvcConfigFromJSONTyped(json, false);
}
export function ConfigSvcConfigFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'data': json['data'],
        'dataJson': json['dataJson'] == null ? undefined : json['dataJson'],
        'id': json['id'] == null ? undefined : json['id'],
        'namespace': json['namespace'] == null ? undefined : json['namespace'],
    };
}
export function ConfigSvcConfigToJSON(json) {
    return ConfigSvcConfigToJSONTyped(json, false);
}
export function ConfigSvcConfigToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'data': value['data'],
        'dataJson': value['dataJson'],
        'id': value['id'],
        'namespace': value['namespace'],
    };
}
