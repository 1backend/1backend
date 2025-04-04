/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ContainerSvcGetContainerSummaryResponse interface.
 */
export function instanceOfContainerSvcGetContainerSummaryResponse(value) {
    if (!('logs' in value) || value['logs'] === undefined)
        return false;
    if (!('status' in value) || value['status'] === undefined)
        return false;
    if (!('summary' in value) || value['summary'] === undefined)
        return false;
    return true;
}
export function ContainerSvcGetContainerSummaryResponseFromJSON(json) {
    return ContainerSvcGetContainerSummaryResponseFromJSONTyped(json, false);
}
export function ContainerSvcGetContainerSummaryResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'logs': json['logs'],
        'status': json['status'],
        'summary': json['summary'],
    };
}
export function ContainerSvcGetContainerSummaryResponseToJSON(json) {
    return ContainerSvcGetContainerSummaryResponseToJSONTyped(json, false);
}
export function ContainerSvcGetContainerSummaryResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'logs': value['logs'],
        'status': value['status'],
        'summary': value['summary'],
    };
}
