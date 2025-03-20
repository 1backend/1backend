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
 * @interface ContainerSvcCapabilities
 */
export interface ContainerSvcCapabilities {
    /**
     * GPUEnabled specifies whether GPU support is enabled for the container.
     * @type {boolean}
     * @memberof ContainerSvcCapabilities
     */
    gpuEnabled?: boolean;
}
/**
 * Check if a given object implements the ContainerSvcCapabilities interface.
 */
export declare function instanceOfContainerSvcCapabilities(value: object): value is ContainerSvcCapabilities;
export declare function ContainerSvcCapabilitiesFromJSON(json: any): ContainerSvcCapabilities;
export declare function ContainerSvcCapabilitiesFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcCapabilities;
export declare function ContainerSvcCapabilitiesToJSON(json: any): ContainerSvcCapabilities;
export declare function ContainerSvcCapabilitiesToJSONTyped(value?: ContainerSvcCapabilities | null, ignoreDiscriminator?: boolean): any;
