/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
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
    data: {
        [key: string]: any;
    };
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
export declare function instanceOfConfigSvcConfig(value: object): value is ConfigSvcConfig;
export declare function ConfigSvcConfigFromJSON(json: any): ConfigSvcConfig;
export declare function ConfigSvcConfigFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigSvcConfig;
export declare function ConfigSvcConfigToJSON(json: any): ConfigSvcConfig;
export declare function ConfigSvcConfigToJSONTyped(value?: ConfigSvcConfig | null, ignoreDiscriminator?: boolean): any;
