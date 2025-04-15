/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ContainerSvcLabel interface.
 */
export function instanceOfContainerSvcLabel(value) {
    if (!('key' in value) || value['key'] === undefined)
        return false;
    if (!('value' in value) || value['value'] === undefined)
        return false;
    return true;
}
export function ContainerSvcLabelFromJSON(json) {
    return ContainerSvcLabelFromJSONTyped(json, false);
}
export function ContainerSvcLabelFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'key': json['key'],
        'value': json['value'],
    };
}
export function ContainerSvcLabelToJSON(json) {
    return ContainerSvcLabelToJSONTyped(json, false);
}
export function ContainerSvcLabelToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'key': value['key'],
        'value': value['value'],
    };
}
