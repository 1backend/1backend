/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ContainerSvcContainer } from './ContainerSvcContainer';
/**
 *
 * @export
 * @interface ContainerSvcListContainersResponse
 */
export interface ContainerSvcListContainersResponse {
    /**
     *
     * @type {Array<ContainerSvcContainer>}
     * @memberof ContainerSvcListContainersResponse
     */
    containers?: Array<ContainerSvcContainer>;
}
/**
 * Check if a given object implements the ContainerSvcListContainersResponse interface.
 */
export declare function instanceOfContainerSvcListContainersResponse(value: object): value is ContainerSvcListContainersResponse;
export declare function ContainerSvcListContainersResponseFromJSON(json: any): ContainerSvcListContainersResponse;
export declare function ContainerSvcListContainersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcListContainersResponse;
export declare function ContainerSvcListContainersResponseToJSON(json: any): ContainerSvcListContainersResponse;
export declare function ContainerSvcListContainersResponseToJSONTyped(value?: ContainerSvcListContainersResponse | null, ignoreDiscriminator?: boolean): any;
