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
 * @interface ModelSvcEnvVar
 */
export interface ModelSvcEnvVar {
    /**
     * Key is the environment variable name.
     * @type {string}
     * @memberof ModelSvcEnvVar
     */
    key?: string;
    /**
     * Value is the environment variable value.
     * @type {string}
     * @memberof ModelSvcEnvVar
     */
    value?: string;
}
/**
 * Check if a given object implements the ModelSvcEnvVar interface.
 */
export declare function instanceOfModelSvcEnvVar(value: object): value is ModelSvcEnvVar;
export declare function ModelSvcEnvVarFromJSON(json: any): ModelSvcEnvVar;
export declare function ModelSvcEnvVarFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcEnvVar;
export declare function ModelSvcEnvVarToJSON(json: any): ModelSvcEnvVar;
export declare function ModelSvcEnvVarToJSONTyped(value?: ModelSvcEnvVar | null, ignoreDiscriminator?: boolean): any;
