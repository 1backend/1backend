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

export class RegistrySvcRegisterInstanceRequest {
    /**
    * The ID of the deployment that this instance is an instance of.
    */
    'deploymentId'?: string;
    /**
    * Host of the instance address. Required if URL is not provided
    */
    'host'?: string;
    'id'?: string;
    /**
    * IP of the instance address. Optional: to register by IP instead of host
    */
    'ip'?: string;
    /**
    * Path of the instance address. Optional (e.g., \"/api\")
    */
    'path'?: string;
    /**
    * Port of the instance address. Required if URL is not provided
    */
    'port'?: number;
    /**
    * Scheme of the instance address. Required if URL is not provided.
    */
    'scheme'?: string;
    /**
    * Full address URL of the instance.
    */
    'url': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "deploymentId",
            "baseName": "deploymentId",
            "type": "string"
        },
        {
            "name": "host",
            "baseName": "host",
            "type": "string"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "ip",
            "baseName": "ip",
            "type": "string"
        },
        {
            "name": "path",
            "baseName": "path",
            "type": "string"
        },
        {
            "name": "port",
            "baseName": "port",
            "type": "number"
        },
        {
            "name": "scheme",
            "baseName": "scheme",
            "type": "string"
        },
        {
            "name": "url",
            "baseName": "url",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return RegistrySvcRegisterInstanceRequest.attributeTypeMap;
    }
}

