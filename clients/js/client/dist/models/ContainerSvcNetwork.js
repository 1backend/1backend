/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ContainerSvcNetwork interface.
 */
export function instanceOfContainerSvcNetwork(value) {
    return true;
}
export function ContainerSvcNetworkFromJSON(json) {
    return ContainerSvcNetworkFromJSONTyped(json, false);
}
export function ContainerSvcNetworkFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'ipAddress': json['ipAddress'] == null ? undefined : json['ipAddress'],
        'macAddress': json['macAddress'] == null ? undefined : json['macAddress'],
        'mode': json['mode'] == null ? undefined : json['mode'],
    };
}
export function ContainerSvcNetworkToJSON(json) {
    return ContainerSvcNetworkToJSONTyped(json, false);
}
export function ContainerSvcNetworkToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'ipAddress': value['ipAddress'],
        'macAddress': value['macAddress'],
        'mode': value['mode'],
    };
}
