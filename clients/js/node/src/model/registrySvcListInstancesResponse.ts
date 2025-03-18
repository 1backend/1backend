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
import { RegistrySvcInstance } from './registrySvcInstance';

export class RegistrySvcListInstancesResponse {
    'instances'?: Array<RegistrySvcInstance>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "instances",
            "baseName": "instances",
            "type": "Array<RegistrySvcInstance>"
        }    ];

    static getAttributeTypeMap() {
        return RegistrySvcListInstancesResponse.attributeTypeMap;
    }
}

