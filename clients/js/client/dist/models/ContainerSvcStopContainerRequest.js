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
 * Check if a given object implements the ContainerSvcStopContainerRequest interface.
 */
export function instanceOfContainerSvcStopContainerRequest(value) {
    return true;
}
export function ContainerSvcStopContainerRequestFromJSON(json) {
    return ContainerSvcStopContainerRequestFromJSONTyped(json, false);
}
export function ContainerSvcStopContainerRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
    };
}
export function ContainerSvcStopContainerRequestToJSON(json) {
    return ContainerSvcStopContainerRequestToJSONTyped(json, false);
}
export function ContainerSvcStopContainerRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'id': value['id'],
        'name': value['name'],
    };
}
