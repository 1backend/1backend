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
import { ModelSvcContainerFromJSON, ModelSvcContainerToJSON, } from './ModelSvcContainer';
/**
 * Check if a given object implements the ModelSvcCudaParameters interface.
 */
export function instanceOfModelSvcCudaParameters(value) {
    return true;
}
export function ModelSvcCudaParametersFromJSON(json) {
    return ModelSvcCudaParametersFromJSONTyped(json, false);
}
export function ModelSvcCudaParametersFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'container': json['container'] == null ? undefined : ModelSvcContainerFromJSON(json['container']),
        'cudaVersionPrecision': json['cudaVersionPrecision'] == null ? undefined : json['cudaVersionPrecision'],
        'defaultCudaVersion': json['defaultCudaVersion'] == null ? undefined : json['defaultCudaVersion'],
        'defaultCudnnVersion': json['defaultCudnnVersion'] == null ? undefined : json['defaultCudnnVersion'],
    };
}
export function ModelSvcCudaParametersToJSON(json) {
    return ModelSvcCudaParametersToJSONTyped(json, false);
}
export function ModelSvcCudaParametersToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'container': ModelSvcContainerToJSON(value['container']),
        'cudaVersionPrecision': value['cudaVersionPrecision'],
        'defaultCudaVersion': value['defaultCudaVersion'],
        'defaultCudnnVersion': value['defaultCudnnVersion'],
    };
}
