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
import { RegistrySvcLanguage } from './registrySvcLanguage';

export class RegistrySvcClient {
    /**
    * Programming language.
    */
    'language': RegistrySvcLanguage;
    /**
    * The URL of the client.
    */
    'url': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "language",
            "baseName": "language",
            "type": "RegistrySvcLanguage"
        },
        {
            "name": "url",
            "baseName": "url",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return RegistrySvcClient.attributeTypeMap;
    }
}

export namespace RegistrySvcClient {
}
