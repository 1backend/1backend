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
 * @interface ContainerSvcResources
 */
export interface ContainerSvcResources {
    /**
     * CPU cores allocated to the container (e.g., 0.5 = 500m, 2 = 2 cores).
     * @type {number}
     * @memberof ContainerSvcResources
     */
    cpu?: number;
    /**
     * Disk space allocated to the container in megabytes.
     * @type {number}
     * @memberof ContainerSvcResources
     */
    diskMB?: number;
    /**
     * Memory allocated to the container in megabytes.
     * @type {number}
     * @memberof ContainerSvcResources
     */
    memoryMB?: number;
}
/**
 * Check if a given object implements the ContainerSvcResources interface.
 */
export declare function instanceOfContainerSvcResources(value: object): value is ContainerSvcResources;
export declare function ContainerSvcResourcesFromJSON(json: any): ContainerSvcResources;
export declare function ContainerSvcResourcesFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcResources;
export declare function ContainerSvcResourcesToJSON(json: any): ContainerSvcResources;
export declare function ContainerSvcResourcesToJSONTyped(value?: ContainerSvcResources | null, ignoreDiscriminator?: boolean): any;
