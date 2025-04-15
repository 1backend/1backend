/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the StableDiffusionTxt2ImgRequest interface.
 */
export function instanceOfStableDiffusionTxt2ImgRequest(value) {
    return true;
}
export function StableDiffusionTxt2ImgRequestFromJSON(json) {
    return StableDiffusionTxt2ImgRequestFromJSONTyped(json, false);
}
export function StableDiffusionTxt2ImgRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'alwayson_scripts': json['alwayson_scripts'] == null ? undefined : json['alwayson_scripts'],
        'batch_size': json['batch_size'] == null ? undefined : json['batch_size'],
        'cfg_scale': json['cfg_scale'] == null ? undefined : json['cfg_scale'],
        'comments': json['comments'] == null ? undefined : json['comments'],
        'denoising_strength': json['denoising_strength'] == null ? undefined : json['denoising_strength'],
        'disable_extra_networks': json['disable_extra_networks'] == null ? undefined : json['disable_extra_networks'],
        'do_not_save_grid': json['do_not_save_grid'] == null ? undefined : json['do_not_save_grid'],
        'do_not_save_samples': json['do_not_save_samples'] == null ? undefined : json['do_not_save_samples'],
        'enable_hr': json['enable_hr'] == null ? undefined : json['enable_hr'],
        'eta': json['eta'] == null ? undefined : json['eta'],
        'firstpass_image': json['firstpass_image'] == null ? undefined : json['firstpass_image'],
        'firstphase_height': json['firstphase_height'] == null ? undefined : json['firstphase_height'],
        'firstphase_width': json['firstphase_width'] == null ? undefined : json['firstphase_width'],
        'force_task_id': json['force_task_id'] == null ? undefined : json['force_task_id'],
        'height': json['height'] == null ? undefined : json['height'],
        'hr_checkpoint_name': json['hr_checkpoint_name'] == null ? undefined : json['hr_checkpoint_name'],
        'hr_negative_prompt': json['hr_negative_prompt'] == null ? undefined : json['hr_negative_prompt'],
        'hr_prompt': json['hr_prompt'] == null ? undefined : json['hr_prompt'],
        'hr_resize_x': json['hr_resize_x'] == null ? undefined : json['hr_resize_x'],
        'hr_resize_y': json['hr_resize_y'] == null ? undefined : json['hr_resize_y'],
        'hr_sampler_name': json['hr_sampler_name'] == null ? undefined : json['hr_sampler_name'],
        'hr_scale': json['hr_scale'] == null ? undefined : json['hr_scale'],
        'hr_scheduler': json['hr_scheduler'] == null ? undefined : json['hr_scheduler'],
        'hr_second_pass_steps': json['hr_second_pass_steps'] == null ? undefined : json['hr_second_pass_steps'],
        'hr_upscaler': json['hr_upscaler'] == null ? undefined : json['hr_upscaler'],
        'infotext': json['infotext'] == null ? undefined : json['infotext'],
        'n_iter': json['n_iter'] == null ? undefined : json['n_iter'],
        'negative_prompt': json['negative_prompt'] == null ? undefined : json['negative_prompt'],
        'override_settings': json['override_settings'] == null ? undefined : json['override_settings'],
        'override_settings_restore_afterwards': json['override_settings_restore_afterwards'] == null ? undefined : json['override_settings_restore_afterwards'],
        'prompt': json['prompt'] == null ? undefined : json['prompt'],
        'refiner_checkpoint': json['refiner_checkpoint'] == null ? undefined : json['refiner_checkpoint'],
        'refiner_switch_at': json['refiner_switch_at'] == null ? undefined : json['refiner_switch_at'],
        'restore_faces': json['restore_faces'] == null ? undefined : json['restore_faces'],
        's_churn': json['s_churn'] == null ? undefined : json['s_churn'],
        's_min_uncond': json['s_min_uncond'] == null ? undefined : json['s_min_uncond'],
        's_noise': json['s_noise'] == null ? undefined : json['s_noise'],
        's_tmax': json['s_tmax'] == null ? undefined : json['s_tmax'],
        's_tmin': json['s_tmin'] == null ? undefined : json['s_tmin'],
        'sampler_index': json['sampler_index'] == null ? undefined : json['sampler_index'],
        'sampler_name': json['sampler_name'] == null ? undefined : json['sampler_name'],
        'save_images': json['save_images'] == null ? undefined : json['save_images'],
        'scheduler': json['scheduler'] == null ? undefined : json['scheduler'],
        'script_args': json['script_args'] == null ? undefined : json['script_args'],
        'script_name': json['script_name'] == null ? undefined : json['script_name'],
        'seed': json['seed'] == null ? undefined : json['seed'],
        'seed_resize_from_h': json['seed_resize_from_h'] == null ? undefined : json['seed_resize_from_h'],
        'seed_resize_from_w': json['seed_resize_from_w'] == null ? undefined : json['seed_resize_from_w'],
        'send_images': json['send_images'] == null ? undefined : json['send_images'],
        'steps': json['steps'] == null ? undefined : json['steps'],
        'styles': json['styles'] == null ? undefined : json['styles'],
        'subseed': json['subseed'] == null ? undefined : json['subseed'],
        'subseed_strength': json['subseed_strength'] == null ? undefined : json['subseed_strength'],
        'tiling': json['tiling'] == null ? undefined : json['tiling'],
        'width': json['width'] == null ? undefined : json['width'],
    };
}
export function StableDiffusionTxt2ImgRequestToJSON(json) {
    return StableDiffusionTxt2ImgRequestToJSONTyped(json, false);
}
export function StableDiffusionTxt2ImgRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'alwayson_scripts': value['alwayson_scripts'],
        'batch_size': value['batch_size'],
        'cfg_scale': value['cfg_scale'],
        'comments': value['comments'],
        'denoising_strength': value['denoising_strength'],
        'disable_extra_networks': value['disable_extra_networks'],
        'do_not_save_grid': value['do_not_save_grid'],
        'do_not_save_samples': value['do_not_save_samples'],
        'enable_hr': value['enable_hr'],
        'eta': value['eta'],
        'firstpass_image': value['firstpass_image'],
        'firstphase_height': value['firstphase_height'],
        'firstphase_width': value['firstphase_width'],
        'force_task_id': value['force_task_id'],
        'height': value['height'],
        'hr_checkpoint_name': value['hr_checkpoint_name'],
        'hr_negative_prompt': value['hr_negative_prompt'],
        'hr_prompt': value['hr_prompt'],
        'hr_resize_x': value['hr_resize_x'],
        'hr_resize_y': value['hr_resize_y'],
        'hr_sampler_name': value['hr_sampler_name'],
        'hr_scale': value['hr_scale'],
        'hr_scheduler': value['hr_scheduler'],
        'hr_second_pass_steps': value['hr_second_pass_steps'],
        'hr_upscaler': value['hr_upscaler'],
        'infotext': value['infotext'],
        'n_iter': value['n_iter'],
        'negative_prompt': value['negative_prompt'],
        'override_settings': value['override_settings'],
        'override_settings_restore_afterwards': value['override_settings_restore_afterwards'],
        'prompt': value['prompt'],
        'refiner_checkpoint': value['refiner_checkpoint'],
        'refiner_switch_at': value['refiner_switch_at'],
        'restore_faces': value['restore_faces'],
        's_churn': value['s_churn'],
        's_min_uncond': value['s_min_uncond'],
        's_noise': value['s_noise'],
        's_tmax': value['s_tmax'],
        's_tmin': value['s_tmin'],
        'sampler_index': value['sampler_index'],
        'sampler_name': value['sampler_name'],
        'save_images': value['save_images'],
        'scheduler': value['scheduler'],
        'script_args': value['script_args'],
        'script_name': value['script_name'],
        'seed': value['seed'],
        'seed_resize_from_h': value['seed_resize_from_h'],
        'seed_resize_from_w': value['seed_resize_from_w'],
        'send_images': value['send_images'],
        'steps': value['steps'],
        'styles': value['styles'],
        'subseed': value['subseed'],
        'subseed_strength': value['subseed_strength'],
        'tiling': value['tiling'],
        'width': value['width'],
    };
}
