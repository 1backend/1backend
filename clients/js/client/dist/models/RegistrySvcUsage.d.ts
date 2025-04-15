/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface RegistrySvcUsage
 */
export interface RegistrySvcUsage {
    /**
     * Usage percentage.
     * @type {number}
     * @memberof RegistrySvcUsage
     */
    percent?: number;
    /**
     * Total available amount (in bytes).
     * @type {number}
     * @memberof RegistrySvcUsage
     */
    total?: number;
    /**
     * Used amount (in bytes).
     * @type {number}
     * @memberof RegistrySvcUsage
     */
    used?: number;
}
/**
 * Check if a given object implements the RegistrySvcUsage interface.
 */
export declare function instanceOfRegistrySvcUsage(value: object): value is RegistrySvcUsage;
export declare function RegistrySvcUsageFromJSON(json: any): RegistrySvcUsage;
export declare function RegistrySvcUsageFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcUsage;
export declare function RegistrySvcUsageToJSON(json: any): RegistrySvcUsage;
export declare function RegistrySvcUsageToJSONTyped(value?: RegistrySvcUsage | null, ignoreDiscriminator?: boolean): any;
