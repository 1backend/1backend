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
/**
 *
 * @export
 * @interface ContainerSvcErrorResponse
 */
export interface ContainerSvcErrorResponse {
    /**
     *
     * @type {string}
     * @memberof ContainerSvcErrorResponse
     */
    error?: string;
}
/**
 * Check if a given object implements the ContainerSvcErrorResponse interface.
 */
export declare function instanceOfContainerSvcErrorResponse(value: object): value is ContainerSvcErrorResponse;
export declare function ContainerSvcErrorResponseFromJSON(json: any): ContainerSvcErrorResponse;
export declare function ContainerSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcErrorResponse;
export declare function ContainerSvcErrorResponseToJSON(json: any): ContainerSvcErrorResponse;
export declare function ContainerSvcErrorResponseToJSONTyped(value?: ContainerSvcErrorResponse | null, ignoreDiscriminator?: boolean): any;
