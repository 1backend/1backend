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
 * @interface ContainerSvcGetContainerSummaryResponse
 */
export interface ContainerSvcGetContainerSummaryResponse {
    /**
     *
     * @type {string}
     * @memberof ContainerSvcGetContainerSummaryResponse
     */
    logs: string;
    /**
     *
     * @type {string}
     * @memberof ContainerSvcGetContainerSummaryResponse
     */
    status: string;
    /**
     * DEPRECATED. Summary contains both Status and Logs.
     * @type {string}
     * @memberof ContainerSvcGetContainerSummaryResponse
     */
    summary: string;
}
/**
 * Check if a given object implements the ContainerSvcGetContainerSummaryResponse interface.
 */
export declare function instanceOfContainerSvcGetContainerSummaryResponse(value: object): value is ContainerSvcGetContainerSummaryResponse;
export declare function ContainerSvcGetContainerSummaryResponseFromJSON(json: any): ContainerSvcGetContainerSummaryResponse;
export declare function ContainerSvcGetContainerSummaryResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcGetContainerSummaryResponse;
export declare function ContainerSvcGetContainerSummaryResponseToJSON(json: any): ContainerSvcGetContainerSummaryResponse;
export declare function ContainerSvcGetContainerSummaryResponseToJSONTyped(value?: ContainerSvcGetContainerSummaryResponse | null, ignoreDiscriminator?: boolean): any;
