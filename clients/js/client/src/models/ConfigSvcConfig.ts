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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ConfigSvcConfig
 */
export interface ConfigSvcConfig {
    /**
     * 
     * @type {{ [key: string]: any; }}
     * @memberof ConfigSvcConfig
     */
    data: { [key: string]: any; };
    /**
     * 
     * @type {string}
     * @memberof ConfigSvcConfig
     */
    dataJson?: string;
    /**
     * 
     * @type {string}
     * @memberof ConfigSvcConfig
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof ConfigSvcConfig
     */
    namespace?: string;
}

/**
 * Check if a given object implements the ConfigSvcConfig interface.
 */
export function instanceOfConfigSvcConfig(value: object): value is ConfigSvcConfig {
    if (!('data' in value) || value['data'] === undefined) return false;
    return true;
}

export function ConfigSvcConfigFromJSON(json: any): ConfigSvcConfig {
    return ConfigSvcConfigFromJSONTyped(json, false);
}

export function ConfigSvcConfigFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigSvcConfig {
    if (json == null) {
        return json;
    }
    return {
        
        'data': json['data'],
        'dataJson': json['dataJson'] == null ? undefined : json['dataJson'],
        'id': json['id'] == null ? undefined : json['id'],
        'namespace': json['namespace'] == null ? undefined : json['namespace'],
    };
}

export function ConfigSvcConfigToJSON(json: any): ConfigSvcConfig {
    return ConfigSvcConfigToJSONTyped(json, false);
}

export function ConfigSvcConfigToJSONTyped(value?: ConfigSvcConfig | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'data': value['data'],
        'dataJson': value['dataJson'],
        'id': value['id'],
        'namespace': value['namespace'],
    };
}

