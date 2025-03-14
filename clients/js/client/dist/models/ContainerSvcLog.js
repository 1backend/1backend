/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ContainerSvcLog interface.
 */
export function instanceOfContainerSvcLog(value) {
    return true;
}
export function ContainerSvcLogFromJSON(json) {
    return ContainerSvcLogFromJSONTyped(json, false);
}
export function ContainerSvcLogFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'containerId': json['containerId'] == null ? undefined : json['containerId'],
        'content': json['content'] == null ? undefined : json['content'],
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'id': json['id'] == null ? undefined : json['id'],
        'nodeId': json['nodeId'] == null ? undefined : json['nodeId'],
    };
}
export function ContainerSvcLogToJSON(json) {
    return ContainerSvcLogToJSONTyped(json, false);
}
export function ContainerSvcLogToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'containerId': value['containerId'],
        'content': value['content'],
        'createdAt': value['createdAt'],
        'id': value['id'],
        'nodeId': value['nodeId'],
    };
}
