/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ContainerSvcImagePullableResponse interface.
 */
export function instanceOfContainerSvcImagePullableResponse(value) {
    if (!('pullable' in value) || value['pullable'] === undefined)
        return false;
    return true;
}
export function ContainerSvcImagePullableResponseFromJSON(json) {
    return ContainerSvcImagePullableResponseFromJSONTyped(json, false);
}
export function ContainerSvcImagePullableResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'pullable': json['pullable'],
    };
}
export function ContainerSvcImagePullableResponseToJSON(json) {
    return ContainerSvcImagePullableResponseToJSONTyped(json, false);
}
export function ContainerSvcImagePullableResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'pullable': value['pullable'],
    };
}
