/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 */
export const PromptSvcPromptStatus = {
    PromptStatusScheduled: 'scheduled',
    PromptStatusRunning: 'running',
    PromptStatusCompleted: 'completed',
    PromptStatusErrored: 'errored',
    PromptStatusAbandoned: 'abandoned',
    PromptStatusCanceled: 'canceled'
};
export function instanceOfPromptSvcPromptStatus(value) {
    for (const key in PromptSvcPromptStatus) {
        if (Object.prototype.hasOwnProperty.call(PromptSvcPromptStatus, key)) {
            if (PromptSvcPromptStatus[key] === value) {
                return true;
            }
        }
    }
    return false;
}
export function PromptSvcPromptStatusFromJSON(json) {
    return PromptSvcPromptStatusFromJSONTyped(json, false);
}
export function PromptSvcPromptStatusFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
export function PromptSvcPromptStatusToJSON(value) {
    return value;
}
export function PromptSvcPromptStatusToJSONTyped(value, ignoreDiscriminator) {
    return value;
}
