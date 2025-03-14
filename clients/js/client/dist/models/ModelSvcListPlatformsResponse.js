/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { ModelSvcPlatformFromJSON, ModelSvcPlatformToJSON, } from './ModelSvcPlatform';
/**
 * Check if a given object implements the ModelSvcListPlatformsResponse interface.
 */
export function instanceOfModelSvcListPlatformsResponse(value) {
    if (!('platforms' in value) || value['platforms'] === undefined)
        return false;
    return true;
}
export function ModelSvcListPlatformsResponseFromJSON(json) {
    return ModelSvcListPlatformsResponseFromJSONTyped(json, false);
}
export function ModelSvcListPlatformsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'platforms': (json['platforms'].map(ModelSvcPlatformFromJSON)),
    };
}
export function ModelSvcListPlatformsResponseToJSON(json) {
    return ModelSvcListPlatformsResponseToJSONTyped(json, false);
}
export function ModelSvcListPlatformsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'platforms': (value['platforms'].map(ModelSvcPlatformToJSON)),
    };
}
