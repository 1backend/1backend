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
import { ConfigSvcConfigFromJSON, ConfigSvcConfigToJSON, } from './ConfigSvcConfig';
/**
 * Check if a given object implements the ConfigSvcGetConfigResponse interface.
 */
export function instanceOfConfigSvcGetConfigResponse(value) {
    return true;
}
export function ConfigSvcGetConfigResponseFromJSON(json) {
    return ConfigSvcGetConfigResponseFromJSONTyped(json, false);
}
export function ConfigSvcGetConfigResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'config': json['config'] == null ? undefined : ConfigSvcConfigFromJSON(json['config']),
    };
}
export function ConfigSvcGetConfigResponseToJSON(json) {
    return ConfigSvcGetConfigResponseToJSONTyped(json, false);
}
export function ConfigSvcGetConfigResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'config': ConfigSvcConfigToJSON(value['config']),
    };
}
