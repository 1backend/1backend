/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { PromptSvcLlamaCppParameters } from './promptSvcLlamaCppParameters';
import { PromptSvcStableDiffusionParameters } from './promptSvcStableDiffusionParameters';

export class PromptSvcEngineParameters {
    'llamaCppParameters'?: PromptSvcLlamaCppParameters;
    'stableDiffusion'?: PromptSvcStableDiffusionParameters;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "llamaCppParameters",
            "baseName": "llamaCppParameters",
            "type": "PromptSvcLlamaCppParameters"
        },
        {
            "name": "stableDiffusion",
            "baseName": "stableDiffusion",
            "type": "PromptSvcStableDiffusionParameters"
        }    ];

    static getAttributeTypeMap() {
        return PromptSvcEngineParameters.attributeTypeMap;
    }
}

