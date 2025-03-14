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

export class RegistrySvcImageSpec {
    /**
    * InternalPorts are the ports the container will listen on internally
    */
    'internalPorts': Array<number>;
    /**
    * Name is the container image name/URL to use for the container
    */
    'name': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "internalPorts",
            "baseName": "internalPorts",
            "type": "Array<number>"
        },
        {
            "name": "name",
            "baseName": "name",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return RegistrySvcImageSpec.attributeTypeMap;
    }
}

