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
import { PromptSvcStreamChunk } from './promptSvcStreamChunk';

export class PromptSvcTypesResponse {
    'chunk'?: PromptSvcStreamChunk;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "chunk",
            "baseName": "chunk",
            "type": "PromptSvcStreamChunk"
        }    ];

    static getAttributeTypeMap() {
        return PromptSvcTypesResponse.attributeTypeMap;
    }
}

