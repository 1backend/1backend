/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.6.1
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class ConfigSvcConfig {
    'app'?: string;
    'data': { [key: string]: any; };
    'dataJson'?: string;
    'id'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "app",
            "baseName": "app",
            "type": "string"
        },
        {
            "name": "data",
            "baseName": "data",
            "type": "{ [key: string]: any; }"
        },
        {
            "name": "dataJson",
            "baseName": "dataJson",
            "type": "string"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return ConfigSvcConfig.attributeTypeMap;
    }
}

