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
import { ContainerSvcLogFromJSON, ContainerSvcLogToJSON, } from './ContainerSvcLog';
/**
 * Check if a given object implements the ContainerSvcListLogsResponse interface.
 */
export function instanceOfContainerSvcListLogsResponse(value) {
    return true;
}
export function ContainerSvcListLogsResponseFromJSON(json) {
    return ContainerSvcListLogsResponseFromJSONTyped(json, false);
}
export function ContainerSvcListLogsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'logs': json['logs'] == null ? undefined : (json['logs'].map(ContainerSvcLogFromJSON)),
    };
}
export function ContainerSvcListLogsResponseToJSON(json) {
    return ContainerSvcListLogsResponseToJSONTyped(json, false);
}
export function ContainerSvcListLogsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'logs': value['logs'] == null ? undefined : (value['logs'].map(ContainerSvcLogToJSON)),
    };
}
