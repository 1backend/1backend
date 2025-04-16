/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class ContainerSvcPortMapping {
    'host': number;
    'internal': number;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "host",
            "baseName": "host",
            "type": "number"
        },
        {
            "name": "internal",
            "baseName": "internal",
            "type": "number"
        }    ];

    static getAttributeTypeMap() {
        return ContainerSvcPortMapping.attributeTypeMap;
    }
}

