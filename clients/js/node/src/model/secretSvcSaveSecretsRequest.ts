/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { SecretSvcSecret } from './secretSvcSecret';

export class SecretSvcSaveSecretsRequest {
    'secrets'?: Array<SecretSvcSecret>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "secrets",
            "baseName": "secrets",
            "type": "Array<SecretSvcSecret>"
        }    ];

    static getAttributeTypeMap() {
        return SecretSvcSaveSecretsRequest.attributeTypeMap;
    }
}

