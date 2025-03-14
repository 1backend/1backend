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
import { ModelSvcContainerFromJSON, ModelSvcContainerToJSON, } from './ModelSvcContainer';
/**
 * Check if a given object implements the ModelSvcDefaultParameters interface.
 */
export function instanceOfModelSvcDefaultParameters(value) {
    return true;
}
export function ModelSvcDefaultParametersFromJSON(json) {
    return ModelSvcDefaultParametersFromJSONTyped(json, false);
}
export function ModelSvcDefaultParametersFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'container': json['container'] == null ? undefined : ModelSvcContainerFromJSON(json['container']),
    };
}
export function ModelSvcDefaultParametersToJSON(json) {
    return ModelSvcDefaultParametersToJSONTyped(json, false);
}
export function ModelSvcDefaultParametersToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'container': ModelSvcContainerToJSON(value['container']),
    };
}
