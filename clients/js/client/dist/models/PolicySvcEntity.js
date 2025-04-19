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
/**
 *
 * @export
 */
export const PolicySvcEntity = {
    EntityUserID: 'userId',
    EntityIP: 'ip'
};
export function instanceOfPolicySvcEntity(value) {
    for (const key in PolicySvcEntity) {
        if (Object.prototype.hasOwnProperty.call(PolicySvcEntity, key)) {
            if (PolicySvcEntity[key] === value) {
                return true;
            }
        }
    }
    return false;
}
export function PolicySvcEntityFromJSON(json) {
    return PolicySvcEntityFromJSONTyped(json, false);
}
export function PolicySvcEntityFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
export function PolicySvcEntityToJSON(value) {
    return value;
}
export function PolicySvcEntityToJSONTyped(value, ignoreDiscriminator) {
    return value;
}
