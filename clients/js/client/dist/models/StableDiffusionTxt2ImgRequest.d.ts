/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
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
 * @interface StableDiffusionTxt2ImgRequest
 */
export interface StableDiffusionTxt2ImgRequest {
    /**
     *
     * @type {{ [key: string]: string; }}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    alwayson_scripts?: {
        [key: string]: string;
    };
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    batch_size?: number;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    cfg_scale?: number;
    /**
     *
     * @type {{ [key: string]: string; }}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    comments?: {
        [key: string]: string;
    };
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    denoising_strength?: number;
    /**
     *
     * @type {boolean}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    disable_extra_networks?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    do_not_save_grid?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    do_not_save_samples?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    enable_hr?: boolean;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    eta?: number;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    firstpass_image?: string;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    firstphase_height?: number;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    firstphase_width?: number;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    force_task_id?: string;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    height?: number;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    hr_checkpoint_name?: string;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    hr_negative_prompt?: string;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    hr_prompt?: string;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    hr_resize_x?: number;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    hr_resize_y?: number;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    hr_sampler_name?: string;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    hr_scale?: number;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    hr_scheduler?: string;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    hr_second_pass_steps?: number;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    hr_upscaler?: string;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    infotext?: string;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    n_iter?: number;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    negative_prompt?: string;
    /**
     *
     * @type {{ [key: string]: string; }}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    override_settings?: {
        [key: string]: string;
    };
    /**
     *
     * @type {boolean}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    override_settings_restore_afterwards?: boolean;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    prompt?: string;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    refiner_checkpoint?: string;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    refiner_switch_at?: number;
    /**
     *
     * @type {boolean}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    restore_faces?: boolean;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    s_churn?: number;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    s_min_uncond?: number;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    s_noise?: number;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    s_tmax?: number;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    s_tmin?: number;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    sampler_index?: string;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    sampler_name?: string;
    /**
     *
     * @type {boolean}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    save_images?: boolean;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    scheduler?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    script_args?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    script_name?: string;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    seed?: number;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    seed_resize_from_h?: number;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    seed_resize_from_w?: number;
    /**
     *
     * @type {boolean}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    send_images?: boolean;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    steps?: number;
    /**
     *
     * @type {Array<string>}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    styles?: Array<string>;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    subseed?: number;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    subseed_strength?: number;
    /**
     *
     * @type {boolean}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    tiling?: boolean;
    /**
     *
     * @type {number}
     * @memberof StableDiffusionTxt2ImgRequest
     */
    width?: number;
}
/**
 * Check if a given object implements the StableDiffusionTxt2ImgRequest interface.
 */
export declare function instanceOfStableDiffusionTxt2ImgRequest(value: object): value is StableDiffusionTxt2ImgRequest;
export declare function StableDiffusionTxt2ImgRequestFromJSON(json: any): StableDiffusionTxt2ImgRequest;
export declare function StableDiffusionTxt2ImgRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): StableDiffusionTxt2ImgRequest;
export declare function StableDiffusionTxt2ImgRequestToJSON(json: any): StableDiffusionTxt2ImgRequest;
export declare function StableDiffusionTxt2ImgRequestToJSONTyped(value?: StableDiffusionTxt2ImgRequest | null, ignoreDiscriminator?: boolean): any;
