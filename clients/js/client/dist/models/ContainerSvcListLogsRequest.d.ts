/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
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
 * @interface ContainerSvcListLogsRequest
 */
export interface ContainerSvcListLogsRequest {
    /**
     *
     * @type {string}
     * @memberof ContainerSvcListLogsRequest
     */
    containerId?: string;
    /**
     *
     * @type {number}
     * @memberof ContainerSvcListLogsRequest
     */
    limit?: number;
    /**
     *
     * @type {string}
     * @memberof ContainerSvcListLogsRequest
     */
    nodeId?: string;
}
/**
 * Check if a given object implements the ContainerSvcListLogsRequest interface.
 */
export declare function instanceOfContainerSvcListLogsRequest(value: object): value is ContainerSvcListLogsRequest;
export declare function ContainerSvcListLogsRequestFromJSON(json: any): ContainerSvcListLogsRequest;
export declare function ContainerSvcListLogsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcListLogsRequest;
export declare function ContainerSvcListLogsRequestToJSON(json: any): ContainerSvcListLogsRequest;
export declare function ContainerSvcListLogsRequestToJSONTyped(value?: ContainerSvcListLogsRequest | null, ignoreDiscriminator?: boolean): any;
