/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { ContainerSvcContainer } from './containerSvcContainer';

export class ContainerSvcListContainersResponse {
    'containers'?: Array<ContainerSvcContainer>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "containers",
            "baseName": "containers",
            "type": "Array<ContainerSvcContainer>"
        }    ];

    static getAttributeTypeMap() {
        return ContainerSvcListContainersResponse.attributeTypeMap;
    }
}

