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

export class RegistrySvcAPISpec {
    /**
    * Additional metadata about the API (e.g., author, license, etc.)
    */
    'metadata'?: { [key: string]: string; };
    /**
    * Protocol type (e.g., OpenAPI, Swagger, etc.)
    */
    'protocolType'?: string;
    /**
    * URL to the OpenAPI file or other API definition
    */
    'url'?: string;
    /**
    * Version of the API specification (e.g., 3.0.0)
    */
    'version'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "metadata",
            "baseName": "metadata",
            "type": "{ [key: string]: string; }"
        },
        {
            "name": "protocolType",
            "baseName": "protocolType",
            "type": "string"
        },
        {
            "name": "url",
            "baseName": "url",
            "type": "string"
        },
        {
            "name": "version",
            "baseName": "version",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return RegistrySvcAPISpec.attributeTypeMap;
    }
}

