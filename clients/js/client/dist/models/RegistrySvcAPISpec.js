/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the RegistrySvcAPISpec interface.
 */
export function instanceOfRegistrySvcAPISpec(value) {
    return true;
}
export function RegistrySvcAPISpecFromJSON(json) {
    return RegistrySvcAPISpecFromJSONTyped(json, false);
}
export function RegistrySvcAPISpecFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'metadata': json['metadata'] == null ? undefined : json['metadata'],
        'protocolType': json['protocolType'] == null ? undefined : json['protocolType'],
        'url': json['url'] == null ? undefined : json['url'],
        'version': json['version'] == null ? undefined : json['version'],
    };
}
export function RegistrySvcAPISpecToJSON(json) {
    return RegistrySvcAPISpecToJSONTyped(json, false);
}
export function RegistrySvcAPISpecToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'metadata': value['metadata'],
        'protocolType': value['protocolType'],
        'url': value['url'],
        'version': value['version'],
    };
}
