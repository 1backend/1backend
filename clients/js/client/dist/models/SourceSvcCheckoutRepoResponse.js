/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the SourceSvcCheckoutRepoResponse interface.
 */
export function instanceOfSourceSvcCheckoutRepoResponse(value) {
    return true;
}
export function SourceSvcCheckoutRepoResponseFromJSON(json) {
    return SourceSvcCheckoutRepoResponseFromJSONTyped(json, false);
}
export function SourceSvcCheckoutRepoResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'dir': json['dir'] == null ? undefined : json['dir'],
    };
}
export function SourceSvcCheckoutRepoResponseToJSON(json) {
    return SourceSvcCheckoutRepoResponseToJSONTyped(json, false);
}
export function SourceSvcCheckoutRepoResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'dir': value['dir'],
    };
}
