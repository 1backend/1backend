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
 * Check if a given object implements the ContainerSvcContainerIsRunningResponse interface.
 */
export function instanceOfContainerSvcContainerIsRunningResponse(value) {
    if (!('isRunning' in value) || value['isRunning'] === undefined)
        return false;
    return true;
}
export function ContainerSvcContainerIsRunningResponseFromJSON(json) {
    return ContainerSvcContainerIsRunningResponseFromJSONTyped(json, false);
}
export function ContainerSvcContainerIsRunningResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'isRunning': json['isRunning'],
    };
}
export function ContainerSvcContainerIsRunningResponseToJSON(json) {
    return ContainerSvcContainerIsRunningResponseToJSONTyped(json, false);
}
export function ContainerSvcContainerIsRunningResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'isRunning': value['isRunning'],
    };
}
