/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class RegistrySvcProcess {
    'memoryUsage'?: number;
    'pid'?: number;
    'processName'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "memoryUsage",
            "baseName": "memoryUsage",
            "type": "number"
        },
        {
            "name": "pid",
            "baseName": "pid",
            "type": "number"
        },
        {
            "name": "processName",
            "baseName": "processName",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return RegistrySvcProcess.attributeTypeMap;
    }
}

