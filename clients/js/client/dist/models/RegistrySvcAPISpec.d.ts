/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
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
 * @interface RegistrySvcAPISpec
 */
export interface RegistrySvcAPISpec {
    /**
     * Additional metadata about the API (e.g., author, license, etc.)
     * @type {{ [key: string]: string; }}
     * @memberof RegistrySvcAPISpec
     */
    metadata?: {
        [key: string]: string;
    };
    /**
     * Protocol type (e.g., OpenAPI, Swagger, etc.)
     * @type {string}
     * @memberof RegistrySvcAPISpec
     */
    protocolType?: string;
    /**
     * URL to the OpenAPI file or other API definition
     * @type {string}
     * @memberof RegistrySvcAPISpec
     */
    url?: string;
    /**
     * Version of the API specification (e.g., 3.0.0)
     * @type {string}
     * @memberof RegistrySvcAPISpec
     */
    version?: string;
}
/**
 * Check if a given object implements the RegistrySvcAPISpec interface.
 */
export declare function instanceOfRegistrySvcAPISpec(value: object): value is RegistrySvcAPISpec;
export declare function RegistrySvcAPISpecFromJSON(json: any): RegistrySvcAPISpec;
export declare function RegistrySvcAPISpecFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcAPISpec;
export declare function RegistrySvcAPISpecToJSON(json: any): RegistrySvcAPISpec;
export declare function RegistrySvcAPISpecToJSONTyped(value?: RegistrySvcAPISpec | null, ignoreDiscriminator?: boolean): any;
