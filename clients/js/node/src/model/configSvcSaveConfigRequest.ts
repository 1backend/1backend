/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.3
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class ConfigSvcSaveConfigRequest {
    'data'?: { [key: string]: any; };
    'dataJson'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "data",
            "baseName": "data",
            "type": "{ [key: string]: any; }"
        },
        {
            "name": "dataJson",
            "baseName": "dataJson",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return ConfigSvcSaveConfigRequest.attributeTypeMap;
    }
}

