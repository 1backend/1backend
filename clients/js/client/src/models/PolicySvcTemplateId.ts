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
} as const;
export type PolicySvcTemplateId = typeof PolicySvcTemplateId[keyof typeof PolicySvcTemplateId];


export function instanceOfPolicySvcTemplateId(value: any): boolean {
    for (const key in PolicySvcTemplateId) {
        if (Object.prototype.hasOwnProperty.call(PolicySvcTemplateId, key)) {
            if (PolicySvcTemplateId[key as keyof typeof PolicySvcTemplateId] === value) {
                return true;
            }
        }
    }
    return false;
}

export function PolicySvcTemplateIdFromJSON(json: any): PolicySvcTemplateId {
    return PolicySvcTemplateIdFromJSONTyped(json, false);
}

export function PolicySvcTemplateIdFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcTemplateId {
    return json as PolicySvcTemplateId;
}

export function PolicySvcTemplateIdToJSON(value?: PolicySvcTemplateId | null): any {
    return value as any;
}

export function PolicySvcTemplateIdToJSONTyped(value: any, ignoreDiscriminator: boolean): PolicySvcTemplateId {
    return value as PolicySvcTemplateId;
}

