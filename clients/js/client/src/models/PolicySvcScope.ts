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
export const PolicySvcScope = {
    ScopeEndpoint: 'endpoint',
    ScopeGlobal: 'global'
} as const;
export type PolicySvcScope = typeof PolicySvcScope[keyof typeof PolicySvcScope];


export function instanceOfPolicySvcScope(value: any): boolean {
    for (const key in PolicySvcScope) {
        if (Object.prototype.hasOwnProperty.call(PolicySvcScope, key)) {
            if (PolicySvcScope[key as keyof typeof PolicySvcScope] === value) {
                return true;
            }
        }
    }
    return false;
}

export function PolicySvcScopeFromJSON(json: any): PolicySvcScope {
    return PolicySvcScopeFromJSONTyped(json, false);
}

export function PolicySvcScopeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcScope {
    return json as PolicySvcScope;
}

export function PolicySvcScopeToJSON(value?: PolicySvcScope | null): any {
    return value as any;
}

export function PolicySvcScopeToJSONTyped(value: any, ignoreDiscriminator: boolean): PolicySvcScope {
    return value as PolicySvcScope;
}

