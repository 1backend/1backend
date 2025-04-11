/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
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
export declare const PolicySvcTemplateId: {
    readonly TemplateIdRateLimit: "rate-limit";
    readonly TemplateIdBlocklist: "blocklist";
};
export type PolicySvcTemplateId = typeof PolicySvcTemplateId[keyof typeof PolicySvcTemplateId];
export declare function instanceOfPolicySvcTemplateId(value: any): boolean;
export declare function PolicySvcTemplateIdFromJSON(json: any): PolicySvcTemplateId;
export declare function PolicySvcTemplateIdFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcTemplateId;
export declare function PolicySvcTemplateIdToJSON(value?: PolicySvcTemplateId | null): any;
export declare function PolicySvcTemplateIdToJSONTyped(value: any, ignoreDiscriminator: boolean): PolicySvcTemplateId;
