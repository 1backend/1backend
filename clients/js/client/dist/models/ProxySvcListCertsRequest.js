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
 * Check if a given object implements the ProxySvcListCertsRequest interface.
 */
export function instanceOfProxySvcListCertsRequest(value) {
    return true;
}
export function ProxySvcListCertsRequestFromJSON(json) {
    return ProxySvcListCertsRequestFromJSONTyped(json, false);
}
export function ProxySvcListCertsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'ids': json['ids'] == null ? undefined : json['ids'],
    };
}
export function ProxySvcListCertsRequestToJSON(json) {
    return ProxySvcListCertsRequestToJSONTyped(json, false);
}
export function ProxySvcListCertsRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'ids': value['ids'],
    };
}
