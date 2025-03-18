/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class RegistrySvcListNodesRequest {
    /**
    * Node IDs to filter on
    */
    'ids'?: Array<string>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "ids",
            "baseName": "ids",
            "type": "Array<string>"
        }    ];

    static getAttributeTypeMap() {
        return RegistrySvcListNodesRequest.attributeTypeMap;
    }
}

