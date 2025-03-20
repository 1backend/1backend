/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ContainerSvcErrorResponse interface.
 */
export function instanceOfContainerSvcErrorResponse(value) {
    return true;
}
export function ContainerSvcErrorResponseFromJSON(json) {
    return ContainerSvcErrorResponseFromJSONTyped(json, false);
}
export function ContainerSvcErrorResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'error': json['error'] == null ? undefined : json['error'],
    };
}
export function ContainerSvcErrorResponseToJSON(json) {
    return ContainerSvcErrorResponseToJSONTyped(json, false);
}
export function ContainerSvcErrorResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'error': value['error'],
    };
}
