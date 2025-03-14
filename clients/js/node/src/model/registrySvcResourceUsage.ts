/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { RegistrySvcUsage } from './registrySvcUsage';

export class RegistrySvcResourceUsage {
    /**
    * CPU usage metrics.
    */
    'cpu'?: RegistrySvcUsage;
    /**
    * Disk usage metrics.
    */
    'disk'?: RegistrySvcUsage;
    /**
    * Memory usage metrics.
    */
    'memory'?: RegistrySvcUsage;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "cpu",
            "baseName": "cpu",
            "type": "RegistrySvcUsage"
        },
        {
            "name": "disk",
            "baseName": "disk",
            "type": "RegistrySvcUsage"
        },
        {
            "name": "memory",
            "baseName": "memory",
            "type": "RegistrySvcUsage"
        }    ];

    static getAttributeTypeMap() {
        return RegistrySvcResourceUsage.attributeTypeMap;
    }
}

