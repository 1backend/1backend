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
 * Check if a given object implements the ContainerSvcListLogsRequest interface.
 */
export function instanceOfContainerSvcListLogsRequest(value) {
    return true;
}
export function ContainerSvcListLogsRequestFromJSON(json) {
    return ContainerSvcListLogsRequestFromJSONTyped(json, false);
}
export function ContainerSvcListLogsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'containerId': json['containerId'] == null ? undefined : json['containerId'],
        'limit': json['limit'] == null ? undefined : json['limit'],
        'nodeId': json['nodeId'] == null ? undefined : json['nodeId'],
    };
}
export function ContainerSvcListLogsRequestToJSON(json) {
    return ContainerSvcListLogsRequestToJSONTyped(json, false);
}
export function ContainerSvcListLogsRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'containerId': value['containerId'],
        'limit': value['limit'],
        'nodeId': value['nodeId'],
    };
}
