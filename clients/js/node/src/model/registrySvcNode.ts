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

import { RequestFile } from './models';
import { RegistrySvcGPU } from './registrySvcGPU';
import { RegistrySvcResourceUsage } from './registrySvcResourceUsage';

export class RegistrySvcNode {
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

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "availabilityZone",
            "baseName": "availabilityZone",
            "type": "string"
        },
        {
            "name": "gpus",
            "baseName": "gpus",
            "type": "Array<RegistrySvcGPU>"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "lastHeartbeat",
            "baseName": "lastHeartbeat",
            "type": "string"
        },
        {
            "name": "region",
            "baseName": "region",
            "type": "string"
        },
        {
            "name": "url",
            "baseName": "url",
            "type": "string"
        },
        {
            "name": "usage",
            "baseName": "usage",
            "type": "RegistrySvcResourceUsage"
        }    ];

    static getAttributeTypeMap() {
        return RegistrySvcNode.attributeTypeMap;
    }
}

