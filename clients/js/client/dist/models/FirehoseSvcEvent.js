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
/**
 * Check if a given object implements the FirehoseSvcEvent interface.
 */
export function instanceOfFirehoseSvcEvent(value) {
    return true;
}
export function FirehoseSvcEventFromJSON(json) {
    return FirehoseSvcEventFromJSONTyped(json, false);
}
export function FirehoseSvcEventFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'data': json['data'] == null ? undefined : json['data'],
        'name': json['name'] == null ? undefined : json['name'],
    };
}
export function FirehoseSvcEventToJSON(json) {
    return FirehoseSvcEventToJSONTyped(json, false);
}
export function FirehoseSvcEventToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'data': value['data'],
        'name': value['name'],
    };
}
