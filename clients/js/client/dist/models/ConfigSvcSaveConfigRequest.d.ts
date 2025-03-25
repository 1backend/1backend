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
import type { ConfigSvcConfig } from './ConfigSvcConfig';
/**
 *
 * @export
 * @interface ConfigSvcSaveConfigRequest
 */
export interface ConfigSvcSaveConfigRequest {
    /**
     *
     * @type {ConfigSvcConfig}
     * @memberof ConfigSvcSaveConfigRequest
     */
    config?: ConfigSvcConfig;
}
/**
 * Check if a given object implements the ConfigSvcSaveConfigRequest interface.
 */
export declare function instanceOfConfigSvcSaveConfigRequest(value: object): value is ConfigSvcSaveConfigRequest;
export declare function ConfigSvcSaveConfigRequestFromJSON(json: any): ConfigSvcSaveConfigRequest;
export declare function ConfigSvcSaveConfigRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigSvcSaveConfigRequest;
export declare function ConfigSvcSaveConfigRequestToJSON(json: any): ConfigSvcSaveConfigRequest;
export declare function ConfigSvcSaveConfigRequestToJSONTyped(value?: ConfigSvcSaveConfigRequest | null, ignoreDiscriminator?: boolean): any;
