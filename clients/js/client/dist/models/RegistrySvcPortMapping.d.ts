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
/**
 *
 * @export
 * @interface RegistrySvcPortMapping
 */
export interface RegistrySvcPortMapping {
    /**
     *
     * @type {number}
     * @memberof RegistrySvcPortMapping
     */
    host: number;
    /**
     *
     * @type {number}
     * @memberof RegistrySvcPortMapping
     */
    internal: number;
}
/**
 * Check if a given object implements the RegistrySvcPortMapping interface.
 */
export declare function instanceOfRegistrySvcPortMapping(value: object): value is RegistrySvcPortMapping;
export declare function RegistrySvcPortMappingFromJSON(json: any): RegistrySvcPortMapping;
export declare function RegistrySvcPortMappingFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcPortMapping;
export declare function RegistrySvcPortMappingToJSON(json: any): RegistrySvcPortMapping;
export declare function RegistrySvcPortMappingToJSONTyped(value?: RegistrySvcPortMapping | null, ignoreDiscriminator?: boolean): any;
