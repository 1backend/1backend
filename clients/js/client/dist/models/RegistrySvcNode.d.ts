/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { RegistrySvcGPU } from './RegistrySvcGPU';
import type { RegistrySvcResourceUsage } from './RegistrySvcResourceUsage';
/**
 *
 * @export
 * @interface RegistrySvcNode
 */
export interface RegistrySvcNode {
    /**
     * The availability zone of the node
     * @type {string}
     * @memberof RegistrySvcNode
     */
    availabilityZone?: string;
    /**
     * List of GPUs available on the node
     * @type {Array<RegistrySvcGPU>}
     * @memberof RegistrySvcNode
     */
    gpus?: Array<RegistrySvcGPU>;
    /**
     * Required: ID of the instance
     * @type {string}
     * @memberof RegistrySvcNode
     */
    id: string;
    /**
     * Last time the instance gave a sign of life
     * @type {string}
     * @memberof RegistrySvcNode
     */
    lastHeartbeat?: string;
    /**
     * The region of the node
     * @type {string}
     * @memberof RegistrySvcNode
     */
    region?: string;
    /**
     * URL of the daemon running on the node.
     * If not configured defaults to hostname + default 1Backend server port.
     * @type {string}
     * @memberof RegistrySvcNode
     */
    url: string;
    /**
     * Resource usage metrics of the node.
     * @type {RegistrySvcResourceUsage}
     * @memberof RegistrySvcNode
     */
    usage?: RegistrySvcResourceUsage;
}
/**
 * Check if a given object implements the RegistrySvcNode interface.
 */
export declare function instanceOfRegistrySvcNode(value: object): value is RegistrySvcNode;
export declare function RegistrySvcNodeFromJSON(json: any): RegistrySvcNode;
export declare function RegistrySvcNodeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcNode;
export declare function RegistrySvcNodeToJSON(json: any): RegistrySvcNode;
export declare function RegistrySvcNodeToJSONTyped(value?: RegistrySvcNode | null, ignoreDiscriminator?: boolean): any;
