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
 * @interface RegistrySvcProcess
 */
export interface RegistrySvcProcess {
    /**
     *
     * @type {number}
     * @memberof RegistrySvcProcess
     */
    memoryUsage?: number;
    /**
     *
     * @type {number}
     * @memberof RegistrySvcProcess
     */
    pid?: number;
    /**
     *
     * @type {string}
     * @memberof RegistrySvcProcess
     */
    processName?: string;
}
/**
 * Check if a given object implements the RegistrySvcProcess interface.
 */
export declare function instanceOfRegistrySvcProcess(value: object): value is RegistrySvcProcess;
export declare function RegistrySvcProcessFromJSON(json: any): RegistrySvcProcess;
export declare function RegistrySvcProcessFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcProcess;
export declare function RegistrySvcProcessToJSON(json: any): RegistrySvcProcess;
export declare function RegistrySvcProcessToJSONTyped(value?: RegistrySvcProcess | null, ignoreDiscriminator?: boolean): any;
