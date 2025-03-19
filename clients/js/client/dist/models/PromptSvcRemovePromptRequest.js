/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the PromptSvcRemovePromptRequest interface.
 */
export function instanceOfPromptSvcRemovePromptRequest(value) {
    return true;
}
export function PromptSvcRemovePromptRequestFromJSON(json) {
    return PromptSvcRemovePromptRequestFromJSONTyped(json, false);
}
export function PromptSvcRemovePromptRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'promptId': json['promptId'] == null ? undefined : json['promptId'],
    };
}
export function PromptSvcRemovePromptRequestToJSON(json) {
    return PromptSvcRemovePromptRequestToJSONTyped(json, false);
}
export function PromptSvcRemovePromptRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'promptId': value['promptId'],
    };
}
