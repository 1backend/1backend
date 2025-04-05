/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class ContainerSvcCapabilities {
    /**
    * GPUEnabled specifies whether GPU support is enabled for the container.
    */
    'gpuEnabled'?: boolean;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "gpuEnabled",
            "baseName": "gpuEnabled",
            "type": "boolean"
        }    ];

    static getAttributeTypeMap() {
        return ContainerSvcCapabilities.attributeTypeMap;
    }
}

