/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class ContainerSvcLog {
    /**
    * ContainerId is the raw underlying container ID. Eg. Docker container id. Node local.
    */
    'containerId'?: string;
    'content'?: string;
    'createdAt'?: string;
    'id'?: string;
    /**
    * Node Id Please see the documentation for the envar OPENORCH_NODE_ID
    */
    'nodeId'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "containerId",
            "baseName": "containerId",
            "type": "string"
        },
        {
            "name": "content",
            "baseName": "content",
            "type": "string"
        },
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "string"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "nodeId",
            "baseName": "nodeId",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return ContainerSvcLog.attributeTypeMap;
    }
}

