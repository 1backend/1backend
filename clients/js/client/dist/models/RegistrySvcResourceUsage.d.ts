/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { RegistrySvcUsage } from './RegistrySvcUsage';
/**
 *
 * @export
 * @interface RegistrySvcResourceUsage
 */
export interface RegistrySvcResourceUsage {
    /**
     * CPU usage metrics.
     * @type {RegistrySvcUsage}
     * @memberof RegistrySvcResourceUsage
     */
    cpu?: RegistrySvcUsage;
    /**
     * Disk usage metrics.
     * @type {RegistrySvcUsage}
     * @memberof RegistrySvcResourceUsage
     */
    disk?: RegistrySvcUsage;
    /**
     * Memory usage metrics.
     * @type {RegistrySvcUsage}
     * @memberof RegistrySvcResourceUsage
     */
    memory?: RegistrySvcUsage;
}
/**
 * Check if a given object implements the RegistrySvcResourceUsage interface.
 */
export declare function instanceOfRegistrySvcResourceUsage(value: object): value is RegistrySvcResourceUsage;
export declare function RegistrySvcResourceUsageFromJSON(json: any): RegistrySvcResourceUsage;
export declare function RegistrySvcResourceUsageFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcResourceUsage;
export declare function RegistrySvcResourceUsageToJSON(json: any): RegistrySvcResourceUsage;
export declare function RegistrySvcResourceUsageToJSONTyped(value?: RegistrySvcResourceUsage | null, ignoreDiscriminator?: boolean): any;
