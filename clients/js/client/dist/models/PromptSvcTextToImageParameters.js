/* tslint:disable */
/* eslint-disable */
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
/**
 * Check if a given object implements the PromptSvcTextToImageParameters interface.
 */
export function instanceOfPromptSvcTextToImageParameters(value) {
    return true;
}
export function PromptSvcTextToImageParametersFromJSON(json) {
    return PromptSvcTextToImageParametersFromJSONTyped(json, false);
}
export function PromptSvcTextToImageParametersFromJSONTyped(json, ignoreDiscriminator) {
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
export function PromptSvcTextToImageParametersToJSON(json) {
    return PromptSvcTextToImageParametersToJSONTyped(json, false);
}
export function PromptSvcTextToImageParametersToJSONTyped(value, ignoreDiscriminator = false) {
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
