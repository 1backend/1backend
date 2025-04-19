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
import { StableDiffusionTxt2ImgRequestFromJSON, StableDiffusionTxt2ImgRequestToJSON, } from './StableDiffusionTxt2ImgRequest';
/**
 * Check if a given object implements the PromptSvcStableDiffusionParameters interface.
 */
export function instanceOfPromptSvcStableDiffusionParameters(value) {
    return true;
}
export function PromptSvcStableDiffusionParametersFromJSON(json) {
    return PromptSvcStableDiffusionParametersFromJSONTyped(json, false);
}
export function PromptSvcStableDiffusionParametersFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'txt2Img': json['txt2Img'] == null ? undefined : StableDiffusionTxt2ImgRequestFromJSON(json['txt2Img']),
    };
}
export function PromptSvcStableDiffusionParametersToJSON(json) {
    return PromptSvcStableDiffusionParametersToJSONTyped(json, false);
}
export function PromptSvcStableDiffusionParametersToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'txt2Img': StableDiffusionTxt2ImgRequestToJSON(value['txt2Img']),
    };
}
