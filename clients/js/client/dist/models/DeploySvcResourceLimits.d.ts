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
 * @interface DeploySvcResourceLimits
 */
export interface DeploySvcResourceLimits {
    /**
     * CPU limit, e.g., "500m" for 0.5 cores
     * @type {string}
     * @memberof DeploySvcResourceLimits
     */
    cpu?: string;
    /**
     * Memory limit, e.g., "128Mi"
     * @type {string}
     * @memberof DeploySvcResourceLimits
     */
    memory?: string;
    /**
     * Optional: GPU VRAM requirement, e.g., "48GB"
     * @type {string}
     * @memberof DeploySvcResourceLimits
     */
    vram?: string;
}
/**
 * Check if a given object implements the DeploySvcResourceLimits interface.
 */
export declare function instanceOfDeploySvcResourceLimits(value: object): value is DeploySvcResourceLimits;
export declare function DeploySvcResourceLimitsFromJSON(json: any): DeploySvcResourceLimits;
export declare function DeploySvcResourceLimitsFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcResourceLimits;
export declare function DeploySvcResourceLimitsToJSON(json: any): DeploySvcResourceLimits;
export declare function DeploySvcResourceLimitsToJSONTyped(value?: DeploySvcResourceLimits | null, ignoreDiscriminator?: boolean): any;
