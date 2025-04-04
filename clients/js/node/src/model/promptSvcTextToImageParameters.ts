/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class PromptSvcTextToImageParameters {
    /**
    * Alternative way to specify dimensions (e.g., \"16:9\", \"1:1\").
    */
    'aspectRatio'?: string;
    /**
    * Number of images to generate per batch.
    */
    'batchSize'?: number;
    /**
    * Controls how much variation is introduced in image modifications.
    */
    'denoisingStrength'?: number;
    /**
    * Whether to apply AI-based upscaling.
    */
    'enableUpscaling'?: boolean;
    /**
    * Output format for the generated image (png, jpg, webp, etc.).
    */
    'format'?: string;
    /**
    * How closely the output should follow the prompt.
    */
    'guidanceScale'?: number;
    'height'?: number;
    /**
    * A negative prompt to specify what should be avoided in the image.
    */
    'negativePrompt'?: string;
    /**
    * Number of batches to generate.
    */
    'numIterations'?: number;
    /**
    * The primary prompt for generating the image. Defaults to the top-level prompt if not specified. If both are provided (which should be avoided), this field takes precedence.
    */
    'prompt'?: string;
    /**
    * Preset quality settings (e.g., Low, Medium, High, Ultra).
    */
    'qualityPreset'?: string;
    /**
    * Whether to enhance facial details for portraits.
    */
    'restoreFaces'?: boolean;
    /**
    * Specifies the sampling method used during generation.
    */
    'scheduler'?: string;
    /**
    * Optional seed for reproducibility. If not set, a random seed is used.
    */
    'seed'?: number;
    /**
    * Number of inference steps for image generation.
    */
    'steps'?: number;
    /**
    * List of artistic styles or themes to apply.
    */
    'styles'?: Array<string>;
    /**
    * Image dimensions (width and height in pixels).
    */
    'width'?: number;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "aspectRatio",
            "baseName": "aspectRatio",
            "type": "string"
        },
        {
            "name": "batchSize",
            "baseName": "batchSize",
            "type": "number"
        },
        {
            "name": "denoisingStrength",
            "baseName": "denoisingStrength",
            "type": "number"
        },
        {
            "name": "enableUpscaling",
            "baseName": "enableUpscaling",
            "type": "boolean"
        },
        {
            "name": "format",
            "baseName": "format",
            "type": "string"
        },
        {
            "name": "guidanceScale",
            "baseName": "guidanceScale",
            "type": "number"
        },
        {
            "name": "height",
            "baseName": "height",
            "type": "number"
        },
        {
            "name": "negativePrompt",
            "baseName": "negativePrompt",
            "type": "string"
        },
        {
            "name": "numIterations",
            "baseName": "numIterations",
            "type": "number"
        },
        {
            "name": "prompt",
            "baseName": "prompt",
            "type": "string"
        },
        {
            "name": "qualityPreset",
            "baseName": "qualityPreset",
            "type": "string"
        },
        {
            "name": "restoreFaces",
            "baseName": "restoreFaces",
            "type": "boolean"
        },
        {
            "name": "scheduler",
            "baseName": "scheduler",
            "type": "string"
        },
        {
            "name": "seed",
            "baseName": "seed",
            "type": "number"
        },
        {
            "name": "steps",
            "baseName": "steps",
            "type": "number"
        },
        {
            "name": "styles",
            "baseName": "styles",
            "type": "Array<string>"
        },
        {
            "name": "width",
            "baseName": "width",
            "type": "number"
        }    ];

    static getAttributeTypeMap() {
        return PromptSvcTextToImageParameters.attributeTypeMap;
    }
}

