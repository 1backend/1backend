/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
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
export const PolicySvcScope = {
    ScopeEndpoint: 'endpoint',
    ScopeGlobal: 'global'
};
export function instanceOfPolicySvcScope(value) {
    for (const key in PolicySvcScope) {
        if (Object.prototype.hasOwnProperty.call(PolicySvcScope, key)) {
            if (PolicySvcScope[key] === value) {
                return true;
            }
        }
    }
    return false;
}
export function PolicySvcScopeFromJSON(json) {
    return PolicySvcScopeFromJSONTyped(json, false);
}
export function PolicySvcScopeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
export function PolicySvcScopeToJSON(value) {
    return value;
}
export function PolicySvcScopeToJSONTyped(value, ignoreDiscriminator) {
    return value;
}
