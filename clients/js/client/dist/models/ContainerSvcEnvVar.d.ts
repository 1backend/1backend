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
/**
 *
 * @export
 * @interface ContainerSvcEnvVar
 */
export interface ContainerSvcEnvVar {
    /**
     * Key is the environment variable name.
     * @type {string}
     * @memberof ContainerSvcEnvVar
     */
    key: string;
    /**
     * Value is the environment variable value.
     * @type {string}
     * @memberof ContainerSvcEnvVar
     */
    value: string;
}
/**
 * Check if a given object implements the ContainerSvcEnvVar interface.
 */
export declare function instanceOfContainerSvcEnvVar(value: object): value is ContainerSvcEnvVar;
export declare function ContainerSvcEnvVarFromJSON(json: any): ContainerSvcEnvVar;
export declare function ContainerSvcEnvVarFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcEnvVar;
export declare function ContainerSvcEnvVarToJSON(json: any): ContainerSvcEnvVar;
export declare function ContainerSvcEnvVarToJSONTyped(value?: ContainerSvcEnvVar | null, ignoreDiscriminator?: boolean): any;
