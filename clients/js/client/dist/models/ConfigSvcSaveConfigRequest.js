/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { ConfigSvcConfigFromJSON, ConfigSvcConfigToJSON, } from './ConfigSvcConfig';
/**
 * Check if a given object implements the ConfigSvcSaveConfigRequest interface.
 */
export function instanceOfConfigSvcSaveConfigRequest(value) {
    return true;
}
export function ConfigSvcSaveConfigRequestFromJSON(json) {
    return ConfigSvcSaveConfigRequestFromJSONTyped(json, false);
}
export function ConfigSvcSaveConfigRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'config': json['config'] == null ? undefined : ConfigSvcConfigFromJSON(json['config']),
    };
}
export function ConfigSvcSaveConfigRequestToJSON(json) {
    return ConfigSvcSaveConfigRequestToJSONTyped(json, false);
}
export function ConfigSvcSaveConfigRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'config': ConfigSvcConfigToJSON(value['config']),
    };
}
