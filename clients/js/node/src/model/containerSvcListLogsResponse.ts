/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { ContainerSvcLog } from './containerSvcLog';

export class ContainerSvcListLogsResponse {
    'logs'?: Array<ContainerSvcLog>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "logs",
            "baseName": "logs",
            "type": "Array<ContainerSvcLog>"
        }    ];

    static getAttributeTypeMap() {
        return ContainerSvcListLogsResponse.attributeTypeMap;
    }
}

