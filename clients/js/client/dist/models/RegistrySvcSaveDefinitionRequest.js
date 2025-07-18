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
import { RegistrySvcDefinitionFromJSON, RegistrySvcDefinitionToJSON, } from './RegistrySvcDefinition';
/**
 * Check if a given object implements the RegistrySvcSaveDefinitionRequest interface.
 */
export function instanceOfRegistrySvcSaveDefinitionRequest(value) {
    return true;
}
export function RegistrySvcSaveDefinitionRequestFromJSON(json) {
    return RegistrySvcSaveDefinitionRequestFromJSONTyped(json, false);
}
export function RegistrySvcSaveDefinitionRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'definition': json['definition'] == null ? undefined : RegistrySvcDefinitionFromJSON(json['definition']),
    };
}
export function RegistrySvcSaveDefinitionRequestToJSON(json) {
    return RegistrySvcSaveDefinitionRequestToJSONTyped(json, false);
}
export function RegistrySvcSaveDefinitionRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'definition': RegistrySvcDefinitionToJSON(value['definition']),
    };
}
