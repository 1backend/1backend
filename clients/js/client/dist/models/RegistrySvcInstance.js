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
import { RegistrySvcInstanceStatusFromJSON, RegistrySvcInstanceStatusToJSON, } from './RegistrySvcInstanceStatus';
/**
 * Check if a given object implements the RegistrySvcInstance interface.
 */
export function instanceOfRegistrySvcInstance(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('status' in value) || value['status'] === undefined)
        return false;
    if (!('url' in value) || value['url'] === undefined)
        return false;
    return true;
}
export function RegistrySvcInstanceFromJSON(json) {
    return RegistrySvcInstanceFromJSONTyped(json, false);
}
export function RegistrySvcInstanceFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'deploymentId': json['deploymentId'] == null ? undefined : json['deploymentId'],
        'details': json['details'] == null ? undefined : json['details'],
        'host': json['host'] == null ? undefined : json['host'],
        'id': json['id'],
        'ip': json['ip'] == null ? undefined : json['ip'],
        'lastHeartbeat': json['lastHeartbeat'] == null ? undefined : json['lastHeartbeat'],
        'nodeUrl': json['nodeUrl'] == null ? undefined : json['nodeUrl'],
        'path': json['path'] == null ? undefined : json['path'],
        'port': json['port'] == null ? undefined : json['port'],
        'scheme': json['scheme'] == null ? undefined : json['scheme'],
        'slug': json['slug'] == null ? undefined : json['slug'],
        'status': RegistrySvcInstanceStatusFromJSON(json['status']),
        'tags': json['tags'] == null ? undefined : json['tags'],
        'url': json['url'],
    };
}
export function RegistrySvcInstanceToJSON(json) {
    return RegistrySvcInstanceToJSONTyped(json, false);
}
export function RegistrySvcInstanceToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'deploymentId': value['deploymentId'],
        'details': value['details'],
        'host': value['host'],
        'id': value['id'],
        'ip': value['ip'],
        'lastHeartbeat': value['lastHeartbeat'],
        'nodeUrl': value['nodeUrl'],
        'path': value['path'],
        'port': value['port'],
        'scheme': value['scheme'],
        'slug': value['slug'],
        'status': RegistrySvcInstanceStatusToJSON(value['status']),
        'tags': value['tags'],
        'url': value['url'],
    };
}
