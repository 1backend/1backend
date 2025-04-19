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
 * Check if a given object implements the RegistrySvcRegisterInstanceRequest interface.
 */
export function instanceOfRegistrySvcRegisterInstanceRequest(value) {
    if (!('url' in value) || value['url'] === undefined)
        return false;
    return true;
}
export function RegistrySvcRegisterInstanceRequestFromJSON(json) {
    return RegistrySvcRegisterInstanceRequestFromJSONTyped(json, false);
}
export function RegistrySvcRegisterInstanceRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'deploymentId': json['deploymentId'] == null ? undefined : json['deploymentId'],
        'host': json['host'] == null ? undefined : json['host'],
        'id': json['id'] == null ? undefined : json['id'],
        'ip': json['ip'] == null ? undefined : json['ip'],
        'path': json['path'] == null ? undefined : json['path'],
        'port': json['port'] == null ? undefined : json['port'],
        'scheme': json['scheme'] == null ? undefined : json['scheme'],
        'url': json['url'],
    };
}
export function RegistrySvcRegisterInstanceRequestToJSON(json) {
    return RegistrySvcRegisterInstanceRequestToJSONTyped(json, false);
}
export function RegistrySvcRegisterInstanceRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'deploymentId': value['deploymentId'],
        'host': value['host'],
        'id': value['id'],
        'ip': value['ip'],
        'path': value['path'],
        'port': value['port'],
        'scheme': value['scheme'],
        'url': value['url'],
    };
}
