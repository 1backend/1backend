/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
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
export const PolicySvcTemplateId = {
    TemplateIdRateLimit: 'rate-limit',
    TemplateIdBlocklist: 'blocklist'
};
export function instanceOfPolicySvcTemplateId(value) {
    for (const key in PolicySvcTemplateId) {
        if (Object.prototype.hasOwnProperty.call(PolicySvcTemplateId, key)) {
            if (PolicySvcTemplateId[key] === value) {
                return true;
            }
        }
    }
    return false;
}
export function PolicySvcTemplateIdFromJSON(json) {
    return PolicySvcTemplateIdFromJSONTyped(json, false);
}
export function PolicySvcTemplateIdFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
export function PolicySvcTemplateIdToJSON(value) {
    return value;
}
export function PolicySvcTemplateIdToJSONTyped(value, ignoreDiscriminator) {
    return value;
}
