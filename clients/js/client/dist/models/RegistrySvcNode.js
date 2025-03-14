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
import { RegistrySvcGPUFromJSON, RegistrySvcGPUToJSON, } from './RegistrySvcGPU';
import { RegistrySvcResourceUsageFromJSON, RegistrySvcResourceUsageToJSON, } from './RegistrySvcResourceUsage';
/**
 * Check if a given object implements the RegistrySvcNode interface.
 */
export function instanceOfRegistrySvcNode(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('url' in value) || value['url'] === undefined)
        return false;
    return true;
}
export function RegistrySvcNodeFromJSON(json) {
    return RegistrySvcNodeFromJSONTyped(json, false);
}
export function RegistrySvcNodeFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'availabilityZone': json['availabilityZone'] == null ? undefined : json['availabilityZone'],
        'gpus': json['gpus'] == null ? undefined : (json['gpus'].map(RegistrySvcGPUFromJSON)),
        'id': json['id'],
        'lastHeartbeat': json['lastHeartbeat'] == null ? undefined : json['lastHeartbeat'],
        'region': json['region'] == null ? undefined : json['region'],
        'url': json['url'],
        'usage': json['usage'] == null ? undefined : RegistrySvcResourceUsageFromJSON(json['usage']),
    };
}
export function RegistrySvcNodeToJSON(json) {
    return RegistrySvcNodeToJSONTyped(json, false);
}
export function RegistrySvcNodeToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'availabilityZone': value['availabilityZone'],
        'gpus': value['gpus'] == null ? undefined : (value['gpus'].map(RegistrySvcGPUToJSON)),
        'id': value['id'],
        'lastHeartbeat': value['lastHeartbeat'],
        'region': value['region'],
        'url': value['url'],
        'usage': RegistrySvcResourceUsageToJSON(value['usage']),
    };
}
