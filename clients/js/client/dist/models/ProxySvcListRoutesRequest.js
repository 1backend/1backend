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
 * Check if a given object implements the ProxySvcListRoutesRequest interface.
 */
export function instanceOfProxySvcListRoutesRequest(value) {
    return true;
}
export function ProxySvcListRoutesRequestFromJSON(json) {
    return ProxySvcListRoutesRequestFromJSONTyped(json, false);
}
export function ProxySvcListRoutesRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'ids': json['ids'] == null ? undefined : json['ids'],
    };
}
export function ProxySvcListRoutesRequestToJSON(json) {
    return ProxySvcListRoutesRequestToJSONTyped(json, false);
}
export function ProxySvcListRoutesRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'ids': value['ids'],
    };
}
