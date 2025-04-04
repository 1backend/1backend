/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { ModelSvcModelFromJSON, ModelSvcModelToJSON, } from './ModelSvcModel';
import { ModelSvcPlatformFromJSON, ModelSvcPlatformToJSON, } from './ModelSvcPlatform';
/**
 * Check if a given object implements the ModelSvcGetModelResponse interface.
 */
export function instanceOfModelSvcGetModelResponse(value) {
    if (!('_exists' in value) || value['_exists'] === undefined)
        return false;
    if (!('model' in value) || value['model'] === undefined)
        return false;
    if (!('platform' in value) || value['platform'] === undefined)
        return false;
    return true;
}
export function ModelSvcGetModelResponseFromJSON(json) {
    return ModelSvcGetModelResponseFromJSONTyped(json, false);
}
export function ModelSvcGetModelResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        '_exists': json['exists'],
        'model': ModelSvcModelFromJSON(json['model']),
        'platform': ModelSvcPlatformFromJSON(json['platform']),
    };
}
export function ModelSvcGetModelResponseToJSON(json) {
    return ModelSvcGetModelResponseToJSONTyped(json, false);
}
export function ModelSvcGetModelResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'exists': value['_exists'],
        'model': ModelSvcModelToJSON(value['model']),
        'platform': ModelSvcPlatformToJSON(value['platform']),
    };
}
