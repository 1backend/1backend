/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { PolicySvcRateLimitParameters } from './PolicySvcRateLimitParameters';
import type { PolicySvcBlocklistParameters } from './PolicySvcBlocklistParameters';
/**
 *
 * @export
 * @interface PolicySvcParameters
 */
export interface PolicySvcParameters {
    /**
     *
     * @type {PolicySvcBlocklistParameters}
     * @memberof PolicySvcParameters
     */
    blocklist?: PolicySvcBlocklistParameters;
    /**
     *
     * @type {PolicySvcRateLimitParameters}
     * @memberof PolicySvcParameters
     */
    rateLimit?: PolicySvcRateLimitParameters;
}
/**
 * Check if a given object implements the PolicySvcParameters interface.
 */
export declare function instanceOfPolicySvcParameters(value: object): value is PolicySvcParameters;
export declare function PolicySvcParametersFromJSON(json: any): PolicySvcParameters;
export declare function PolicySvcParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcParameters;
export declare function PolicySvcParametersToJSON(json: any): PolicySvcParameters;
export declare function PolicySvcParametersToJSONTyped(value?: PolicySvcParameters | null, ignoreDiscriminator?: boolean): any;
