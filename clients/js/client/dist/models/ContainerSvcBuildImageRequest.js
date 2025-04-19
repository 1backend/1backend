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
 * Check if a given object implements the ContainerSvcBuildImageRequest interface.
 */
export function instanceOfContainerSvcBuildImageRequest(value) {
    if (!('contextPath' in value) || value['contextPath'] === undefined)
        return false;
    if (!('name' in value) || value['name'] === undefined)
        return false;
    return true;
}
export function ContainerSvcBuildImageRequestFromJSON(json) {
    return ContainerSvcBuildImageRequestFromJSONTyped(json, false);
}
export function ContainerSvcBuildImageRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'contextPath': json['contextPath'],
        'dockerfilePath': json['dockerfilePath'] == null ? undefined : json['dockerfilePath'],
        'name': json['name'],
    };
}
export function ContainerSvcBuildImageRequestToJSON(json) {
    return ContainerSvcBuildImageRequestToJSONTyped(json, false);
}
export function ContainerSvcBuildImageRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'contextPath': value['contextPath'],
        'dockerfilePath': value['dockerfilePath'],
        'name': value['name'],
    };
}
