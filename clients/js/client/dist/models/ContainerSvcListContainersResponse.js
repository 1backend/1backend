/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { ContainerSvcContainerFromJSON, ContainerSvcContainerToJSON, } from './ContainerSvcContainer';
/**
 * Check if a given object implements the ContainerSvcListContainersResponse interface.
 */
export function instanceOfContainerSvcListContainersResponse(value) {
    return true;
}
export function ContainerSvcListContainersResponseFromJSON(json) {
    return ContainerSvcListContainersResponseFromJSONTyped(json, false);
}
export function ContainerSvcListContainersResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'containers': json['containers'] == null ? undefined : (json['containers'].map(ContainerSvcContainerFromJSON)),
    };
}
export function ContainerSvcListContainersResponseToJSON(json) {
    return ContainerSvcListContainersResponseToJSONTyped(json, false);
}
export function ContainerSvcListContainersResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'containers': value['containers'] == null ? undefined : (value['containers'].map(ContainerSvcContainerToJSON)),
    };
}
