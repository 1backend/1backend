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

export class ContainerSvcListLogsRequest {
    'containerId'?: string;
    'limit'?: number;
    'nodeId'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "containerId",
            "baseName": "containerId",
            "type": "string"
        },
        {
            "name": "limit",
            "baseName": "limit",
            "type": "number"
        },
        {
            "name": "nodeId",
            "baseName": "nodeId",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return ContainerSvcListLogsRequest.attributeTypeMap;
    }
}

