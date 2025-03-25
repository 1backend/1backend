/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class ContainerSvcDaemonInfoResponse {
    'address'?: string;
    'available': boolean;
    'error'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "address",
            "baseName": "address",
            "type": "string"
        },
        {
            "name": "available",
            "baseName": "available",
            "type": "boolean"
        },
        {
            "name": "error",
            "baseName": "error",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return ContainerSvcDaemonInfoResponse.attributeTypeMap;
    }
}

