/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class ContainerSvcContainerIsRunningResponse {
    'isRunning': boolean;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "isRunning",
            "baseName": "isRunning",
            "type": "boolean"
        }    ];

    static getAttributeTypeMap() {
        return ContainerSvcContainerIsRunningResponse.attributeTypeMap;
    }
}

