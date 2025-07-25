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

import { RequestFile } from './models';

export class ContainerSvcResources {
    /**
    * CPU cores allocated to the container (e.g., 0.5 = 500m, 2 = 2 cores).
    */
    'cpu'?: number;
    /**
    * Disk space allocated to the container in megabytes.
    */
    'diskMB'?: number;
    /**
    * Memory allocated to the container in megabytes.
    */
    'memoryMB'?: number;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "cpu",
            "baseName": "cpu",
            "type": "number"
        },
        {
            "name": "diskMB",
            "baseName": "diskMB",
            "type": "number"
        },
        {
            "name": "memoryMB",
            "baseName": "memoryMB",
            "type": "number"
        }    ];

    static getAttributeTypeMap() {
        return ContainerSvcResources.attributeTypeMap;
    }
}

