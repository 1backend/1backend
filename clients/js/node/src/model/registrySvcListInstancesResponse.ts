/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
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

