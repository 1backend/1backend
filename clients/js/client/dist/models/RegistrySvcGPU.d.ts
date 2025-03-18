/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { RegistrySvcProcess } from './RegistrySvcProcess';
/**
 *
 * @export
 * @interface RegistrySvcGPU
 */
export interface RegistrySvcGPU {
    /**
     *
     * @type {string}
     * @memberof RegistrySvcGPU
     */
    busId?: string;
    /**
     *
     * @type {string}
     * @memberof RegistrySvcGPU
     */
    computeMode?: string;
    /**
     *
     * @type {string}
     * @memberof RegistrySvcGPU
     */
    cudaVersion?: string;
    /**
     *
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    gpuUtilization?: number;
    /**
     * Id Node.URL + IntraNodeId
     * @type {string}
     * @memberof RegistrySvcGPU
     */
    id?: string;
    /**
     *
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    intraNodeId?: number;
    /**
     *
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    memoryTotal?: number;
    /**
     *
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    memoryUsage?: number;
    /**
     *
     * @type {string}
     * @memberof RegistrySvcGPU
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof RegistrySvcGPU
     */
    performanceState?: string;
    /**
     *
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    powerCap?: number;
    /**
     *
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    powerUsage?: number;
    /**
     *
     * @type {Array<RegistrySvcProcess>}
     * @memberof RegistrySvcGPU
     */
    processDetails?: Array<RegistrySvcProcess>;
    /**
     *
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    temperature?: number;
}
/**
 * Check if a given object implements the RegistrySvcGPU interface.
 */
export declare function instanceOfRegistrySvcGPU(value: object): value is RegistrySvcGPU;
export declare function RegistrySvcGPUFromJSON(json: any): RegistrySvcGPU;
export declare function RegistrySvcGPUFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcGPU;
export declare function RegistrySvcGPUToJSON(json: any): RegistrySvcGPU;
export declare function RegistrySvcGPUToJSONTyped(value?: RegistrySvcGPU | null, ignoreDiscriminator?: boolean): any;
