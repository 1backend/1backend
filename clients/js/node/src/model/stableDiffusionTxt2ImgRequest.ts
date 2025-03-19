/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class StableDiffusionTxt2ImgRequest {
    'alwayson_scripts'?: { [key: string]: string; };
    'batch_size'?: number;
    'cfg_scale'?: number;
    'comments'?: { [key: string]: string; };
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
    'override_settings'?: { [key: string]: string; };
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

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "alwayson_scripts",
            "baseName": "alwayson_scripts",
            "type": "{ [key: string]: string; }"
        },
        {
            "name": "batch_size",
            "baseName": "batch_size",
            "type": "number"
        },
        {
            "name": "cfg_scale",
            "baseName": "cfg_scale",
            "type": "number"
        },
        {
            "name": "comments",
            "baseName": "comments",
            "type": "{ [key: string]: string; }"
        },
        {
            "name": "denoising_strength",
            "baseName": "denoising_strength",
            "type": "number"
        },
        {
            "name": "disable_extra_networks",
            "baseName": "disable_extra_networks",
            "type": "boolean"
        },
        {
            "name": "do_not_save_grid",
            "baseName": "do_not_save_grid",
            "type": "boolean"
        },
        {
            "name": "do_not_save_samples",
            "baseName": "do_not_save_samples",
            "type": "boolean"
        },
        {
            "name": "enable_hr",
            "baseName": "enable_hr",
            "type": "boolean"
        },
        {
            "name": "eta",
            "baseName": "eta",
            "type": "number"
        },
        {
            "name": "firstpass_image",
            "baseName": "firstpass_image",
            "type": "string"
        },
        {
            "name": "firstphase_height",
            "baseName": "firstphase_height",
            "type": "number"
        },
        {
            "name": "firstphase_width",
            "baseName": "firstphase_width",
            "type": "number"
        },
        {
            "name": "force_task_id",
            "baseName": "force_task_id",
            "type": "string"
        },
        {
            "name": "height",
            "baseName": "height",
            "type": "number"
        },
        {
            "name": "hr_checkpoint_name",
            "baseName": "hr_checkpoint_name",
            "type": "string"
        },
        {
            "name": "hr_negative_prompt",
            "baseName": "hr_negative_prompt",
            "type": "string"
        },
        {
            "name": "hr_prompt",
            "baseName": "hr_prompt",
            "type": "string"
        },
        {
            "name": "hr_resize_x",
            "baseName": "hr_resize_x",
            "type": "number"
        },
        {
            "name": "hr_resize_y",
            "baseName": "hr_resize_y",
            "type": "number"
        },
        {
            "name": "hr_sampler_name",
            "baseName": "hr_sampler_name",
            "type": "string"
        },
        {
            "name": "hr_scale",
            "baseName": "hr_scale",
            "type": "number"
        },
        {
            "name": "hr_scheduler",
            "baseName": "hr_scheduler",
            "type": "string"
        },
        {
            "name": "hr_second_pass_steps",
            "baseName": "hr_second_pass_steps",
            "type": "number"
        },
        {
            "name": "hr_upscaler",
            "baseName": "hr_upscaler",
            "type": "string"
        },
        {
            "name": "infotext",
            "baseName": "infotext",
            "type": "string"
        },
        {
            "name": "n_iter",
            "baseName": "n_iter",
            "type": "number"
        },
        {
            "name": "negative_prompt",
            "baseName": "negative_prompt",
            "type": "string"
        },
        {
            "name": "override_settings",
            "baseName": "override_settings",
            "type": "{ [key: string]: string; }"
        },
        {
            "name": "override_settings_restore_afterwards",
            "baseName": "override_settings_restore_afterwards",
            "type": "boolean"
        },
        {
            "name": "prompt",
            "baseName": "prompt",
            "type": "string"
        },
        {
            "name": "refiner_checkpoint",
            "baseName": "refiner_checkpoint",
            "type": "string"
        },
        {
            "name": "refiner_switch_at",
            "baseName": "refiner_switch_at",
            "type": "number"
        },
        {
            "name": "restore_faces",
            "baseName": "restore_faces",
            "type": "boolean"
        },
        {
            "name": "s_churn",
            "baseName": "s_churn",
            "type": "number"
        },
        {
            "name": "s_min_uncond",
            "baseName": "s_min_uncond",
            "type": "number"
        },
        {
            "name": "s_noise",
            "baseName": "s_noise",
            "type": "number"
        },
        {
            "name": "s_tmax",
            "baseName": "s_tmax",
            "type": "number"
        },
        {
            "name": "s_tmin",
            "baseName": "s_tmin",
            "type": "number"
        },
        {
            "name": "sampler_index",
            "baseName": "sampler_index",
            "type": "string"
        },
        {
            "name": "sampler_name",
            "baseName": "sampler_name",
            "type": "string"
        },
        {
            "name": "save_images",
            "baseName": "save_images",
            "type": "boolean"
        },
        {
            "name": "scheduler",
            "baseName": "scheduler",
            "type": "string"
        },
        {
            "name": "script_args",
            "baseName": "script_args",
            "type": "Array<string>"
        },
        {
            "name": "script_name",
            "baseName": "script_name",
            "type": "string"
        },
        {
            "name": "seed",
            "baseName": "seed",
            "type": "number"
        },
        {
            "name": "seed_resize_from_h",
            "baseName": "seed_resize_from_h",
            "type": "number"
        },
        {
            "name": "seed_resize_from_w",
            "baseName": "seed_resize_from_w",
            "type": "number"
        },
        {
            "name": "send_images",
            "baseName": "send_images",
            "type": "boolean"
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
            "name": "subseed",
            "baseName": "subseed",
            "type": "number"
        },
        {
            "name": "subseed_strength",
            "baseName": "subseed_strength",
            "type": "number"
        },
        {
            "name": "tiling",
            "baseName": "tiling",
            "type": "boolean"
        },
        {
            "name": "width",
            "baseName": "width",
            "type": "number"
        }    ];

    static getAttributeTypeMap() {
        return StableDiffusionTxt2ImgRequest.attributeTypeMap;
    }
}

