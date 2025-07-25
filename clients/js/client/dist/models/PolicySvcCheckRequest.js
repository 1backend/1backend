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
/**
 * Check if a given object implements the PolicySvcCheckRequest interface.
 */
export function instanceOfPolicySvcCheckRequest(value) {
    return true;
}
export function PolicySvcCheckRequestFromJSON(json) {
    return PolicySvcCheckRequestFromJSONTyped(json, false);
}
export function PolicySvcCheckRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'endpoint': json['endpoint'] == null ? undefined : json['endpoint'],
        'ip': json['ip'] == null ? undefined : json['ip'],
        'method': json['method'] == null ? undefined : json['method'],
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}
export function PolicySvcCheckRequestToJSON(json) {
    return PolicySvcCheckRequestToJSONTyped(json, false);
}
export function PolicySvcCheckRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'endpoint': value['endpoint'],
        'ip': value['ip'],
        'method': value['method'],
        'userId': value['userId'],
    };
}
