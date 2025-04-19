/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ContainerSvcListContainersRequest
 */
export interface ContainerSvcListContainersRequest {
    /**
     *
     * @type {string}
     * @memberof ContainerSvcListContainersRequest
     */
    containerId?: string;
    /**
     *
     * @type {number}
     * @memberof ContainerSvcListContainersRequest
     */
    limit?: number;
    /**
     *
     * @type {string}
     * @memberof ContainerSvcListContainersRequest
     */
    nodeId?: string;
}
/**
 * Check if a given object implements the ContainerSvcListContainersRequest interface.
 */
export declare function instanceOfContainerSvcListContainersRequest(value: object): value is ContainerSvcListContainersRequest;
export declare function ContainerSvcListContainersRequestFromJSON(json: any): ContainerSvcListContainersRequest;
export declare function ContainerSvcListContainersRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcListContainersRequest;
export declare function ContainerSvcListContainersRequestToJSON(json: any): ContainerSvcListContainersRequest;
export declare function ContainerSvcListContainersRequestToJSONTyped(value?: ContainerSvcListContainersRequest | null, ignoreDiscriminator?: boolean): any;
