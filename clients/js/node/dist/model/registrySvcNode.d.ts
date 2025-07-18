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
import { RegistrySvcGPU } from './registrySvcGPU';
import { RegistrySvcResourceUsage } from './registrySvcResourceUsage';
export declare class RegistrySvcNode {
    /**
    * The availability zone of the node
    */
    'availabilityZone'?: string;
    /**
    * List of GPUs available on the node
    */
    'gpus'?: Array<RegistrySvcGPU>;
    /**
    * Required: ID of the instance
    */
    'id': string;
    /**
    * Last time the instance gave a sign of life
    */
    'lastHeartbeat'?: string;
    /**
    * The region of the node
    */
    'region'?: string;
    /**
    * URL of the daemon running on the node. If not configured defaults to hostname + default 1Backend server port.
    */
    'url': string;
    /**
    * Resource usage metrics of the node.
    */
    'usage'?: RegistrySvcResourceUsage;
    static discriminator: string | undefined;
    static attributeTypeMap: Array<{
        name: string;
        baseName: string;
        type: string;
    }>;
    static getAttributeTypeMap(): {
        name: string;
        baseName: string;
        type: string;
    }[];
}
