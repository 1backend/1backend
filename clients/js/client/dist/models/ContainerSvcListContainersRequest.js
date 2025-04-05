/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ContainerSvcListContainersRequest interface.
 */
export function instanceOfContainerSvcListContainersRequest(value) {
    return true;
}
export function ContainerSvcListContainersRequestFromJSON(json) {
    return ContainerSvcListContainersRequestFromJSONTyped(json, false);
}
export function ContainerSvcListContainersRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'containerId': json['containerId'] == null ? undefined : json['containerId'],
        'limit': json['limit'] == null ? undefined : json['limit'],
        'nodeId': json['nodeId'] == null ? undefined : json['nodeId'],
    };
}
export function ContainerSvcListContainersRequestToJSON(json) {
    return ContainerSvcListContainersRequestToJSONTyped(json, false);
}
export function ContainerSvcListContainersRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'containerId': value['containerId'],
        'limit': value['limit'],
        'nodeId': value['nodeId'],
    };
}
