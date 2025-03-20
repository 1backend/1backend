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
/**
 *
 * @export
 */
export const RegistrySvcInstanceStatus = {
    InstanceStatusUnknown: 'Unknown',
    InstanceStatusHealthy: 'Healthy',
    InstanceStatusDegraded: 'Degraded',
    InstanceStatusUnreachable: 'Unreachable',
    InstanceStatusError: 'Error'
};
export function instanceOfRegistrySvcInstanceStatus(value) {
    for (const key in RegistrySvcInstanceStatus) {
        if (Object.prototype.hasOwnProperty.call(RegistrySvcInstanceStatus, key)) {
            if (RegistrySvcInstanceStatus[key] === value) {
                return true;
            }
        }
    }
    return false;
}
export function RegistrySvcInstanceStatusFromJSON(json) {
    return RegistrySvcInstanceStatusFromJSONTyped(json, false);
}
export function RegistrySvcInstanceStatusFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
export function RegistrySvcInstanceStatusToJSON(value) {
    return value;
}
export function RegistrySvcInstanceStatusToJSONTyped(value, ignoreDiscriminator) {
    return value;
}
