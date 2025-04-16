/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { PolicySvcTemplateIdFromJSON, PolicySvcTemplateIdToJSON, } from './PolicySvcTemplateId';
import { PolicySvcParametersFromJSON, PolicySvcParametersToJSON, } from './PolicySvcParameters';
/**
 * Check if a given object implements the PolicySvcInstance interface.
 */
export function instanceOfPolicySvcInstance(value) {
    if (!('parameters' in value) || value['parameters'] === undefined)
        return false;
    if (!('templateId' in value) || value['templateId'] === undefined)
        return false;
    return true;
}
export function PolicySvcInstanceFromJSON(json) {
    return PolicySvcInstanceFromJSONTyped(json, false);
}
export function PolicySvcInstanceFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'endpoint': json['endpoint'] == null ? undefined : json['endpoint'],
        'id': json['id'] == null ? undefined : json['id'],
        'parameters': PolicySvcParametersFromJSON(json['parameters']),
        'templateId': PolicySvcTemplateIdFromJSON(json['templateId']),
    };
}
export function PolicySvcInstanceToJSON(json) {
    return PolicySvcInstanceToJSONTyped(json, false);
}
export function PolicySvcInstanceToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'endpoint': value['endpoint'],
        'id': value['id'],
        'parameters': PolicySvcParametersToJSON(value['parameters']),
        'templateId': PolicySvcTemplateIdToJSON(value['templateId']),
    };
}
