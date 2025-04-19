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
import { ModelSvcModelFromJSON, ModelSvcModelToJSON, } from './ModelSvcModel';
/**
 * Check if a given object implements the ModelSvcListModelsResponse interface.
 */
export function instanceOfModelSvcListModelsResponse(value) {
    return true;
}
export function ModelSvcListModelsResponseFromJSON(json) {
    return ModelSvcListModelsResponseFromJSONTyped(json, false);
}
export function ModelSvcListModelsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'models': json['models'] == null ? undefined : (json['models'].map(ModelSvcModelFromJSON)),
    };
}
export function ModelSvcListModelsResponseToJSON(json) {
    return ModelSvcListModelsResponseToJSONTyped(json, false);
}
export function ModelSvcListModelsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'models': value['models'] == null ? undefined : (value['models'].map(ModelSvcModelToJSON)),
    };
}
