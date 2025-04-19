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
import { RegistrySvcInstanceFromJSON, RegistrySvcInstanceToJSON, } from './RegistrySvcInstance';
/**
 * Check if a given object implements the RegistrySvcListInstancesResponse interface.
 */
export function instanceOfRegistrySvcListInstancesResponse(value) {
    return true;
}
export function RegistrySvcListInstancesResponseFromJSON(json) {
    return RegistrySvcListInstancesResponseFromJSONTyped(json, false);
}
export function RegistrySvcListInstancesResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'instances': json['instances'] == null ? undefined : (json['instances'].map(RegistrySvcInstanceFromJSON)),
    };
}
export function RegistrySvcListInstancesResponseToJSON(json) {
    return RegistrySvcListInstancesResponseToJSONTyped(json, false);
}
export function RegistrySvcListInstancesResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'instances': value['instances'] == null ? undefined : (value['instances'].map(RegistrySvcInstanceToJSON)),
    };
}
