/* tslint:disable */
/* eslint-disable */
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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface PromptSvcTextToImageParameters
 */
export interface PromptSvcTextToImageParameters {
    /**
     * Alternative way to specify dimensions (e.g., "16:9", "1:1").
     * @type {string}
     * @memberof PromptSvcTextToImageParameters
     */
    aspectRatio?: string;
    /**
     * Number of images to generate per batch.
     * @type {number}
     * @memberof PromptSvcTextToImageParameters
     */
    batchSize?: number;
    /**
     * Controls how much variation is introduced in image modifications.
     * @type {number}
     * @memberof PromptSvcTextToImageParameters
     */
    denoisingStrength?: number;
    /**
     * Whether to apply AI-based upscaling.
     * @type {boolean}
     * @memberof PromptSvcTextToImageParameters
     */
    enableUpscaling?: boolean;
    /**
     * Output format for the generated image (png, jpg, webp, etc.).
     * @type {string}
     * @memberof PromptSvcTextToImageParameters
     */
    format?: string;
    /**
     * How closely the output should follow the prompt.
     * @type {number}
     * @memberof PromptSvcTextToImageParameters
     */
    guidanceScale?: number;
    /**
     * 
     * @type {number}
     * @memberof PromptSvcTextToImageParameters
     */
    height?: number;
    /**
     * A negative prompt to specify what should be avoided in the image.
     * @type {string}
     * @memberof PromptSvcTextToImageParameters
     */
    negativePrompt?: string;
    /**
     * Number of batches to generate.
     * @type {number}
     * @memberof PromptSvcTextToImageParameters
     */
    numIterations?: number;
    /**
     * The primary prompt for generating the image.
     * Defaults to the top-level prompt if not specified.
     * If both are provided (which should be avoided), this field takes precedence.
     * @type {string}
     * @memberof PromptSvcTextToImageParameters
     */
    prompt?: string;
    /**
     * Preset quality settings (e.g., Low, Medium, High, Ultra).
     * @type {string}
     * @memberof PromptSvcTextToImageParameters
     */
    qualityPreset?: string;
    /**
     * Whether to enhance facial details for portraits.
     * @type {boolean}
     * @memberof PromptSvcTextToImageParameters
     */
    restoreFaces?: boolean;
    /**
     * Specifies the sampling method used during generation.
     * @type {string}
     * @memberof PromptSvcTextToImageParameters
     */
    scheduler?: string;
    /**
     * Optional seed for reproducibility. If not set, a random seed is used.
     * @type {number}
     * @memberof PromptSvcTextToImageParameters
     */
    seed?: number;
    /**
     * Number of inference steps for image generation.
     * @type {number}
     * @memberof PromptSvcTextToImageParameters
     */
    steps?: number;
    /**
     * List of artistic styles or themes to apply.
     * @type {Array<string>}
     * @memberof PromptSvcTextToImageParameters
     */
    styles?: Array<string>;
    /**
     * Image dimensions (width and height in pixels).
     * @type {number}
     * @memberof PromptSvcTextToImageParameters
     */
    width?: number;
}

/**
 * Check if a given object implements the PromptSvcTextToImageParameters interface.
 */
export function instanceOfPromptSvcTextToImageParameters(value: object): value is PromptSvcTextToImageParameters {
    return true;
}

export function PromptSvcTextToImageParametersFromJSON(json: any): PromptSvcTextToImageParameters {
    return PromptSvcTextToImageParametersFromJSONTyped(json, false);
}

export function PromptSvcTextToImageParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcTextToImageParameters {
    if (json == null) {
        return json;
    }
    return {
        
        'aspectRatio': json['aspectRatio'] == null ? undefined : json['aspectRatio'],
        'batchSize': json['batchSize'] == null ? undefined : json['batchSize'],
        'denoisingStrength': json['denoisingStrength'] == null ? undefined : json['denoisingStrength'],
        'enableUpscaling': json['enableUpscaling'] == null ? undefined : json['enableUpscaling'],
        'format': json['format'] == null ? undefined : json['format'],
        'guidanceScale': json['guidanceScale'] == null ? undefined : json['guidanceScale'],
        'height': json['height'] == null ? undefined : json['height'],
        'negativePrompt': json['negativePrompt'] == null ? undefined : json['negativePrompt'],
        'numIterations': json['numIterations'] == null ? undefined : json['numIterations'],
        'prompt': json['prompt'] == null ? undefined : json['prompt'],
        'qualityPreset': json['qualityPreset'] == null ? undefined : json['qualityPreset'],
        'restoreFaces': json['restoreFaces'] == null ? undefined : json['restoreFaces'],
        'scheduler': json['scheduler'] == null ? undefined : json['scheduler'],
        'seed': json['seed'] == null ? undefined : json['seed'],
        'steps': json['steps'] == null ? undefined : json['steps'],
        'styles': json['styles'] == null ? undefined : json['styles'],
        'width': json['width'] == null ? undefined : json['width'],
    };
}

export function PromptSvcTextToImageParametersToJSON(json: any): PromptSvcTextToImageParameters {
    return PromptSvcTextToImageParametersToJSONTyped(json, false);
}

export function PromptSvcTextToImageParametersToJSONTyped(value?: PromptSvcTextToImageParameters | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'aspectRatio': value['aspectRatio'],
        'batchSize': value['batchSize'],
        'denoisingStrength': value['denoisingStrength'],
        'enableUpscaling': value['enableUpscaling'],
        'format': value['format'],
        'guidanceScale': value['guidanceScale'],
        'height': value['height'],
        'negativePrompt': value['negativePrompt'],
        'numIterations': value['numIterations'],
        'prompt': value['prompt'],
        'qualityPreset': value['qualityPreset'],
        'restoreFaces': value['restoreFaces'],
        'scheduler': value['scheduler'],
        'seed': value['seed'],
        'steps': value['steps'],
        'styles': value['styles'],
        'width': value['width'],
    };
}

