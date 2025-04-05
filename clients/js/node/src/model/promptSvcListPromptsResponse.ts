/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { PromptSvcPrompt } from './promptSvcPrompt';

export class PromptSvcListPromptsResponse {
    'after'?: object;
    'count'?: number;
    'prompts'?: Array<PromptSvcPrompt>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "after",
            "baseName": "after",
            "type": "object"
        },
        {
            "name": "count",
            "baseName": "count",
            "type": "number"
        },
        {
            "name": "prompts",
            "baseName": "prompts",
            "type": "Array<PromptSvcPrompt>"
        }    ];

    static getAttributeTypeMap() {
        return PromptSvcListPromptsResponse.attributeTypeMap;
    }
}

