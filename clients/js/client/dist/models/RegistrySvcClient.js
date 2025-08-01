/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { RegistrySvcLanguageFromJSON, RegistrySvcLanguageToJSON, } from './RegistrySvcLanguage';
/**
 * Check if a given object implements the RegistrySvcClient interface.
 */
export function instanceOfRegistrySvcClient(value) {
    if (!('language' in value) || value['language'] === undefined)
        return false;
    if (!('url' in value) || value['url'] === undefined)
        return false;
    return true;
}
export function RegistrySvcClientFromJSON(json) {
    return RegistrySvcClientFromJSONTyped(json, false);
}
export function RegistrySvcClientFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'language': RegistrySvcLanguageFromJSON(json['language']),
        'url': json['url'],
    };
}
export function RegistrySvcClientToJSON(json) {
    return RegistrySvcClientToJSONTyped(json, false);
}
export function RegistrySvcClientToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'language': RegistrySvcLanguageToJSON(value['language']),
        'url': value['url'],
    };
}
