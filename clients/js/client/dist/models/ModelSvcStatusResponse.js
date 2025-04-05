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
import { ModelSvcModelStatusFromJSON, ModelSvcModelStatusToJSON, } from './ModelSvcModelStatus';
/**
 * Check if a given object implements the ModelSvcStatusResponse interface.
 */
export function instanceOfModelSvcStatusResponse(value) {
    return true;
}
export function ModelSvcStatusResponseFromJSON(json) {
    return ModelSvcStatusResponseFromJSONTyped(json, false);
}
export function ModelSvcStatusResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'status': json['status'] == null ? undefined : ModelSvcModelStatusFromJSON(json['status']),
    };
}
export function ModelSvcStatusResponseToJSON(json) {
    return ModelSvcStatusResponseToJSONTyped(json, false);
}
export function ModelSvcStatusResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'status': ModelSvcModelStatusToJSON(value['status']),
    };
}
