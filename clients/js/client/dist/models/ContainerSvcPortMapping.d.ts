/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ContainerSvcPortMapping
 */
export interface ContainerSvcPortMapping {
    /**
     *
     * @type {number}
     * @memberof ContainerSvcPortMapping
     */
    host: number;
    /**
     *
     * @type {number}
     * @memberof ContainerSvcPortMapping
     */
    internal: number;
}
/**
 * Check if a given object implements the ContainerSvcPortMapping interface.
 */
export declare function instanceOfContainerSvcPortMapping(value: object): value is ContainerSvcPortMapping;
export declare function ContainerSvcPortMappingFromJSON(json: any): ContainerSvcPortMapping;
export declare function ContainerSvcPortMappingFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcPortMapping;
export declare function ContainerSvcPortMappingToJSON(json: any): ContainerSvcPortMapping;
export declare function ContainerSvcPortMappingToJSONTyped(value?: ContainerSvcPortMapping | null, ignoreDiscriminator?: boolean): any;
