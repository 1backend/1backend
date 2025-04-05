/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { PromptSvcEngineParameters } from './promptSvcEngineParameters';
import { PromptSvcParameters } from './promptSvcParameters';

export class PromptSvcPromptRequest {
    /**
    * AI engine/platform (eg. Llama, Stable Diffusion) specific parameters
    */
    'engineParameters'?: PromptSvcEngineParameters;
    /**
    * Id is the unique ID of the prompt.
    */
    'id'?: string;
    /**
    * MaxRetries specified how many times the system should retry a prompt when it keeps erroring.
    */
    'maxRetries'?: number;
    /**
    * ModelId is just the 1Backend internal ID of the model.
    */
    'modelId'?: string;
    /**
    * AI engine/platform (eg. Llama, Stable Diffusion) agnostic parameters. Use these high level parameters when you don\'t care about the actual engine, only the functionality (eg. text to image, image to image) it provides.
    */
    'parameters'?: PromptSvcParameters;
    /**
    * Prompt is the message itself eg. \"What\'s a banana?
    */
    'prompt': string;
    /**
    * Sync drives whether prompt add request should wait and hang until the prompt is done executing. By default the prompt just gets put on a queue and the client will just subscribe to a Thread Stream. For quick and dirty scripting however it\'s often times easier to do things synchronously. In those cases set Sync to true.
    */
    'sync'?: boolean;
    /**
    * ThreadId is the ID of the thread a prompt belongs to. Clients subscribe to Thread Streams to see the answer to a prompt, or set `prompt.sync` to true for a blocking answer.
    */
    'threadId'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "engineParameters",
            "baseName": "engineParameters",
            "type": "PromptSvcEngineParameters"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "maxRetries",
            "baseName": "maxRetries",
            "type": "number"
        },
        {
            "name": "modelId",
            "baseName": "modelId",
            "type": "string"
        },
        {
            "name": "parameters",
            "baseName": "parameters",
            "type": "PromptSvcParameters"
        },
        {
            "name": "prompt",
            "baseName": "prompt",
            "type": "string"
        },
        {
            "name": "sync",
            "baseName": "sync",
            "type": "boolean"
        },
        {
            "name": "threadId",
            "baseName": "threadId",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return PromptSvcPromptRequest.attributeTypeMap;
    }
}

