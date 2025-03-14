/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ContainerSvcPortMapping } from './ContainerSvcPortMapping';
/**
 *
 * @export
 * @interface ContainerSvcRunContainerResponse
 */
export interface ContainerSvcRunContainerResponse {
    /**
     * Ports is returned here as host ports might get mapped dynamically.
     * @type {Array<ContainerSvcPortMapping>}
     * @memberof ContainerSvcRunContainerResponse
     */
    ports?: Array<ContainerSvcPortMapping>;
    /**
     *
     * @type {boolean}
     * @memberof ContainerSvcRunContainerResponse
     */
    started?: boolean;
}
/**
 * Check if a given object implements the ContainerSvcRunContainerResponse interface.
 */
export declare function instanceOfContainerSvcRunContainerResponse(value: object): value is ContainerSvcRunContainerResponse;
export declare function ContainerSvcRunContainerResponseFromJSON(json: any): ContainerSvcRunContainerResponse;
export declare function ContainerSvcRunContainerResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcRunContainerResponse;
export declare function ContainerSvcRunContainerResponseToJSON(json: any): ContainerSvcRunContainerResponse;
export declare function ContainerSvcRunContainerResponseToJSONTyped(value?: ContainerSvcRunContainerResponse | null, ignoreDiscriminator?: boolean): any;
