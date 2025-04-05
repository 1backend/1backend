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
/**
 *
 * @export
 * @interface PolicySvcBlocklistParameters
 */
export interface PolicySvcBlocklistParameters {
    /**
     *
     * @type {Array<string>}
     * @memberof PolicySvcBlocklistParameters
     */
    blockedIPs?: Array<string>;
}
/**
 * Check if a given object implements the PolicySvcBlocklistParameters interface.
 */
export declare function instanceOfPolicySvcBlocklistParameters(value: object): value is PolicySvcBlocklistParameters;
export declare function PolicySvcBlocklistParametersFromJSON(json: any): PolicySvcBlocklistParameters;
export declare function PolicySvcBlocklistParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcBlocklistParameters;
export declare function PolicySvcBlocklistParametersToJSON(json: any): PolicySvcBlocklistParameters;
export declare function PolicySvcBlocklistParametersToJSONTyped(value?: PolicySvcBlocklistParameters | null, ignoreDiscriminator?: boolean): any;
