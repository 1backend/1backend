/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class PromptSvcRemovePromptRequest {
    'promptId'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "promptId",
            "baseName": "promptId",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return PromptSvcRemovePromptRequest.attributeTypeMap;
    }
}

