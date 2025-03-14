/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { RegistrySvcDefinitionFromJSON, RegistrySvcDefinitionToJSON, } from './RegistrySvcDefinition';
/**
 * Check if a given object implements the RegistrySvcListDefinitionsResponse interface.
 */
export function instanceOfRegistrySvcListDefinitionsResponse(value) {
    return true;
}
export function RegistrySvcListDefinitionsResponseFromJSON(json) {
    return RegistrySvcListDefinitionsResponseFromJSONTyped(json, false);
}
export function RegistrySvcListDefinitionsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'definitions': json['definitions'] == null ? undefined : (json['definitions'].map(RegistrySvcDefinitionFromJSON)),
    };
}
export function RegistrySvcListDefinitionsResponseToJSON(json) {
    return RegistrySvcListDefinitionsResponseToJSONTyped(json, false);
}
export function RegistrySvcListDefinitionsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'definitions': value['definitions'] == null ? undefined : (value['definitions'].map(RegistrySvcDefinitionToJSON)),
    };
}
