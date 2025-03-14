/**
 * OpenOrch
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

export class ContainerSvcNetwork {
    /**
    * IPAddress is the assigned IP address of the container.
    */
    'ipAddress'?: string;
    /**
    * MacAddress is the container\'s MAC address if applicable.
    */
    'macAddress'?: string;
    /**
    * Mode specifies the container\'s network mode (e.g., bridge, host, none, custom).
    */
    'mode'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "ipAddress",
            "baseName": "ipAddress",
            "type": "string"
        },
        {
            "name": "macAddress",
            "baseName": "macAddress",
            "type": "string"
        },
        {
            "name": "mode",
            "baseName": "mode",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return ContainerSvcNetwork.attributeTypeMap;
    }
}

