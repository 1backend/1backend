/**
 * 1Backend
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
import { PromptSvcEngineParameters } from './promptSvcEngineParameters';
import { PromptSvcParameters } from './promptSvcParameters';
import { PromptSvcPromptStatus } from './promptSvcPromptStatus';
import { PromptSvcPromptType } from './promptSvcPromptType';

export class PromptSvcPrompt {
    /**
    * CreatedAt is the time of the prompt creation.
    */
    'createdAt'?: string;
    /**
    * AI engine/platform (eg. LlamaCpp, Stable Diffusion) specific parameters
    */
    'engineParameters'?: PromptSvcEngineParameters;
    /**
    * Error that arose during prompt execution, if any.
    */
    'error'?: string;
    /**
    * Id is the unique ID of the prompt.
    */
    'id'?: string;
    /**
    * LastRun is the time of the last prompt run.
    */
    'lastRun'?: string;
    /**
    * MaxRetries specified how many times the system should retry a prompt when it keeps erroring.
    */
    'maxRetries'?: number;
    /**
    * ModelId is just the 1Backend internal ID of the model.
    */
    'modelId'?: string;
    /**
    * AI engine/platform (eg. LlamaCpp, Stable Diffusion) agnostic parameters. Use these high level parameters when you don\'t care about the actual engine, only the functionality (eg. text to image, image to image) it provides.
    */
    'parameters'?: PromptSvcParameters;
    /**
    * Prompt is the message itself eg. \"What\'s a banana?
    */
    'prompt': string;
    'requestMessageId'?: string;
    'responseMessageId'?: string;
    /**
    * RunCount is the number of times the prompt was retried due to errors
    */
    'runCount'?: number;
    /**
    * Status of the prompt.
    */
    'status'?: PromptSvcPromptStatus;
    /**
    * Sync drives whether prompt add request should wait and hang until the prompt is done executing. By default the prompt just gets put on a queue and the client will just subscribe to a Thread Stream. For quick and dirty scripting however it\'s often times easier to do things syncronously. In those cases set Sync to true.
    */
    'sync'?: boolean;
    /**
    * ThreadId is the ID of the thread a prompt belongs to. Clients subscribe to Thread Streams to see the answer to a prompt, or set `prompt.sync` to true for a blocking answer.
    */
    'threadId'?: string;
    /**
    * Type is inferred from the `Parameters` or `EngineParameters` field. Eg. A LLamaCpp prompt will be \"Text-to-Text\", a Stabel Diffusion one will be \"Text-to-Image\" etc.
    */
    'type'?: PromptSvcPromptType;
    /**
    * UpdatedAt is the last time the prompt was updated.
    */
    'updatedAt'?: string;
    /**
    * UserId contains the ID of the user who submitted the prompt.
    */
    'userId'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "string"
        },
        {
            "name": "engineParameters",
            "baseName": "engineParameters",
            "type": "PromptSvcEngineParameters"
        },
        {
            "name": "error",
            "baseName": "error",
            "type": "string"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "lastRun",
            "baseName": "lastRun",
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
            "name": "requestMessageId",
            "baseName": "requestMessageId",
            "type": "string"
        },
        {
            "name": "responseMessageId",
            "baseName": "responseMessageId",
            "type": "string"
        },
        {
            "name": "runCount",
            "baseName": "runCount",
            "type": "number"
        },
        {
            "name": "status",
            "baseName": "status",
            "type": "PromptSvcPromptStatus"
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
        },
        {
            "name": "type",
            "baseName": "type",
            "type": "PromptSvcPromptType"
        },
        {
            "name": "updatedAt",
            "baseName": "updatedAt",
            "type": "string"
        },
        {
            "name": "userId",
            "baseName": "userId",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return PromptSvcPrompt.attributeTypeMap;
    }
}

export namespace PromptSvcPrompt {
}
