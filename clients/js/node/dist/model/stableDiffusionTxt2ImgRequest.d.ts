/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export declare class StableDiffusionTxt2ImgRequest {
    'alwayson_scripts'?: {
        [key: string]: string;
    };
    'batch_size'?: number;
    'cfg_scale'?: number;
    'comments'?: {
        [key: string]: string;
    };
    'denoising_strength'?: number;
    'disable_extra_networks'?: boolean;
    'do_not_save_grid'?: boolean;
    'do_not_save_samples'?: boolean;
    'enable_hr'?: boolean;
    'eta'?: number;
    'firstpass_image'?: string;
    'firstphase_height'?: number;
    'firstphase_width'?: number;
    'force_task_id'?: string;
    'height'?: number;
    'hr_checkpoint_name'?: string;
    'hr_negative_prompt'?: string;
    'hr_prompt'?: string;
    'hr_resize_x'?: number;
    'hr_resize_y'?: number;
    'hr_sampler_name'?: string;
    'hr_scale'?: number;
    'hr_scheduler'?: string;
    'hr_second_pass_steps'?: number;
    'hr_upscaler'?: string;
    'infotext'?: string;
    'n_iter'?: number;
    'negative_prompt'?: string;
    'override_settings'?: {
        [key: string]: string;
    };
    'override_settings_restore_afterwards'?: boolean;
    'prompt'?: string;
    'refiner_checkpoint'?: string;
    'refiner_switch_at'?: number;
    'restore_faces'?: boolean;
    's_churn'?: number;
    's_min_uncond'?: number;
    's_noise'?: number;
    's_tmax'?: number;
    's_tmin'?: number;
    'sampler_index'?: string;
    'sampler_name'?: string;
    'save_images'?: boolean;
    'scheduler'?: string;
    'script_args'?: Array<string>;
    'script_name'?: string;
    'seed'?: number;
    'seed_resize_from_h'?: number;
    'seed_resize_from_w'?: number;
    'send_images'?: boolean;
    'steps'?: number;
    'styles'?: Array<string>;
    'subseed'?: number;
    'subseed_strength'?: number;
    'tiling'?: boolean;
    'width'?: number;
    static discriminator: string | undefined;
    static attributeTypeMap: Array<{
        name: string;
        baseName: string;
        type: string;
    }>;
    static getAttributeTypeMap(): {
        name: string;
        baseName: string;
        type: string;
    }[];
}
