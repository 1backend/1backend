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
export declare function instanceOfPromptSvcTextToImageParameters(value: object): value is PromptSvcTextToImageParameters;
export declare function PromptSvcTextToImageParametersFromJSON(json: any): PromptSvcTextToImageParameters;
export declare function PromptSvcTextToImageParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcTextToImageParameters;
export declare function PromptSvcTextToImageParametersToJSON(json: any): PromptSvcTextToImageParameters;
export declare function PromptSvcTextToImageParametersToJSONTyped(value?: PromptSvcTextToImageParameters | null, ignoreDiscriminator?: boolean): any;
